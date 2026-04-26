package content

func init() {
	Register(&Package{
		Name:       "encoding/base64",
		ImportPath: "encoding/base64",
		Category:   "Encoding",
		Summary:    "Base64 encoding. Standard, URL-safe, with or without padding — four ready-made encodings.",
		Sections: []Section{
			{
				Title: "The four built-in encodings",
				Description: "StdEncoding and URLEncoding both pad with =. RawStdEncoding and RawURLEncoding skip padding — best for URLs and JWTs.",
				Examples: []Example{
					{
						Title: "StdEncoding — classic, padded",
						Code: `s := base64.StdEncoding.EncodeToString([]byte("hello!"))
fmt.Println(s)   // aGVsbG8h

raw, _ := base64.StdEncoding.DecodeString(s)
fmt.Println(string(raw))`,
						Output: `aGVsbG8h
hello!
`,
					},
					{
						Title: "URLEncoding — uses - and _ instead of + and /",
						Notes: "Safe in URLs and filenames without further escaping.",
						Code: `base64.URLEncoding.EncodeToString([]byte{0xff, 0xe6})  // "_-Y="`,
					},
					{
						Title: "RawURLEncoding — no padding",
						Notes: "Use this for JWT header/payload/signature segments.",
						Code: `base64.RawURLEncoding.EncodeToString([]byte("hi"))   // "aGk"`,
					},
				},
			},
			{
				Title: "Streaming",
				Examples: []Example{
					{
						Title: "Encoder to any io.Writer",
						Notes: "Must be closed to flush any final bytes.",
						Code: `enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
enc.Write([]byte("hello world"))
enc.Close()   // writes trailing '=' padding
fmt.Println()`,
					},
				},
			},
		},
	})
}
