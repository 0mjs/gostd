package content

func init() {
	Register(&Package{
		Name:       "encoding/base32",
		ImportPath: "encoding/base32",
		Category:   "Encoding",
		Summary:    "Base32 encoding. Same API shape as encoding/base64 — pick base32 when you need case-insensitive output (e.g., TOTP secrets, some filenames).",
		Sections: []Section{
			{
				Title: "The four encodings",
				Description: "StdEncoding (A–Z2–7, RFC 4648), HexEncoding (0–9A–V, extended hex), plus their Raw* variants without padding.",
				Examples: []Example{
					{
						Title: "EncodeToString / DecodeString",
						Code: `s := base32.StdEncoding.EncodeToString([]byte("hello"))
fmt.Println(s)  // "NBSWY3DP"

b, _ := base32.StdEncoding.DecodeString(s)
fmt.Println(string(b))`,
					},
				},
			},
		},
	})
}
