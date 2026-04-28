package content

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"sync"
)

type CoverageStats struct {
	Curated   int
	Generated int
	Total     int
}

var (
	ensureOnce  sync.Once
	ensureStats CoverageStats
	ensureErr   error
)

func EnsureGeneratedPackages() (CoverageStats, error) {
	ensureOnce.Do(func() {
		curated := len(registry)
		root := filepath.Join(runtime.GOROOT(), "src")
		if _, err := os.Stat(root); err != nil {
			if os.IsNotExist(err) {
				ensureStats.Curated = curated
				ensureStats.Total = len(registry)
				return
			}
			ensureErr = err
			return
		}
		importCounts, err := scanStdImportCounts(root)
		if err != nil {
			ensureErr = err
			return
		}
		pkgs, err := scanStandardPackages(root, importCounts)
		if err != nil {
			ensureErr = err
			return
		}
		for _, p := range pkgs {
			if existing, ok := registry[p.ImportPath]; ok {
				existing.Sections = append(existing.Sections, generatedReferenceSections(p.Sections)...)
				existing.Advisories = p.Advisories
				continue
			}
			Register(p)
			ensureStats.Generated++
		}
		ensureStats.Curated = curated
		ensureStats.Total = len(registry)
	})
	return ensureStats, ensureErr
}

func generatedReferenceSections(sections []Section) []Section {
	if len(sections) <= 1 {
		return nil
	}
	return sections[1:]
}

func scanStandardPackages(root string, importCounts map[string]int) ([]*Package, error) {
	var pkgs []*Package
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == "." {
			return nil
		}
		if skipPackageDir(rel) {
			return filepath.SkipDir
		}
		if !isPublicStdPackage(rel) {
			return nil
		}

		pkg, err := generatedPackage(path, rel, importCounts[rel])
		if err != nil {
			return err
		}
		if pkg != nil {
			pkgs = append(pkgs, pkg)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	slices.SortFunc(pkgs, func(a, b *Package) int {
		return strings.Compare(a.ImportPath, b.ImportPath)
	})
	return pkgs, nil
}

func generatedPackage(dir, importPath string, importCount int) (*Package, error) {
	if !hasPackageSources(dir) {
		return nil, nil
	}

	fset := token.NewFileSet()
	astPkgs, err := parser.ParseDir(fset, dir, sourceFilter, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	if len(astPkgs) == 0 {
		return nil, nil
	}

	docPkg := buildDocPackage(astPkgs, importPath)
	if docPkg == nil {
		return nil, nil
	}

	fullDoc := cleanDoc(docPkg.Doc)
	summary := synopsis(fullDoc)
	if summary == "" {
		summary = "Standard library package " + importPath + "."
	}
	if fullDoc == "" {
		fullDoc = summary
	}

	sections := []Section{
		{
			Title:       "Overview",
			Description: fullDoc,
			Examples: []Example{
				{
					Title:    "Start here",
					Notes:    generatedStartNotes(importPath),
					Code:     generatedStartCode(importPath),
					Language: generatedStartLanguage(importPath),
				},
			},
		},
	}
	sections = append(sections, packageReferenceSections(docPkg, importPath)...)

	return &Package{
		Name:       importPath,
		ImportPath: importPath,
		Category:   categoryForImportPath(importPath),
		Summary:    summary,
		Generated:  true,
		Sections:   sections,
		Advisories: buildAdvisories(importPath, fullDoc, importCount),
	}, nil
}

func scanStdImportCounts(root string) (map[string]int, error) {
	packages, err := stdPackagePaths(root)
	if err != nil {
		return nil, err
	}
	public := map[string]bool{}
	for _, pkg := range packages {
		public[pkg] = true
	}

	counts := map[string]int{}
	for _, pkg := range packages {
		dir := filepath.Join(root, filepath.FromSlash(pkg))
		fset := token.NewFileSet()
		astPkgs, err := parser.ParseDir(fset, dir, sourceFilter, parser.ImportsOnly)
		if err != nil {
			return nil, err
		}

		seen := map[string]bool{}
		for _, astPkg := range astPkgs {
			for _, file := range astPkg.Files {
				for _, spec := range file.Imports {
					imp := strings.Trim(spec.Path.Value, `"`)
					if imp == pkg || !public[imp] || seen[imp] {
						continue
					}
					seen[imp] = true
					counts[imp]++
				}
			}
		}
	}

	return counts, nil
}

func stdPackagePaths(root string) ([]string, error) {
	var packages []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == "." {
			return nil
		}
		if skipPackageDir(rel) {
			return filepath.SkipDir
		}
		if !isPublicStdPackage(rel) || !hasPackageSources(path) {
			return nil
		}
		packages = append(packages, rel)
		return nil
	})
	if err != nil {
		return nil, err
	}
	slices.Sort(packages)
	return packages, nil
}

