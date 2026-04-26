package content

func init() {
	Register(&Package{
		Name:       "crypto/rc4",
		ImportPath: "crypto/rc4",
		Category:   "Crypto",
		Summary:    "RC4 stream cipher. BROKEN. Only for reading legacy protocol data you cannot avoid.",
		Sections: []Section{
			{
				Title: "Usage (legacy only)",
				Examples: []Example{
					{Title: "NewCipher", Code: `c, err := rc4.NewCipher(key)
if err != nil { log.Fatal(err) }
out := make([]byte, len(in))
c.XORKeyStream(out, in)
// Do NOT use for anything new. Use AES-GCM or ChaCha20-Poly1305.`},
				},
			},
		},
	})
}
