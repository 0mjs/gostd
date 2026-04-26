package content

func init() {
	Register(&Package{
		Name:       "crypto/sha1",
		ImportPath: "crypto/sha1",
		Category:   "Crypto",
		Summary:    "SHA-1 hashing. BROKEN for security. Remains for Git-compatible hashing and legacy protocols.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "Sum",
						Code: `sum := sha1.Sum([]byte("hello"))
fmt.Printf("%x\n", sum)`,
					},
				},
			},
		},
	})
}
