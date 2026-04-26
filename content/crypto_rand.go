package content

func init() {
	Register(&Package{
		Name:       "crypto/rand",
		ImportPath: "crypto/rand",
		Category:   "Crypto",
		Summary:    "Cryptographically-secure random bytes. Reach for this for tokens, keys, nonces. NEVER math/rand for security.",
		Sections: []Section{
			{
				Title: "Random bytes",
				Examples: []Example{
					{
						Title: "Fill a buffer",
						Code: `buf := make([]byte, 32)
_, err := rand.Read(buf)
if err != nil { log.Fatal(err) }
fmt.Printf("%x\n", buf)`,
					},
					{
						Title: "Build a session token",
						Code: `b := make([]byte, 32)
rand.Read(b)
token := base64.RawURLEncoding.EncodeToString(b)`,
					},
				},
			},
			{
				Title: "Bounded integers",
				Examples: []Example{
					{
						Title: "Int — uniform random *big.Int in [0, max)",
						Code: `max := big.NewInt(1_000_000)
n, _ := rand.Int(rand.Reader, max)
fmt.Println(n)`,
					},
					{
						Title: "Text (Go 1.24+) — random ASCII token",
						Code: `fmt.Println(rand.Text())   // "Hk8v9Q5xA3qUJ..." (26 chars default)`,
					},
				},
			},
			{
				Title: "rand.Reader",
				Description: "An io.Reader you can pass anywhere a random source is asked for — RSA key gen, TLS cert creation, etc.",
			},
		},
	})
}
