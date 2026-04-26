package content

func init() {
	Register(&Package{
		Name:       "crypto/hkdf",
		ImportPath: "crypto/hkdf",
		Category:   "Crypto",
		Summary:    "HMAC-based Key Derivation Function. Turn a shared secret (from ECDH, say) into one or more strong keys.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "Derive a 32-byte key",
						Code: `key, err := hkdf.Key(sha256.New, sharedSecret, salt, []byte("info"), 32)
if err != nil { log.Fatal(err) }`,
					},
				},
			},
		},
	})
}
