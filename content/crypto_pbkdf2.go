package content

func init() {
	Register(&Package{
		Name:       "crypto/pbkdf2",
		ImportPath: "crypto/pbkdf2",
		Category:   "Crypto",
		Summary:    "Password-Based Key Derivation. Use a high iteration count. For new password hashing prefer Argon2 (golang.org/x/crypto).",
		Sections: []Section{
			{
				Title: "Derive a key from a password",
				Examples: []Example{
					{
						Title: "Key",
						Code: `key, err := pbkdf2.Key(sha256.New, password, salt, 600_000, 32)
if err != nil { log.Fatal(err) }`,
					},
				},
			},
		},
	})
}
