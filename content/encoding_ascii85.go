package content

func init() {
	Register(&Package{
		Name:       "encoding/ascii85",
		ImportPath: "encoding/ascii85",
		Category:   "Encoding",
		Summary:    "Ascii85 (aka btoa) — 4 bytes become 5 printable chars. More compact than base64, much rarer.",
		Sections: []Section{
			{
				Title: "Encode / Decode",
				Examples: []Example{
					{
						Title: "One-shot Encode",
						Code: `src := []byte("Man is distinguished, not only by his reason")
dst := make([]byte, ascii85.MaxEncodedLen(len(src)))
n := ascii85.Encode(dst, src)
fmt.Println(string(dst[:n]))`,
					},
				},
			},
		},
	})
}
