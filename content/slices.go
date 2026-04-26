package content

func init() {
	Register(&Package{
		Name:       "slices",
		ImportPath: "slices",
		Category:   "Collections",
		Summary:    "Generic slice operations (Go 1.21+). The modern answer to 'does stdlib have a Contains?' — yes.",
		Sections: []Section{
			{
				Title: "Search and membership",
				Examples: []Example{
					{
						Title: "Contains, Index",
						Code: `s := []string{"go", "rust", "ts"}
slices.Contains(s, "rust")  // true
slices.Index(s, "rust")     // 1
slices.Index(s, "zig")      // -1`,
					},
					{
						Title: "ContainsFunc / IndexFunc — predicate variants",
						Code: `nums := []int{2, 4, 6, 9}
slices.ContainsFunc(nums, func(n int) bool { return n%2 == 1 }) // true`,
					},
					{
						Title: "Min, Max, MinFunc, MaxFunc",
						Notes: "Min/Max panic on an empty slice. Use MinFunc/MaxFunc with cmp.Compare-style functions for custom ordering.",
						Code: `slices.Min([]int{3, 1, 2})  // 1
slices.Max([]int{3, 1, 2})  // 3`,
					},
				},
			},
			{
				Title: "Sorting and ordering",
				Examples: []Example{
					{
						Title: "Sort / SortStable / SortFunc",
						Notes: "Sort works on any ordered type. SortFunc takes a cmp function returning -1/0/+1 — use cmp.Compare.",
						Code: `s := []int{3, 1, 4, 1, 5}
slices.Sort(s)
fmt.Println(s)

type P struct{ Name string; Age int }
people := []P{{"A", 40}, {"B", 30}}
slices.SortFunc(people, func(a, b P) int {
    return cmp.Compare(a.Age, b.Age)
})`,
					},
					{
						Title: "BinarySearch",
						Code: `s := []int{1, 3, 5, 7}
i, found := slices.BinarySearch(s, 5)
fmt.Println(i, found)  // 2 true`,
					},
					{
						Title: "IsSorted / IsSortedFunc",
						Code: `slices.IsSorted([]int{1, 2, 3})  // true`,
					},
				},
			},
			{
				Title: "Mutation",
				Description: "These modify the input slice in place and return the updated slice (length may change).",
				Examples: []Example{
					{
						Title: "Insert, Delete, Replace",
						Code: `s := []int{1, 2, 5}
s = slices.Insert(s, 2, 3, 4)       // [1 2 3 4 5]
s = slices.Delete(s, 1, 3)          // remove indices [1,3) → [1 4 5]
s = slices.Replace(s, 0, 1, 9, 8)   // replace [0,1) with 9,8 → [9 8 4 5]`,
					},
					{
						Title: "Reverse",
						Code: `s := []int{1, 2, 3}
slices.Reverse(s)   // [3 2 1]`,
					},
					{
						Title: "Compact / CompactFunc — dedupe adjacent",
						Notes: "Only removes *adjacent* duplicates — Sort first if you want full dedup.",
						Code: `s := []int{1, 1, 2, 3, 3, 3, 4}
s = slices.Compact(s)  // [1 2 3 4]`,
					},
				},
			},
			{
				Title: "Copying and equality",
				Examples: []Example{
					{
						Title: "Equal / EqualFunc",
						Code: `slices.Equal([]int{1, 2}, []int{1, 2})  // true`,
					},
					{
						Title: "Clone — shallow copy",
						Code: `a := []int{1, 2, 3}
b := slices.Clone(a)
b[0] = 99
fmt.Println(a, b)  // [1 2 3] [99 2 3]`,
					},
					{
						Title: "Concat (1.22+)",
						Code: `x := slices.Concat([]int{1, 2}, []int{3}, []int{4, 5})
fmt.Println(x)  // [1 2 3 4 5]`,
					},
				},
			},
			{
				Title: "Iteration (1.23+)",
				Examples: []Example{
					{
						Title: "All, Values, Backward — range-over-func",
						Code: `for i, v := range slices.All([]string{"a", "b", "c"}) {
    fmt.Println(i, v)
}
for v := range slices.Backward([]int{1, 2, 3}) {
    fmt.Println(v)
}`,
					},
				},
			},
		},
	})
}
