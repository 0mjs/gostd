package content

func init() {
	Register(&Package{
		Name:       "cmp",
		ImportPath: "cmp",
		Category:   "Collections",
		Summary:    "Generic comparison helpers (Go 1.21+). The glue between the old world and slices.SortFunc / maps / binary search.",
		Sections: []Section{
			{
				Title: "Compare and Less",
				Examples: []Example{
					{
						Title: "cmp.Compare — the -1/0/+1 function",
						Notes: "Used with slices.SortFunc, slices.BinarySearchFunc, and similar. Handles NaN sensibly.",
						Code: `cmp.Compare(1, 2)   // -1
cmp.Compare(2, 2)   //  0
cmp.Compare(3, 2)   // +1`,
					},
					{
						Title: "Less — the boolean version",
						Code: `cmp.Less(1.0, 2.0)  // true`,
					},
					{
						Title: "Or — fall through to the next comparison",
						Notes: "Build multi-key comparators: first compare field A, if equal, compare field B.",
						Code: `type P struct{ Last, First string; Age int }
slices.SortFunc(people, func(a, b P) int {
    return cmp.Or(
        cmp.Compare(a.Last, b.Last),
        cmp.Compare(a.First, b.First),
        cmp.Compare(a.Age, b.Age),
    )
})`,
					},
					{
						Title: "Or for zero-value fallback (any type)",
						Notes: "cmp.Or on non-comparable-semantics returns the first non-zero value — a generic COALESCE.",
						Code: `name := cmp.Or(userInput, envVar, "default")`,
					},
				},
			},
		},
	})
}