func sourceFilter(fi fs.FileInfo) bool {
	name := fi.Name()
	if fi.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
		return false
	}
	return !strings.HasPrefix(name, ".") && !strings.HasPrefix(name, "_")
}

func hasPackageSources(dir string) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, f := range files {
		name := f.Name()
		if f.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_") {
			continue
		}
		return true
	}
	return false
}

func buildDocPackage(astPkgs map[string]*ast.Package, importPath string) *doc.Package {
	names := make([]string, 0, len(astPkgs))
	for name := range astPkgs {
		if strings.HasSuffix(name, "_test") {
			continue
		}
		names = append(names, name)
	}
	slices.Sort(names)
	if len(names) == 0 {
		return nil
	}

	if !strings.HasPrefix(importPath, "cmd/") {
		if pkg, ok := astPkgs[packageQualifier(importPath)]; ok {
			return doc.New(pkg, importPath, doc.AllDecls)
		}
		for _, name := range names {
			if name == "main" {
				continue
			}
			return doc.New(astPkgs[name], importPath, doc.AllDecls)
		}
	}

	if pkg, ok := astPkgs["main"]; ok {
		return doc.New(pkg, importPath, doc.AllDecls)
	}
	return doc.New(astPkgs[names[0]], importPath, doc.AllDecls)
}

func packageReferenceSections(docPkg *doc.Package, importPath string) []Section {
	var sections []Section

	if section := packageLevelSection(docPkg, importPath); len(section.Examples) > 0 {
		sections = append(sections, section)
	}

	for _, typ := range docPkg.Types {
		if !ast.IsExported(typ.Name) {
			continue
		}
		if section := typeSection(typ, importPath); len(section.Examples) > 0 {
			sections = append(sections, section)
		}
	}
	return sections
}

func packageLevelSection(docPkg *doc.Package, importPath string) Section {
	section := Section{
		Title:       "Package-level API",
		Description: "Exported constants, variables, and functions from `" + importPath + "`.",
	}

	for _, c := range docPkg.Consts {
		if ex, ok := valueExample(c, "const", importPath); ok {
			section.Examples = append(section.Examples, ex)
		}
	}
	for _, v := range docPkg.Vars {
		if ex, ok := valueExample(v, "var", importPath); ok {
			section.Examples = append(section.Examples, ex)
		}
	}
	for _, fn := range docPkg.Funcs {
		if !ast.IsExported(fn.Name) {
			continue
		}
		section.Examples = append(section.Examples, functionExample(fn, importPath, "function"))
	}

	return section
}

