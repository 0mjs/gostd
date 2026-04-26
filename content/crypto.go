package content

func init() {
	Register(&Package{
		Name:       "crypto",
		ImportPath: "crypto",
		Category:   "Crypto",
		Summary:    "The root crypto package. Defines Hash (an enum of registered hash funcs), the Signer and Decrypter interfaces, and little else. Real work happens in subpackages.",
		Sections: []Section{
			{
				Title: "Hash — an enum of algorithms",
				Description: "Subpackages (crypto/sha256, crypto/sha512, crypto/md5...) register themselves on import. crypto.Hash lets generic code name them.",
				Examples: []Example{
					{
						Title: "Pick a hash by name",
						Code: `h := crypto.SHA256.New()   // requires import of crypto/sha256
io.Copy(h, f)
fmt.Printf("%x\n", h.Sum(nil))`,
					},
				},
			},
			{
				Title: "Signer interface",
				Description: "crypto.Signer is the abstraction used by TLS, ssh, and PKI code: anything with a private key you can sign with. *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey all implement it.",
			},
		},
	})
}
