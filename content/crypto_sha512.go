package content

func init() {
	Register(&Package{
		Name:       "crypto/sha512",
		ImportPath: "crypto/sha512",
		Category:   "Crypto",
		Summary:    "SHA-384 and SHA-512 hashes. Same API shape as crypto/sha256.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "New, Sum512, Sum384",
						Code: `sum := sha512.Sum512([]byte("hello"))
fmt.Printf("%x\n", sum)`,
					},
				},
			},
		},
	})
}
