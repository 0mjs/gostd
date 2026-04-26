package content

func init() {
	Register(&Package{
		Name:       "hash/fnv",
		ImportPath: "hash/fnv",
		Category:   "Hashing",
		Summary:    "Fowler–Noll–Vo non-cryptographic hash. Very simple, decent distribution — good for hash tables and bloom filters.",
		Sections: []Section{
			{
				Title: "Construct a hasher",
				Examples: []Example{
					{
						Title: "32-bit, 64-bit, and 128-bit variants",
						Notes: "The 'a' variants (New32a, New64a) scramble better and are the default recommendation.",
						Code: `h := fnv.New64a()
h.Write([]byte("hello"))
fmt.Printf("%016x\n", h.Sum64())`,
					},
				},
			},
		},
	})
}
