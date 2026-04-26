package content

func init() {
	Register(&Package{
		Name:       "testing/quick",
		ImportPath: "testing/quick",
		Category:   "Testing",
		Summary:    "Quickcheck-style property tests. Largely superseded by native fuzzing (go test -fuzz).",
		Sections: []Section{
			{
				Title: "Property check",
				Examples: []Example{
					{Title: "Check", Code: `prop := func(xs []int) bool {
    return len(Reverse(Reverse(xs))) == len(xs)
}
if err := quick.Check(prop, nil); err != nil {
    t.Error(err)
}`},
					{Title: "CheckEqual", Code: `if err := quick.CheckEqual(fast, slow, nil); err != nil {
    t.Error(err)
}`},
				},
			},
		},
	})
}
