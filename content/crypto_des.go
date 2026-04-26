package content

func init() {
	Register(&Package{
		Name:       "crypto/des",
		ImportPath: "crypto/des",
		Category:   "Crypto",
		Summary:    "DES and triple-DES block ciphers. BROKEN for modern security. Use AES.",
		Sections: []Section{
			{
				Title: "Usage (legacy only)",
				Examples: []Example{
					{Title: "NewCipher / NewTripleDESCipher", Code: `block, err := des.NewCipher(key8)         // DES
block3, err := des.NewTripleDESCipher(key24) // 3DES
_ = block; _ = block3
// Combine with cipher.BlockMode (CBC) or cipher.Stream (CTR) — never raw ECB.`},
				},
			},
		},
	})
}
