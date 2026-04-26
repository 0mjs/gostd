package content

func init() {
	Register(&Package{
		Name:       "maps",
		ImportPath: "maps",
		Category:   "Collections",
		Summary:    "Generic map operations (Go 1.21+). The missing helpers for Go's built-in map type.",
		Sections: []Section{
			{
				Title: "Querying",
				Examples: []Example{
					{
						Title: "Keys / Values — iterators (1.23+)",
						Notes: "These now return iter.Seq, not a slice. Use slices.Collect if you need a slice.",
						Code: `m := map[string]int{"a": 1, "b": 2, "c": 3}

for k := range maps.Keys(m) {
    fmt.Println(k)
}

keys := slices.Collect(maps.Keys(m))
slices.Sort(keys)   // map iteration order is random
fmt.Println(keys)`,
					},
					{
						Title: "Equal",
						Code: `maps.Equal(
    map[string]int{"a": 1},
    map[string]int{"a": 1},
) // true`,
					},
				},
			},
			{
				Title: "Mutation",
				Examples: []Example{
					{
						Title: "Clone — shallow copy",
						Code: `m := map[string]int{"a": 1}
n := maps.Clone(m)
n["a"] = 99
fmt.Println(m["a"], n["a"])  // 1 99`,
					},
					{
						Title: "Copy — merge into existing",
						Notes: "Copy(dst, src) overwrites any shared keys in dst with src's values.",
						Code: `a := map[string]int{"x": 1, "y": 2}
b := map[string]int{"y": 20, "z": 30}
maps.Copy(a, b)
fmt.Println(a)  // map[x:1 y:20 z:30]`,
					},
					{
						Title: "DeleteFunc — conditional delete",
						Code: `m := map[string]int{"a": 1, "b": -1, "c": 2}
maps.DeleteFunc(m, func(k string, v int) bool { return v < 0 })`,
					},
				},
			},
			{
				Title: "Building maps from iterators",
				Examples: []Example{
					{
						Title: "Collect — build a map from a Seq2",
						Notes: "The inverse of maps.All. Handy when transforming a slice into a keyed lookup.",
						Code: `pairs := func(yield func(string, int) bool) {
    yield("a", 1)
    yield("b", 2)
}
m := maps.Collect(pairs)
fmt.Println(m)`,
					},
				},
			},
		},
	})
}
