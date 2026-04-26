package content

func init() {
	Register(&Package{
		Name:       "crypto/ed25519",
		ImportPath: "crypto/ed25519",
		Category:   "Crypto",
		Summary:    "Ed25519 signatures. Small keys, small signatures, constant-time by design. Preferred for new signing systems.",
		Sections: []Section{
			{
				Title: "Generate, sign, verify",
				Examples: []Example{
					{
						Title: "End-to-end",
						Code: `pub, priv, err := ed25519.GenerateKey(rand.Reader)
if err != nil { log.Fatal(err) }

msg := []byte("hello")
sig := ed25519.Sign(priv, msg)
ok := ed25519.Verify(pub, msg, sig)
fmt.Println(ok)   // true`,
					},
					{
						Title: "Keys are just byte slices — easy to store",
						Notes: "PublicKeySize=32, PrivateKeySize=64, SignatureSize=64. Serialize with base64 or hex.",
					},
				},
			},
		},
	})
}
