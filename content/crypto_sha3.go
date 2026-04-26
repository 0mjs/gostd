package content

func init() {
	Register(&Package{
		Name:       "crypto/sha3",
		ImportPath: "crypto/sha3",
		Category:   "Crypto",
		Summary:    "Keccak-based hashes: SHA3-256, SHA3-512, SHAKE128, SHAKE256. Different construction than SHA-2.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "Sum256",
						Code: `sum := sha3.Sum256([]byte("hello"))
fmt.Printf("%x\n", sum)`,
					},
				},
			},
		},
	})
}
