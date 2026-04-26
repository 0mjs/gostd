package content

// SearchEntry is one row in the client-side search index served at /search.json.
// The cmd-k palette fetches this once on first open, then fuzzy-matches in the
// browser — no server round-trip per keystroke.
//
// JSON keys are single-letter to keep the payload small — the index has
// thousands of rows.
type SearchEntry struct {
	Kind   string `json:"k"`           // "pkg" | "section" | "example" | "recipe"
	Name   string `json:"n"`           // text shown as the primary hit label
	Detail string `json:"d,omitempty"` // secondary line (summary, parent path, etc.)
	URL    string `json:"u"`           // /pkg/foo or /pkg/foo#anchor or /recipes#anchor
}

// BuildSearchIndex walks every registered package and recipe and produces a
// flat search index. Anchors are computed using the same Slugify+UniqueAnchor
// rules the renderer uses, so links land on the right element.
func BuildSearchIndex() []SearchEntry {
	var out []SearchEntry

	for _, p := range All() {
		out = append(out, SearchEntry{
			Kind:   "pkg",
			Name:   p.ImportPath,
			Detail: p.Summary,
			URL:    "/pkg/" + p.Name,
		})

		anchors := map[string]int{}
		for _, section := range p.Sections {
			sectionAnchor := UniqueAnchor(anchors, section.Title)
			out = append(out, SearchEntry{
				Kind:   "section",
				Name:   section.Title,
				Detail: p.ImportPath,
				URL:    "/pkg/" + p.Name + "#" + sectionAnchor,
			})
			for _, ex := range section.Examples {
				exAnchor := UniqueAnchor(anchors, section.Title+" "+ex.Title)
				out = append(out, SearchEntry{
					Kind:   "example",
					Name:   ex.Title,
					Detail: p.ImportPath + " · " + section.Title,
					URL:    "/pkg/" + p.Name + "#" + exAnchor,
				})
			}
		}
	}

	recAnchors := map[string]int{}
	for _, r := range AllRecipes() {
		anchor := UniqueAnchor(recAnchors, r.Title)
		out = append(out, SearchEntry{
			Kind:   "recipe",
			Name:   r.Title,
			Detail: "recipe · " + r.Group,
			URL:    "/recipes#" + anchor,
		})
	}

	return out
}
