package content

func init() {
	Register(&Package{
		Name:       "hash/adler32",
		ImportPath: "hash/adler32",
		Category:   "Hashing",
		Summary:    "Adler-32 checksum. Used inside zlib. Very fast; weaker than CRC-32 for short inputs.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "Checksum",
						Code: `sum := adler32.Checksum([]byte("hello"))
fmt.Printf("%08x\n", sum)`,
					},
				},
			},
		},
	})
}
