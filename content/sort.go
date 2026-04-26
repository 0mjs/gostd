package content

func init() {
	Register(&Package{
		Name:       "sort",
		ImportPath: "sort",
		Category:   "Collections",
		Summary:    "Sorting for slices and any sort.Interface. For generic slice sorting prefer the newer slices package.",
		Sections: []Section{
			{
				Title: "sort.Slice — the everyday case",
				Examples: []Example{
					{
						Title: "Sort with a less func",
						Code: `people := []struct{ Name string; Age int }{
    {"Ada", 36}, {"Alan", 41}, {"Grace", 85},
}
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
fmt.Println(people)`,
						Output: `[{Ada 36} {Alan 41} {Grace 85}]
`,
					},
					{
						Title: "SliceStable — preserves order of equals",
						Notes: "Use when you sort on a secondary key and want the primary key's relative order preserved.",
						Code: `sort.SliceStable(data, func(i, j int) bool {
    return data[i].Priority < data[j].Priority
})`,
					},
				},
			},
			{
				Title: "Built-in helpers for common slice types",
				Examples: []Example{
					{
						Title: "Ints, Strings, Float64s",
						Code: `nums := []int{3, 1, 4, 1, 5, 9}
sort.Ints(nums)
fmt.Println(nums)                     // [1 1 3 4 5 9]
fmt.Println(sort.IntsAreSorted(nums)) // true`,
						Output: `[1 1 3 4 5 9]
true
`,
					},
				},
			},
			{
				Title: "Binary search",
				Examples: []Example{
					{
						Title: "SearchInts / Search",
						Notes: "Returns the index where x would be inserted to keep the slice sorted. Check if s[i] actually equals x.",
						Code: `s := []int{1, 3, 5, 7}
i := sort.SearchInts(s, 5)
fmt.Println(i, s[i] == 5)   // 2 true

i = sort.SearchInts(s, 4)
fmt.Println(i, i < len(s) && s[i] == 4)   // 2 false`,
					},
				},
			},
			{
				Title: "sort.Interface — when you need full control",
				Description: "For non-slice types or multi-key sorts too complex for sort.Slice. Implement Len, Less, Swap.",
				Examples: []Example{
					{
						Title: "Minimal custom type",
						Code: `type byLen []string
func (b byLen) Len() int           { return len(b) }
func (b byLen) Less(i, j int) bool { return len(b[i]) < len(b[j]) }
func (b byLen) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

words := []string{"banana", "fig", "apple"}
sort.Sort(byLen(words))
fmt.Println(words)`,
						Output: `[fig apple banana]
`,
					},
				},
			},
		},
	})
}
