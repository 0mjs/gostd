package content

func init() {
	Register(&Package{
		Name:       "unicode/utf16",
		ImportPath: "unicode/utf16",
		Category:   "Formatting & Strings",
		Summary:    "Encode and decode UTF-16 (surrogate pairs). Rarely needed in Go — mostly useful when crossing into Windows APIs or JS.",
		Sections: []Section{
			{
				Title: "Encode and Decode",
				Examples: []Example{
					{
						Title: "Encode — runes → uint16 code units",
						Code: `runes := []rune{'h', 'i', '🎉'}
units := utf16.Encode(runes)
fmt.Println(units)`,
						Output: `[104 105 55356 57225]
`,
					},
					{
						Title: "Decode — uint16 code units → runes",
						Code: `runes := utf16.Decode([]uint16{104, 105, 55356, 57225})
fmt.Println(string(runes))`,
						Output: `hi🎉
`,
					},
				},
			},
			{
				Title: "Surrogate helpers",
				Examples: []Example{
					{
						Title: "IsSurrogate, EncodeRune, DecodeRune",
						Code: `hi, lo := utf16.EncodeRune('🎉')
r := utf16.DecodeRune(hi, lo)
fmt.Printf("%U %U -> %c\n", hi, lo, r)`,
					},
				},
			},
		},
	})
}
