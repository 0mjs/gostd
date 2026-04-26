package main

import (
	"cmp"
	"compress/gzip"
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"stdlearn/content"
)

//go:embed templates
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

var tmplFuncs = template.FuncMap{
	"add": func(a, b int) int { return a + b },
}

var tmpl = template.Must(template.New("").Funcs(tmplFuncs).ParseFS(templateFS, "templates/*.html"))

type pageData struct {
	Title      string
	ActivePkg  string
	ActiveNav  string // "home" | "recipes" | "" — controls top-nav highlight
	Categories []content.Category
	Package    *content.Package
	Coverage   content.CoverageStats
	Sections   []sectionView
	TOC        []tocSection
	LongPage   bool
	HasTOC     bool
	StartHere  []content.StartHereEntry
	Recipes    []content.RecipeGroup
	RecipeAnch map[string]string // recipe title -> anchor
}

type sectionView struct {
	Title       string
	Description string
	Examples    []exampleView
	Anchor      string
	ItemCount   int
	Collapsible bool
	Open        bool
}

type exampleView struct {
	Title    string
	Notes    string
	Code     string
	Language string
	Output   string
	Anchor   string
}

type tocSection struct {
	Title     string
	Anchor    string
	ItemCount int
	Items     []tocItem
}

type tocItem struct {
	Title  string
	Anchor string
}

func main() {
	coverage, err := content.EnsureGeneratedPackages()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	staticSub, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(staticSub)))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		home(w, r, coverage)
	})
	mux.HandleFunc("GET /pkg/{name...}", func(w http.ResponseWriter, r *http.Request) {
		pkg(w, r, coverage)
	})
	mux.HandleFunc("GET /recipes", func(w http.ResponseWriter, r *http.Request) {
		recipes(w, r, coverage)
	})
	mux.HandleFunc("GET /search.json", searchJSON)

	addr := ":" + cmp.Or(os.Getenv("PORT"), "8080")
	log.Printf("listening on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func home(w http.ResponseWriter, _ *http.Request, coverage content.CoverageStats) {
	if err := tmpl.ExecuteTemplate(w, "home.html", pageData{
		Title:      "Go stdlib — tour",
		ActiveNav:  "home",
		Categories: content.AllCategories(),
		Coverage:   coverage,
		StartHere:  content.StartHere(),
	}); err != nil {
		log.Printf("home: %v", err)
	}
}

func pkg(w http.ResponseWriter, r *http.Request, coverage content.CoverageStats) {
	name := r.PathValue("name")
	p, ok := content.Get(name)
	if !ok {
		http.NotFound(w, r)
		return
	}
	data := buildPackagePageData(p, coverage)
	if err := tmpl.ExecuteTemplate(w, "package.html", data); err != nil {
		log.Printf("pkg: %v", err)
	}
}

func recipes(w http.ResponseWriter, _ *http.Request, coverage content.CoverageStats) {
	groups := content.RecipeGroups()
	anchors := map[string]int{}
	rAnch := map[string]string{}
	for _, g := range groups {
		for _, r := range g.Recipes {
			rAnch[r.Title] = content.UniqueAnchor(anchors, r.Title)
		}
	}
	if err := tmpl.ExecuteTemplate(w, "recipes.html", pageData{
		Title:      "Recipes — Go stdlib",
		ActiveNav:  "recipes",
		Categories: content.AllCategories(),
		Coverage:   coverage,
		Recipes:    groups,
		RecipeAnch: rAnch,
	}); err != nil {
		log.Printf("recipes: %v", err)
	}
}

func searchJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=60")
	w.Header().Set("Vary", "Accept-Encoding")

	index := content.BuildSearchIndex()
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		if err := json.NewEncoder(gz).Encode(index); err != nil {
			log.Printf("search: %v", err)
		}
		return
	}
	if err := json.NewEncoder(w).Encode(index); err != nil {
		log.Printf("search: %v", err)
	}
}

func buildPackagePageData(p *content.Package, coverage content.CoverageStats) pageData {
	sections, toc, longPage := buildPackageViews(p)
	return pageData{
		Title:      p.ImportPath + " — Go stdlib",
		ActivePkg:  p.Name,
		Categories: content.AllCategories(),
		Package:    p,
		Coverage:   coverage,
		Sections:   sections,
		TOC:        toc,
		LongPage:   longPage,
		HasTOC:     len(toc) > 0,
	}
}

func buildPackageViews(p *content.Package) ([]sectionView, []tocSection, bool) {
	totalExamples := 0
	for _, section := range p.Sections {
		totalExamples += len(section.Examples)
	}
	longPage := len(p.Sections) >= 8 || totalExamples >= 40

	anchors := map[string]int{}
	sections := make([]sectionView, 0, len(p.Sections))
	toc := make([]tocSection, 0, len(p.Sections))

	for i, section := range p.Sections {
		sectionAnchor := content.UniqueAnchor(anchors, section.Title)
		view := sectionView{
			Title:       section.Title,
			Description: section.Description,
			Anchor:      sectionAnchor,
			ItemCount:   len(section.Examples),
			Collapsible: longPage && i >= 2 && len(section.Examples) > 0,
			Open:        !(longPage && i >= 2 && len(section.Examples) > 0),
		}
		tocSection := tocSection{
			Title:     section.Title,
			Anchor:    sectionAnchor,
			ItemCount: len(section.Examples),
		}

		for _, example := range section.Examples {
			exampleAnchor := content.UniqueAnchor(anchors, section.Title+" "+example.Title)
			view.Examples = append(view.Examples, exampleView{
				Title:    example.Title,
				Notes:    example.Notes,
				Code:     example.Code,
				Language: example.Language,
				Output:   example.Output,
				Anchor:   exampleAnchor,
			})
			tocSection.Items = append(tocSection.Items, tocItem{
				Title:  example.Title,
				Anchor: exampleAnchor,
			})
		}

		sections = append(sections, view)
		toc = append(toc, tocSection)
	}

	return sections, toc, longPage
}
