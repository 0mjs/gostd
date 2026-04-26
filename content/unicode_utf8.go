package content

func init() {
	Register(&Package{
		Name:       "unicode/utf8",
		ImportPath: "unicode/utf8",
		Category:   "Formatting & Strings",
		Summary:    "Work with UTF-8-encoded strings at the rune level: counts, encoding, decoding, validation.",
		Sections: []Section{
			{
				Title: "Counting and validating",
				Description: "len(s) gives bytes, not runes. Use utf8 when you need the rune count.",
				Examples: []Example{
					{
						Title: "RuneCountInString vs len",
						Code: `s := "héllo"
fmt.Println(len(s))                          // 6 — bytes
fmt.Println(utf8.RuneCountInString(s))       // 5 — runes`,
						Output: `6
5
`,
					},
					{
						Title: "ValidString — is this a well-formed UTF-8 string?",
						Code: `utf8.ValidString("hello")       // true
utf8.ValidString("a\xffb")      // false`,
					},
				},
			},
			{
				Title: "Decoding and encoding",
				Examples: []Example{
					{
						Title: "DecodeRuneInString — one rune at a time",
						Notes: "Returns the rune and its byte width. Use for manual rune iteration; the range loop does this for you.",
						Code: `s := "héllo"
for i := 0; i < len(s); {
    r, w := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%c at %d (%d bytes)\n", r, i, w)
    i += w
}`,
						Output: `h at 0 (1 bytes)
é at 1 (2 bytes)
l at 3 (1 bytes)
l at 4 (1 bytes)
o at 5 (1 bytes)
`,
					},
					{
						Title: "EncodeRune — write a rune into a []byte",
						Code: `buf := make([]byte, 4)
n := utf8.EncodeRune(buf, '🎉')
fmt.Printf("%d bytes: % x\n", n, buf[:n])`,
						Output: `4 bytes: f0 9f 8e 89
`,
					},
					{
						Title: "range over a string walks runes for you",
						Notes: "A plain for-range on a string already decodes UTF-8 — no utf8 package needed.",
						Code: `for i, r := range "héllo" {
    fmt.Printf("%d: %c\n", i, r)
}`,
						Output: `0: h
1: é
3: l
4: l
5: o
`,
					},
				},
			},
		},
	})
}