func typeSection(typ *doc.Type, importPath string) Section {
	desc := cleanDoc(typ.Doc)
	if desc == "" {
		desc = "Exported type and related API from `" + importPath + "`."
	}

	section := Section{
		Title:       typ.Name,
		Description: desc,
		Examples: []Example{
			{
				Title:    "Declaration",
				Notes:    beginnerNote("type", typ.Name, synopsis(desc)),
				Code:     formatDecl(typ.Decl),
				Language: "go",
			},
		},
	}

	for _, c := range typ.Consts {
		if ex, ok := valueExample(c, "const", importPath); ok {
			section.Examples = append(section.Examples, ex)
		}
	}
	for _, v := range typ.Vars {
		if ex, ok := valueExample(v, "var", importPath); ok {
			section.Examples = append(section.Examples, ex)
		}
	}
	for _, fn := range typ.Funcs {
		if !ast.IsExported(fn.Name) {
			continue
		}
		section.Examples = append(section.Examples, functionExample(fn, importPath, "function"))
	}
	for _, method := range typ.Methods {
		if !ast.IsExported(method.Name) {
			continue
		}
		section.Examples = append(section.Examples, functionExample(method, importPath, "method"))
	}

	return section
}

func valueExample(v *doc.Value, kind, importPath string) (Example, bool) {
	names := exportedNames(v.Names)
	if len(names) == 0 {
		return Example{}, false
	}
	title := strings.Join(names, ", ")
	summary := synopsis(cleanDoc(v.Doc))
	if summary == "" {
		summary = "Exported " + kind + " from `" + importPath + "`."
	}
	return Example{
		Title:    title,
		Notes:    beginnerNote(kind, title, summary),
		Code:     formatDecl(v.Decl),
		Language: "go",
	}, true
}

func functionExample(fn *doc.Func, importPath, kind string) Example {
	summary := synopsis(cleanDoc(fn.Doc))
	if summary == "" {
		summary = "Exported " + kind + " from `" + importPath + "`."
	}
	return Example{
		Title:    fn.Name,
		Notes:    beginnerNote(kind, fn.Name, summary),
		Code:     functionReferenceCode(importPath, fn),
		Language: "go",
	}
}

func functionReferenceCode(importPath string, fn *doc.Func) string {
	signature := formatFuncSignature(fn.Decl)
	callShape := buildCallShape(importPath, fn.Decl)
	if callShape == "" {
		return signature
	}
	return signature + "\n\n" + callShape
}

func buildCallShape(importPath string, decl *ast.FuncDecl) string {
	if decl == nil || decl.Type == nil {
		return ""
	}

	target := packageCallTarget(importPath, decl)
	args := callArgs(decl.Type.Params)
	results := callResultNames(decl.Type.Results)
	call := target + "(" + strings.Join(args, ", ") + ")"

	var lines []string
	if recvType := receiverTypeName(importPath, decl); recvType != "" {
		lines = append(lines, "var value "+recvType)
	}

	switch {
	case len(results) == 0:
		lines = append(lines, call)
	case lastIsError(decl.Type.Results):
		assign := strings.Join(results, ", ") + " := " + call
		lines = append(lines, assign)
		lines = append(lines, "if err != nil {")
		lines = append(lines, "\t// handle the error")
		lines = append(lines, "}")
		for _, name := range results[:len(results)-1] {
			lines = append(lines, "_ = "+name)
		}
	default:
		assign := strings.Join(results, ", ") + " := " + call
		lines = append(lines, assign)
		for _, name := range results {
			lines = append(lines, "_ = "+name)
		}
	}

	return strings.Join(lines, "\n")
}

func packageCallTarget(importPath string, decl *ast.FuncDecl) string {
	if decl.Recv != nil && len(decl.Recv.List) > 0 {
		return "value." + decl.Name.Name
	}
	return packageQualifier(importPath) + "." + decl.Name.Name
}

func callArgs(fields *ast.FieldList) []string {
	if fields == nil || len(fields.List) == 0 {
		return nil
	}
	var args []string
	index := 1
	for _, field := range fields.List {
		if len(field.Names) == 0 {
			args = append(args, fmt.Sprintf("arg%d", index))
			index++
			continue
		}
		for _, name := range field.Names {
			arg := name.Name
			if arg == "" || arg == "_" {
				arg = fmt.Sprintf("arg%d", index)
			}
			args = append(args, arg)
			index++
		}
	}
	return args
}

