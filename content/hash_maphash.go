package content

func init() {
	Register(&Package{
		Name:       "hash/maphash",
		ImportPath: "hash/maphash",
		Category:   "Hashing",
		Summary:    "Fast, randomly-seeded hashing for strings and bytes. The same hash Go's map uses internally — great for building sharded caches.",
		Sections: []Section{
			{
				Title: "Why this exists",
				Description: "Unlike crypto hashes, maphash is deliberately non-portable and per-process random (defeats collision-DoS). Unlike fnv/crc, it's very fast. Use it when you want 'hash a key, pick a shard' — NOT when the hash value needs to survive process restarts.",
			},
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "One-shot helpers (Go 1.19+)",
						Code: `h1 := maphash.String(maphash.MakeSeed(), "hello")
h2 := maphash.Bytes(maphash.MakeSeed(), []byte("hello"))`,
					},
					{
						Title: "Incremental — for composite keys",
						Code: `var h maphash.Hash
h.SetSeed(maphash.MakeSeed())
h.WriteString("user:")
h.WriteString(userID)
sum := h.Sum64()`,
					},
					{
						Title: "Comparable (Go 1.24+) — typed keys",
						Code: `seed := maphash.MakeSeed()
n := maphash.Comparable(seed, struct{ X, Y int }{3, 4})
fmt.Println(n)`,
					},
				},
			},
		},
	})
}
