package content

func init() {
	Register(&Package{
		Name:       "container/list",
		ImportPath: "container/list",
		Category:   "Containers",
		Summary:    "Doubly-linked list. Useful when you need to insert/remove nodes in O(1) given an element pointer — LRU caches, etc.",
		Sections: []Section{
			{
				Title: "Basic operations",
				Examples: []Example{
					{
						Title: "Push, iterate, remove",
						Code: `l := list.New()
l.PushBack("a")
l.PushBack("b")
e := l.PushFront("first")
l.InsertAfter("second", e)

for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}
l.Remove(e)`,
					},
				},
			},
			{
				Title: "LRU cache sketch",
				Description: "A map[K]*list.Element + a list gives you O(1) get/put with O(1) eviction of the tail. This is the canonical use.",
			},
		},
	})
}