func callResultNames(fields *ast.FieldList) []string {
	if fields == nil || len(fields.List) == 0 {
		return nil
	}
	var names []string
	valueIndex := 1
	for _, field := range fields.List {
		typeName := exprString(field.Type)
		if len(field.Names) == 0 {
			names = append(names, resultName(typeName, valueIndex))
			valueIndex++
			continue
		}
		for _, name := range field.Names {
			got := name.Name
			if got == "" || got == "_" {
				got = resultName(typeName, valueIndex)
			}
			names = append(names, got)
			valueIndex++
		}
	}
	return names
}

func resultName(typeName string, index int) string {
	if typeName == "error" {
		return "err"
	}
	if index == 1 {
		return "result"
	}
	return fmt.Sprintf("result%d", index)
}

func lastIsError(fields *ast.FieldList) bool {
	if fields == nil || len(fields.List) == 0 {
		return false
	}
	last := fields.List[len(fields.List)-1]
	return exprString(last.Type) == "error"
}

func receiverTypeName(importPath string, decl *ast.FuncDecl) string {
	if decl.Recv == nil || len(decl.Recv.List) == 0 {
		return ""
	}
	return qualifyReceiverType(importPath, decl.Recv.List[0].Type)
}

func qualifyReceiverType(importPath string, expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return "*" + qualifyReceiverType(importPath, t.X)
	case *ast.Ident:
		return packageQualifier(importPath) + "." + t.Name
	case *ast.IndexExpr, *ast.IndexListExpr, *ast.SelectorExpr:
		return exprString(expr)
	default:
		return exprString(expr)
	}
}

func formatFuncSignature(decl *ast.FuncDecl) string {
	if decl == nil {
		return ""
	}
	copyDecl := *decl
	copyDecl.Body = nil
	return formatNode(&copyDecl)
}

func formatDecl(node ast.Node) string {
	if node == nil {
		return ""
	}
	return formatNode(node)
}

func formatNode(node ast.Node) string {
	var buf bytes.Buffer
	if err := format.Node(&buf, token.NewFileSet(), node); err != nil {
		return ""
	}
	return strings.TrimSpace(buf.String())
}

func exprString(expr ast.Expr) string {
	if expr == nil {
		return ""
	}
	return formatNode(expr)
}

func exportedNames(names []string) []string {
	var out []string
	for _, name := range names {
		if ast.IsExported(name) {
			out = append(out, name)
		}
	}
	return out
}

func beginnerNote(kind, name, summary string) string {
	switch kind {
	case "type":
		return "Reach for `" + name + "` when this value is the center of the workflow you are building. " + summary
	case "method":
		return "Call `" + name + "` after you already have the receiver value. Start with the signature, then use the call shape as a template. " + summary
	case "function":
		return "Use `" + name + "` when you need the behavior described here. The signature shows the contract; the call shape shows how it fits into real code. " + summary
	case "const":
		return "Use `" + name + "` when you need the package-defined constant rather than inventing your own literal. " + summary
	case "var":
		return "Use `" + name + "` when the package exposes a ready-made variable or sentinel you should share with other code. " + summary
	default:
		return summary
	}
}

