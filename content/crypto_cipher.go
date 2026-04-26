package content

func init() {
	Register(&Package{
		Name:       "crypto/cipher",
		ImportPath: "crypto/cipher",
		Category:   "Crypto",
		Summary:    "Block-cipher modes: GCM, CBC, CTR, CFB, OFB. GCM is what you almost always want.",
		Sections: []Section{
			{
				Title: "AES-GCM — authenticated encryption",
				Description: "GCM encrypts AND authenticates. One call to Seal, one call to Open. Never reuse (key, nonce) — generate a fresh nonce per message.",
				Examples: []Example{
					{
						Title: "Encrypt",
						Code: `block, _ := aes.NewCipher(key)
gcm, _ := cipher.NewGCM(block)

nonce := make([]byte, gcm.NonceSize())
rand.Read(nonce)

ciphertext := gcm.Seal(nonce, nonce, plaintext, additionalData)
// ciphertext prepends the nonce for storage`,
					},
					{
						Title: "Decrypt",
						Code: `nonceSize := gcm.NonceSize()
nonce, ct := ciphertext[:nonceSize], ciphertext[nonceSize:]
plaintext, err := gcm.Open(nil, nonce, ct, additionalData)
if err != nil { return errors.New("invalid ciphertext") }`,
					},
				},
			},
			{
				Title: "CBC, CTR, CFB, OFB — when?",
				Description: "Only when interoperating with a legacy system that requires them. They don't authenticate — you must add HMAC yourself and get the order right (encrypt-then-MAC). Prefer GCM.",
			},
		},
	})
}
