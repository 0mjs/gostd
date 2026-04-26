package content

func init() {
	Register(&Package{
		Name:       "crypto/elliptic",
		ImportPath: "crypto/elliptic",
		Category:   "Crypto",
		Summary:    "Named NIST curves (P-224, P-256, P-384, P-521). Most new code uses crypto/ecdsa or crypto/ecdh instead — this is the lower-level curve API.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Pick a curve", Code: `c := elliptic.P256()
// Pass into ecdsa.GenerateKey / ecdh.P256(), etc.`},
					{Title: "Point encoding", Code: `// Marshal/Unmarshal is deprecated in new code;
// use crypto/ecdh for exchange and crypto/ecdsa for signatures.`},
				},
			},
		},
	})
}
