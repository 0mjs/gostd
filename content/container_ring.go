package content

func init() {
	Register(&Package{
		Name:       "container/ring",
		ImportPath: "container/ring",
		Category:   "Containers",
		Summary:    "Circular linked list. Fixed-size ring buffer where advancing past the end loops to the front. Rarely needed.",
		Sections: []Section{
			{
				Title: "Create and walk",
				Examples: []Example{
					{
						Title: "Initialize a ring of 5 and walk it",
						Code: `r := ring.New(5)
for i := 0; i < r.Len(); i++ {
    r.Value = i
    r = r.Next()
}
r.Do(func(x any) { fmt.Println(x) })`,
					},
				},
			},
		},
	})
}
