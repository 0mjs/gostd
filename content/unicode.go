package content

func init() {
	Register(&Package{
		Name:       "unicode",
		ImportPath: "unicode",
		Category:   "Formatting & Strings",
		Summary:    "Rune classification and simple case conversion. The 'is this a letter / digit / space' answers.",
		Sections: []Section{
			{
				Title: "Classifier functions",
				Examples: []Example{
					{
						Title: "IsLetter, IsDigit, IsSpace, IsPunct",
						Code: `unicode.IsLetter('é')  // true
unicode.IsDigit('7')   // true
unicode.IsSpace(' ')   // true
unicode.IsPunct(',')   // true`,
					},
					{
						Title: "IsUpper, IsLower, IsTitle",
						Code: `unicode.IsUpper('A')   // true
unicode.IsLower('a')   // true`,
					},
					{
						Title: "In / Is — check a rune against ranges",
						Notes: "unicode.Is checks one RangeTable; unicode.In accepts many. Ranges like unicode.Latin, unicode.Greek, unicode.Han let you test scripts.",
						Code: `unicode.In('π', unicode.Greek)  // true
unicode.Is(unicode.Han, '汉')    // true`,
					},
				},
			},
			{
				Title: "Case conversion — single rune",
				Examples: []Example{
					{
						Title: "ToUpper, ToLower, ToTitle",
						Notes: "These operate per-rune. For full strings use strings.ToUpper etc.",
						Code: `unicode.ToUpper('ß')  // 'ß' — simple fold
unicode.ToLower('E')  // 'e'`,
					},
				},
			},
		},
	})
}
