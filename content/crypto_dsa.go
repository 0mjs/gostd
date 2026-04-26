package content

func init() {
	Register(&Package{
		Name:       "crypto/dsa",
		ImportPath: "crypto/dsa",
		Category:   "Crypto",
		Summary:    "DSA signatures. Deprecated. Use crypto/ed25519 or crypto/ecdsa instead.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Why not DSA", Code: `// Present for parsing old keys and legacy interop.
// Do not generate new DSA keys in 2025 — use Ed25519 or ECDSA (P-256).`},
				},
			},
		},
	})
}
