package content

func init() {
	Register(&Package{
		Name:       "container/heap",
		ImportPath: "container/heap",
		Category:   "Containers",
		Summary:    "Binary heap / priority queue. Implement heap.Interface (which embeds sort.Interface) and the package drives it.",
		Sections: []Section{
			{
				Title: "A min-heap of ints",
				Examples: []Example{
					{
						Title: "Implement the interface",
						Code: `type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

h := &IntHeap{3, 1, 4, 1, 5}
heap.Init(h)
heap.Push(h, 9)
for h.Len() > 0 {
    fmt.Printf("%d ", heap.Pop(h))
}`,
						Output: `1 1 3 4 5 9 `,
					},
				},
			},
			{
				Title: "Priority queue pattern",
				Description: "Wrap your item type in a struct with a Priority int and an index field, and implement Less accordingly. heap.Fix lets you re-prioritize in-place after mutating a field.",
			},
		},
	})
}
