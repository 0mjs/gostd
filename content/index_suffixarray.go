package content

func init() {
	Register(&Package{
		Name:       "index/suffixarray",
		ImportPath: "index/suffixarray",
		Category:   "Misc",
		Summary:    "Substring index backed by a suffix array. Fast repeated lookups into a single large corpus.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "New + Lookup + FindAllIndex", Code: `data := []byte("banana bandana")
idx := suffixarray.New(data)
positions := idx.Lookup([]byte("an"), -1)
fmt.Println(positions) // offsets of all "an"

re := regexp.MustCompile("ban.")
ranges := idx.FindAllIndex(re, -1)`},
				},
			},
		},
	})
}
