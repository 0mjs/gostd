package content

func init() {
	Register(&Package{
		Name:       "time/tzdata",
		ImportPath: "time/tzdata",
		Category:   "Misc",
		Summary:    "Embed the IANA timezone database into your binary. Import for side effect when deploying to stripped-down containers without /usr/share/zoneinfo.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Blank import", Code: `import _ "time/tzdata"

// Now time.LoadLocation works even in scratch / distroless images.`},
				},
			},
		},
	})
}
