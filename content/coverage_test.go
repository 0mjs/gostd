package content

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestGeneratedCoverageIncludesAllPublicStdPackages(t *testing.T) {
	stats, err := EnsureGeneratedPackages()
	if err != nil {
		t.Fatalf("EnsureGeneratedPackages: %v", err)
	}
	if stats.Total == 0 {
		t.Fatal("expected non-zero package coverage")
	}

	want, err := scanGeneratedPackagesForTest()
	if err != nil {
		t.Fatalf("scanStandardPackages: %v", err)
	}

	have := map[string]bool{}
	for _, p := range All() {
		have[p.ImportPath] = true
	}
	for _, p := range want {
		if !have[p.ImportPath] {
			t.Fatalf("missing std package coverage for %q", p.ImportPath)
		}
	}
}

func TestGeneratedReferenceSectionsAreMergedIntoCoveredPackages(t *testing.T) {
	if _, err := EnsureGeneratedPackages(); err != nil {
		t.Fatalf("EnsureGeneratedPackages: %v", err)
	}

	want, err := scanGeneratedPackagesForTest()
	if err != nil {
		t.Fatalf("scanStandardPackages: %v", err)
	}

	for _, generated := range want {
		got, ok := Get(generated.ImportPath)
		if !ok {
			t.Fatalf("package %q missing after generation", generated.ImportPath)
		}

		for _, wantSection := range generatedReferenceSections(generated.Sections) {
			gotSections := findSections(got.Sections, wantSection.Title)
			if len(gotSections) == 0 {
				t.Fatalf("package %q missing section %q", generated.ImportPath, wantSection.Title)
			}
			for _, wantExample := range wantSection.Examples {
				if !sectionGroupHasExample(gotSections, wantExample.Title) {
					t.Fatalf("package %q missing example %q in section %q", generated.ImportPath, wantExample.Title, wantSection.Title)
				}
			}
		}
	}
}

func findSections(sections []Section, title string) []Section {
	var matched []Section
	for _, section := range sections {
		if section.Title == title {
			matched = append(matched, section)
		}
	}
	return matched
}

func sectionGroupHasExample(sections []Section, title string) bool {
	for _, section := range sections {
		if hasExample(section.Examples, title) {
			return true
		}
	}
	return false
}

func hasExample(examples []Example, title string) bool {
	for _, example := range examples {
		if example.Title == title {
			return true
		}
	}
	return false
}

func TestPackageAdvisories(t *testing.T) {
	if _, err := EnsureGeneratedPackages(); err != nil {
		t.Fatalf("EnsureGeneratedPackages: %v", err)
	}

	tests := []struct {
		pkg  string
		kind string
	}{
		{pkg: "io/ioutil", kind: "deprecated"},
		{pkg: "crypto/md5", kind: "not-recommended"},
		{pkg: "net/smtp", kind: "frozen"},
		{pkg: "database/sql", kind: "unused"},
	}

	for _, tt := range tests {
		p, ok := Get(tt.pkg)
		if !ok {
			t.Fatalf("package %q missing", tt.pkg)
		}
		if !hasAdvisoryKind(p.Advisories, tt.kind) {
			t.Fatalf("package %q missing advisory %q", tt.pkg, tt.kind)
		}
	}

	if p, ok := Get("fmt"); !ok {
		t.Fatal("package fmt missing")
	} else if p.HasAdvisories() {
		t.Fatalf("package fmt should not have advisories: %+v", p.Advisories)
	}
}

func hasAdvisoryKind(advisories []Advisory, kind string) bool {
	for _, advisory := range advisories {
		if advisory.Kind == kind {
			return true
		}
	}
	return false
}

func scanGeneratedPackagesForTest() ([]*Package, error) {
	root := filepath.Join(runtime.GOROOT(), "src")
	importCounts, err := scanStdImportCounts(root)
	if err != nil {
		return nil, err
	}
	return scanStandardPackages(root, importCounts)
}
