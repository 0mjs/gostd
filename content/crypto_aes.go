package content

func init() {
	Register(&Package{
		Name:       "crypto/aes",
		ImportPath: "crypto/aes",
		Category:   "Crypto",
		Summary:    "The AES block cipher. Creates a cipher.Block you combine with a mode from crypto/cipher (GCM, CBC, CTR). Never use AES-ECB.",
		Sections: []Section{
			{
				Title: "Construct the block cipher",
				Examples: []Example{
					{
						Title: "NewCipher — key must be 16, 24, or 32 bytes",
						Code: `key := make([]byte, 32)   // AES-256
rand.Read(key)

block, err := aes.NewCipher(key)
if err != nil { log.Fatal(err) }`,
					},
				},
			},
			{
				Title: "Pair with AES-GCM for authenticated encryption",
				Description: "Use AES-GCM (via crypto/cipher.NewGCM) unless you have an unusual reason not to. See the crypto/cipher page for the full example.",
			},
		},
	})
}
