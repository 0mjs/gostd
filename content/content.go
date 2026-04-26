package content

import (
	"sort"
	"strings"
)

// Package represents a stdlib package we teach.
type Package struct {
	Name       string // e.g. "fmt" or "path/filepath" — used as URL slug and sidebar label
	ImportPath string // e.g. "fmt" or "path/filepath"
	Category   string
	Summary    string
	Sections   []Section
	Generated  bool
	Advisories []Advisory
}

// CheatRow is one line in a package's at-a-glance cheatsheet: a real-world task
// paired with the smallest Go expression that does it.
type CheatRow struct {
	Task string // e.g. "Read a whole file"
	Code string // e.g. `b, err := os.ReadFile("path")`
}

// cheatsheets and tldrs are populated by RegisterCheatsheet, which is safe to
// call from any init() regardless of which file the target package was defined
// in (init order across files is alphabetical and would race otherwise).
var (
	tldrs       = map[string]string{}
	cheatsheets = map[string][]CheatRow{}
)

// RegisterCheatsheet attaches a TL;DR + at-a-glance rows to a package by
// import path. Lookups via the Package methods read these out at render time.
func RegisterCheatsheet(importPath, tldr string, rows []CheatRow) {
	tldrs[importPath] = tldr
	cheatsheets[importPath] = rows
}

type Advisory struct {
	Kind     string
	Label    string
	Message  string
	Priority int
}

type Section struct {
	Title       string
	Description string
	Examples    []Example
}

type Example struct {
	Title    string
	Notes    string // optional prose: when to reach for this, how it differs
	Code     string
	Language string
	Output   string // optional — leave empty if it's not a clean comparable output
}

type Category struct {
	Name     string
	Order    int
	Packages []*Package
}

var (
	registry = map[string]*Package{}
	order    []string
)

// Register adds a package to the global registry. Called from init() funcs.
func Register(p *Package) {
	if _, dup := registry[p.Name]; dup {
		panic("duplicate package: " + p.Name)
	}
	registry[p.Name] = p
	order = append(order, p.Name)
}

func Get(name string) (*Package, bool) {
	p, ok := registry[name]
	return p, ok
}

func All() []*Package {
	out := make([]*Package, 0, len(order))
	for _, n := range order {
		out = append(out, registry[n])
	}
	return out
}

func (p *Package) HasAdvisories() bool {
	return p != nil && len(p.Advisories) > 0
}

func (p *Package) Cheatsheet() []CheatRow {
	if p == nil {
		return nil
	}
	return cheatsheets[p.ImportPath]
}

func (p *Package) HasCheatsheet() bool {
	return p != nil && len(cheatsheets[p.ImportPath]) > 0
}

func (p *Package) TLDR() string {
	if p == nil {
		return ""
	}
	return tldrs[p.ImportPath]
}

func (p *Package) HasTLDR() bool {
	return p != nil && strings.TrimSpace(tldrs[p.ImportPath]) != ""
}

func (p *Package) PrimaryAdvisory() *Advisory {
	if p == nil || len(p.Advisories) == 0 {
		return nil
	}
	best := 0
	for i := 1; i < len(p.Advisories); i++ {
		if p.Advisories[i].Priority < p.Advisories[best].Priority {
			best = i
		}
	}
	return &p.Advisories[best]
}

// categoryOrder controls the order categories are shown in the sidebar.
var categoryOrder = map[string]int{
	"Formatting & Strings":   1,
	"I/O & Files":            2,
	"Time & Context":         3,
	"Concurrency":            4,
	"Collections":            5,
	"Encoding":               6,
	"Hashing":                7,
	"Networking":             8,
	"Errors & Logging":       9,
	"Math":                   10,
	"Crypto":                 11,
	"CLI & Runtime":          12,
	"Archives & Compression": 13,
	"Containers":             14,
	"Testing":                15,
	"Templates":              16,
	"Reflection & Unsafe":    17,
	"Runtime & Debug":        18,
	"Image":                  19,
	"Database":               20,
	"Commands & Toolchain":   21,
	"Go Tooling":             22,
	"Misc":                   23,
}

func AllCategories() []Category {
	byName := map[string]*Category{}
	var cats []*Category
	for _, p := range All() {
		c, ok := byName[p.Category]
		if !ok {
			ord := categoryOrder[p.Category]
			if ord == 0 {
				ord = 99
			}
			c = &Category{Name: p.Category, Order: ord}
			byName[p.Category] = c
			cats = append(cats, c)
		}
		c.Packages = append(c.Packages, p)
	}
	sort.SliceStable(cats, func(i, j int) bool {
		if cats[i].Order != cats[j].Order {
			return cats[i].Order < cats[j].Order
		}
		return cats[i].Name < cats[j].Name
	})
	for _, c := range cats {
		sort.SliceStable(c.Packages, func(i, j int) bool {
			return c.Packages[i].Name < c.Packages[j].Name
		})
	}
	out := make([]Category, len(cats))
	for i, c := range cats {
		out[i] = *c
	}
	return out
}
