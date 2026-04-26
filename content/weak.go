package content

func init() {
	Register(&Package{
		Name:       "weak",
		ImportPath: "weak",
		Category:   "Misc",
		Summary:    "Weak pointers — references that don't prevent GC (Go 1.24+). Build caches and maps that shouldn't keep values alive.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Make + Value", Code: `p := &User{Name: "Ada"}
w := weak.Make(p)

// later:
if u := w.Value(); u != nil {
    fmt.Println(u.Name) // still alive
} else {
    // the target was collected
}`},
				},
			},
		},
	})
}
