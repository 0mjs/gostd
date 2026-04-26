package content

func init() {
	Register(&Package{
		Name:       "crypto/rsa",
		ImportPath: "crypto/rsa",
		Category:   "Crypto",
		Summary:    "RSA signing, OAEP encryption, PSS signatures. Use 2048+ bit keys; prefer Ed25519/ECDSA for new systems.",
		Sections: []Section{
			{
				Title: "Generate a key",
				Examples: []Example{
					{
						Title: "GenerateKey",
						Code: `key, err := rsa.GenerateKey(rand.Reader, 2048)
if err != nil { log.Fatal(err) }`,
					},
				},
			},
			{
				Title: "Signing and encryption",
				Examples: []Example{
					{
						Title: "Sign with PSS (preferred over PKCS#1 v1.5)",
						Code: `hash := sha256.Sum256(msg)
sig, _ := rsa.SignPSS(rand.Reader, key, crypto.SHA256, hash[:], nil)
err := rsa.VerifyPSS(&key.PublicKey, crypto.SHA256, hash[:], sig, nil)`,
					},
					{
						Title: "Encrypt with OAEP",
						Code: `ct, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, &key.PublicKey, msg, nil)
pt, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, ct, nil)`,
					},
				},
			},
		},
	})
}
