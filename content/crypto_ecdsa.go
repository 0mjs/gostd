package content

func init() {
	Register(&Package{
		Name:       "crypto/ecdsa",
		ImportPath: "crypto/ecdsa",
		Category:   "Crypto",
		Summary:    "ECDSA signatures (P-256 etc.). More widely interoperable than Ed25519; used in TLS, JWT ES256.",
		Sections: []Section{
			{
				Title: "Sign and verify",
				Examples: []Example{
					{
						Title: "P-256 round-trip",
						Code: `priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

hash := sha256.Sum256([]byte("hello"))
sig, _ := ecdsa.SignASN1(rand.Reader, priv, hash[:])
ok := ecdsa.VerifyASN1(&priv.PublicKey, hash[:], sig)
fmt.Println(ok)`,
					},
				},
			},
		},
	})
}
