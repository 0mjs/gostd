package content

func init() {
	Register(&Package{
		Name:       "crypto/mlkem",
		ImportPath: "crypto/mlkem",
		Category:   "Crypto",
		Summary:    "ML-KEM (FIPS 203) — post-quantum key encapsulation. Produces a shared secret; combine with a classical KEM for hybrid TLS (Go 1.24+).",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Generate + Encapsulate + Decapsulate", Code: `// Receiver:
dk, err := mlkem.GenerateKey768()
if err != nil { log.Fatal(err) }
pub := dk.EncapsulationKey().Bytes()

// Sender:
ek, _ := mlkem.NewEncapsulationKey768(pub)
ct, ss1 := ek.Encapsulate()

// Receiver recovers the shared secret:
ss2, _ := dk.Decapsulate(ct)
// ss1 == ss2`},
					{Title: "When to use", Code: `// Pair with X25519 as a hybrid: use HKDF over X25519_secret || mlkem_secret
// so you're safe against both classical and quantum attackers.`},
				},
			},
		},
	})
}
