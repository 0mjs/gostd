package content

func init() {
	Register(&Package{
		Name:       "crypto/hmac",
		ImportPath: "crypto/hmac",
		Category:   "Crypto",
		Summary:    "Keyed Message Authentication Code. Pair with any Hash (SHA-256 usually) to authenticate messages.",
		Sections: []Section{
			{
				Title: "Sign and verify",
				Examples: []Example{
					{
						Title: "Sign",
						Code: `mac := hmac.New(sha256.New, []byte("secret"))
mac.Write([]byte("message"))
sig := mac.Sum(nil)
fmt.Printf("%x\n", sig)`,
					},
					{
						Title: "Verify — use hmac.Equal, not bytes.Equal",
						Notes: "hmac.Equal is constant-time. Plain comparison leaks timing info that attackers can exploit.",
						Code: `mac := hmac.New(sha256.New, key)
mac.Write(message)
expected := mac.Sum(nil)

if !hmac.Equal(provided, expected) {
    return errors.New("bad signature")
}`,
					},
				},
			},
			{
				Title: "Common uses",
				Description: "Cookie signing, API signed URLs (AWS-style), webhook signatures (GitHub, Stripe), JWT HS256.",
			},
		},
	})
}
