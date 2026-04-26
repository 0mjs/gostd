package content

func init() {
	Register(&Package{
		Name:       "iter",
		ImportPath: "iter",
		Category:   "Collections",
		Summary:    "Defines Seq and Seq2, the types behind range-over-func (Go 1.23+). Write reusable iterators for any collection.",
		Sections: []Section{
			{
				Title: "The two types",
				Description: "iter.Seq[V] = func(yield func(V) bool). iter.Seq2[K,V] adds a second value. Consumers use range; producers call yield for each item and stop if yield returns false.",
				Examples: []Example{
					{
						Title: "A simple Seq — integers from a to b",
						Code: `func Range(a, b int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for i := a; i < b; i++ {
            if !yield(i) {
                return
            }
        }
    }
}

for n := range Range(0, 5) {
    fmt.Println(n)
}`,
						Output: `0
1
2
3
4
`,
					},
					{
						Title: "A Seq2 — indexed iteration",
						Code: `func Enumerate[T any](s []T) iter.Seq2[int, T] {
    return func(yield func(int, T) bool) {
        for i, v := range s {
            if !yield(i, v) {
                return
            }
        }
    }
}

for i, v := range Enumerate([]string{"a", "b", "c"}) {
    fmt.Println(i, v)
}`,
					},
					{
						Title: "Pull and Pull2 — consume iterators manually",
						Notes: "Use when range is too rigid — e.g., merging two sequences, or adapting to a callback API. You MUST call stop.",
						Code: `next, stop := iter.Pull(Range(0, 3))
defer stop()
for {
    v, ok := next()
    if !ok { break }
    fmt.Println(v)
}`,
					},
				},
			},
			{
				Title: "Where you'll see iter.Seq in stdlib",
				Description: "slices.All / slices.Values / slices.Backward, maps.Keys / maps.Values / maps.All, strings.Lines / strings.Split* (as of 1.24) all return iter.Seq now.",
			},
		},
	})
}
