package content

func init() {
	Register(&Package{
		Name:       "runtime/coverage",
		ImportPath: "runtime/coverage",
		Category:   "Runtime & Debug",
		Summary:    "Programmatically write or clear coverage counters at runtime (for long-running binaries built with -cover).",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Write counters mid-run", Code: `// Binary must be built with: go build -cover
if err := coverage.WriteMetaDir("covdir"); err != nil { log.Fatal(err) }
if err := coverage.WriteCountersDir("covdir"); err != nil { log.Fatal(err) }`},
					{Title: "Reset counters", Code: `coverage.ClearCounters()`},
				},
			},
		},
	})
}
