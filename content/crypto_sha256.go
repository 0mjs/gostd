package content

func init() {
	Register(&Package{
		Name:       "crypto/sha256",
		ImportPath: "crypto/sha256",
		Category:   "Crypto",
		Summary:    "SHA-256 and SHA-224 hashing. The exemplar for every hash package in stdlib (md5, sha1, sha512 follow the same API).",
		Sections: []Section{
			{
				Title: "One-shot: Sum256",
				Examples: []Example{
					{
						Title: "Hash a []byte in one call",
						Code: `sum := sha256.Sum256([]byte("hello"))
fmt.Printf("%x\n", sum)`,
						Output: `2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
`,
					},
				},
			},
			{
				Title: "Streaming: the Hash interface",
				Description: "sha256.New returns a hash.Hash that implements io.Writer. Feed any stream into it — file, HTTP body, concatenated inputs — then call Sum.",
				Examples: []Example{
					{
						Title: "Hash a file without reading it into memory",
						Code: `f, _ := os.Open("big.iso")
defer f.Close()

h := sha256.New()
if _, err := io.Copy(h, f); err != nil {
    log.Fatal(err)
}
fmt.Printf("%x\n", h.Sum(nil))`,
					},
					{
						Title: "Multi-part hashing",
						Notes: "Write each part. h.Sum(b) appends to b without resetting the hash — pass nil for a fresh digest.",
						Code: `h := sha256.New()
h.Write([]byte("hello "))
h.Write([]byte("world"))
fmt.Printf("%x\n", h.Sum(nil))`,
					},
				},
			},
			{
				Title: "HMAC with crypto/hmac",
				Examples: []Example{
					{
						Title: "HMAC-SHA256",
						Notes: "Never use raw SHA-256 for authentication — use HMAC with a secret key.",
						Code: `mac := hmac.New(sha256.New, []byte("secret"))
mac.Write([]byte("message"))
sig := mac.Sum(nil)
fmt.Printf("%x\n", sig)`,
					},
					{
						Title: "hmac.Equal — constant-time comparison",
						Notes: "bytes.Equal and == leak timing info. hmac.Equal doesn't. ALWAYS use it for comparing signatures.",
						Code: `if !hmac.Equal(want, got) {
    return errors.New("bad signature")
}`,
					},
				},
			},
		},
	})
}
