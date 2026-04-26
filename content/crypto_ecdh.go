package content

func init() {
	Register(&Package{
		Name:       "crypto/ecdh",
		ImportPath: "crypto/ecdh",
		Category:   "Crypto",
		Summary:    "Elliptic-curve Diffie-Hellman key exchange. The modern, misuse-resistant replacement for crypto/elliptic's low-level API.",
		Sections: []Section{
			{
				Title: "Key agreement",
				Examples: []Example{
					{
						Title: "Two parties derive the same shared secret",
						Code: `curve := ecdh.X25519()
aliceKey, _ := curve.GenerateKey(rand.Reader)
bobKey, _ := curve.GenerateKey(rand.Reader)

aliceShared, _ := aliceKey.ECDH(bobKey.PublicKey())
bobShared, _   := bobKey.ECDH(aliceKey.PublicKey())
// aliceShared == bobShared — feed to HKDF to derive session keys`,
					},
				},
			},
		},
	})
}