func cleanDoc(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func buildAdvisories(importPath, fullDoc string, importCount int) []Advisory {
	var advisories []Advisory

	if msg := packageDeprecatedMessage(fullDoc); msg != "" {
		advisories = append(advisories, Advisory{
			Kind:     "deprecated",
			Label:    "deprecated",
			Message:  msg,
			Priority: 1,
		})
	}
	if msg := packageInsecureMessage(fullDoc); msg != "" {
		advisories = append(advisories, Advisory{
			Kind:     "not-recommended",
			Label:    "insecure",
			Message:  msg,
			Priority: 2,
		})
	}
	if msg := packageFrozenMessage(fullDoc); msg != "" {
		advisories = append(advisories, Advisory{
			Kind:     "frozen",
			Label:    "frozen",
			Message:  msg,
			Priority: 3,
		})
	}
	if !strings.HasPrefix(importPath, "cmd/") && importCount == 0 {
		advisories = append(advisories, Advisory{
			Kind:     "unused",
			Label:    "unused",
			Message:  "No other public standard-library package imports this package in the Go source tree used by this app. This is informational only, not a recommendation against using it.",
			Priority: 4,
		})
	}

	return advisories
}

func packageDeprecatedMessage(fullDoc string) string {
	return paragraphWithPrefix(fullDoc, "Deprecated:")
}

func packageInsecureMessage(fullDoc string) string {
	for _, paragraph := range splitParagraphs(fullDoc) {
		if strings.Contains(strings.ToLower(paragraph), "cryptographically broken and should not be used") {
			return paragraph
		}
	}
	return ""
}

func packageFrozenMessage(fullDoc string) string {
	for _, paragraph := range splitParagraphs(fullDoc) {
		if strings.Contains(strings.ToLower(paragraph), "frozen and is not accepting new features") {
			return paragraph
		}
	}
	return ""
}

func paragraphWithPrefix(fullDoc, prefix string) string {
	for _, paragraph := range splitParagraphs(fullDoc) {
		if strings.Contains(paragraph, prefix) {
			return paragraph
		}
	}
	return ""
}

func splitParagraphs(s string) []string {
	s = cleanDoc(s)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, "\n\n")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		out = append(out, strings.Join(strings.Fields(part), " "))
	}
	return out
}

func synopsis(s string) string {
	if s == "" {
		return ""
	}
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.Join(strings.Fields(s), " ")
	for _, end := range []string{". ", "! ", "? "} {
		if i := strings.Index(s, end); i >= 0 {
			return strings.TrimSpace(s[:i+1])
		}
	}
	return s
}

func generatedStartNotes(importPath string) string {
	if strings.HasPrefix(importPath, "cmd/") {
		return "Start with the command invocation below, then read the exported API reference on this page when the command also exposes reusable Go packages or flags."
	}
	return "Start with the import, then scan the package-level API and per-type sections below. Read the signature first, then copy the call shape into your own program and replace the placeholders."
}

func generatedStartCode(importPath string) string {
	if strings.HasPrefix(importPath, "cmd/") {
		name := strings.TrimPrefix(importPath, "cmd/")
		switch name {
		case "go":
			return "go help"
		case "gofmt":
			return "gofmt -w ./..."
		default:
			return "go tool " + name
		}
	}
	return "import \"" + importPath + "\""
}

func generatedStartLanguage(importPath string) string {
	if strings.HasPrefix(importPath, "cmd/") {
		return "bash"
	}
	return "go"
}

func packageQualifier(importPath string) string {
	parts := strings.Split(importPath, "/")
	return parts[len(parts)-1]
}

func skipPackageDir(rel string) bool {
	for _, part := range strings.Split(rel, "/") {
		if part == "testdata" || part == "vendor" || part == "internal" {
			return true
		}
		if strings.HasPrefix(part, ".") || strings.HasPrefix(part, "_") {
			return true
		}
	}
	return false
}

func isPublicStdPackage(rel string) bool {
	switch {
	case strings.HasPrefix(rel, "cmd/"):
		return true
	case strings.HasPrefix(rel, "archive/"),
		strings.HasPrefix(rel, "compress/"),
		strings.HasPrefix(rel, "container/"),
		strings.HasPrefix(rel, "crypto/"),
		strings.HasPrefix(rel, "database/"),
		strings.HasPrefix(rel, "debug/"),
		strings.HasPrefix(rel, "encoding/"),
		strings.HasPrefix(rel, "go/"),
		strings.HasPrefix(rel, "hash/"),
		strings.HasPrefix(rel, "html/"),
		strings.HasPrefix(rel, "image/"),
		strings.HasPrefix(rel, "index/"),
		strings.HasPrefix(rel, "io/"),
		strings.HasPrefix(rel, "log/"),
		strings.HasPrefix(rel, "math/"),
		strings.HasPrefix(rel, "mime/"),
		strings.HasPrefix(rel, "net/"),
		strings.HasPrefix(rel, "os/"),
		strings.HasPrefix(rel, "path/"),
		strings.HasPrefix(rel, "regexp/"),
		strings.HasPrefix(rel, "runtime/"),
		strings.HasPrefix(rel, "sync/"),
		strings.HasPrefix(rel, "syscall/"),
		strings.HasPrefix(rel, "testing/"),
		strings.HasPrefix(rel, "text/"),
		strings.HasPrefix(rel, "time/"),
		strings.HasPrefix(rel, "unicode/"):
		return true
	}
	switch rel {
	case "arena", "bufio", "builtin", "bytes", "cmp", "context", "crypto", "embed", "encoding",
		"errors", "expvar", "flag", "fmt", "hash", "html", "image", "io", "iter", "log", "maps",
		"math", "mime", "net", "os", "path", "plugin", "reflect", "regexp", "runtime", "slices",
		"sort", "strconv", "strings", "structs", "sync", "syscall", "testing", "time", "unicode",
		"unique", "unsafe", "weak":
		return true
	}
	return false
}

func categoryForImportPath(importPath string) string {
	switch {
	case strings.HasPrefix(importPath, "cmd/"):
		return "Commands & Toolchain"
	case strings.HasPrefix(importPath, "archive/"), strings.HasPrefix(importPath, "compress/"):
		return "Archives & Compression"
	case importPath == "bufio", importPath == "io", strings.HasPrefix(importPath, "io/"),
		importPath == "os", strings.HasPrefix(importPath, "os/"), importPath == "path",
		strings.HasPrefix(importPath, "path/"):
		return "I/O & Files"
	case importPath == "time", strings.HasPrefix(importPath, "time/"), importPath == "context":
		return "Time & Context"
	case importPath == "sync", strings.HasPrefix(importPath, "sync/"):
		return "Concurrency"
	case importPath == "bytes", importPath == "cmp", importPath == "iter", importPath == "maps",
		importPath == "slices", importPath == "sort", strings.HasPrefix(importPath, "container/"):
		return "Collections"
	case importPath == "encoding", strings.HasPrefix(importPath, "encoding/"), strings.HasPrefix(importPath, "mime/"):
		return "Encoding"
	case importPath == "hash", strings.HasPrefix(importPath, "hash/"):
		return "Hashing"
	case importPath == "net", strings.HasPrefix(importPath, "net/"):
		return "Networking"
	case importPath == "errors", importPath == "expvar", importPath == "log", strings.HasPrefix(importPath, "log/"):
		return "Errors & Logging"
	case importPath == "math", strings.HasPrefix(importPath, "math/"):
		return "Math"
	case importPath == "crypto", strings.HasPrefix(importPath, "crypto/"):
		return "Crypto"
	case importPath == "flag", importPath == "plugin", importPath == "runtime", strings.HasPrefix(importPath, "runtime/"),
		importPath == "syscall", strings.HasPrefix(importPath, "syscall/"), importPath == "unsafe":
		return "CLI & Runtime"
	case strings.HasPrefix(importPath, "testing/"), importPath == "testing":
		return "Testing"
	case strings.HasPrefix(importPath, "text/template"), importPath == "html/template":
		return "Templates"
	case importPath == "reflect", importPath == "structs", importPath == "weak":
		return "Reflection & Unsafe"
	case strings.HasPrefix(importPath, "debug/"):
		return "Runtime & Debug"
	case importPath == "image", strings.HasPrefix(importPath, "image/"):
		return "Image"
	case strings.HasPrefix(importPath, "database/"):
		return "Database"
	case strings.HasPrefix(importPath, "go/"):
		return "Go Tooling"
	case importPath == "fmt", importPath == "html", importPath == "regexp", strings.HasPrefix(importPath, "regexp/"),
		importPath == "strconv", importPath == "strings", importPath == "text/scanner", importPath == "text/tabwriter":
		return "Formatting & Strings"
	default:
		return "Misc"
	}
}
