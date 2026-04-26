package content

func init() {
	Register(&Package{
		Name:       "strconv",
		ImportPath: "strconv",
		Category:   "Formatting & Strings",
		Summary:    "Convert between strings and numeric/boolean types. Lower-level and faster than fmt for pure conversion.",
		Sections: []Section{
			{
				Title: "Integers",
				Examples: []Example{
					{
						Title: "Atoi / Itoa — decimal shortcut",
						Notes: "Atoi is Ascii-to-integer. Use when you have a plain base-10 int-sized value.",
						Code: `n, err := strconv.Atoi("42")
fmt.Println(n, err)

s := strconv.Itoa(-42)
fmt.Println(s)`,
						Output: `42 <nil>
-42
`,
					},
					{
						Title: "ParseInt / FormatInt — any base, any bit size",
						Notes: "base=0 lets the prefix decide (0x=16, 0o=8, 0b=2, else 10). bitSize is the width to fit into.",
						Code: `n, _ := strconv.ParseInt("0xff", 0, 64)   // 255
m, _ := strconv.ParseInt("-10", 10, 8)    // -10 (fits int8)
_, err := strconv.ParseInt("300", 10, 8)  // value out of range

s := strconv.FormatInt(255, 16)           // "ff"
fmt.Println(n, m, err, s)`,
						Output: `255 -10 strconv.ParseInt: parsing "300": value out of range ff
`,
					},
					{
						Title: "ParseUint / FormatUint",
						Code: `u, _ := strconv.ParseUint("1010", 2, 64)  // 10
fmt.Println(u, strconv.FormatUint(10, 2)) // "1010"`,
					},
				},
			},
			{
				Title: "Floats",
				Examples: []Example{
					{
						Title: "ParseFloat — 32 or 64 bit",
						Code: `f, _ := strconv.ParseFloat("3.14", 64)
fmt.Println(f)`,
						Output: `3.14
`,
					},
					{
						Title: "FormatFloat — formats 'g', 'f', 'e', etc.",
						Notes: "prec=-1 means the smallest number of digits that round-trips exactly.",
						Code: `fmt.Println(strconv.FormatFloat(3.14159, 'f', 2, 64))  // 3.14
fmt.Println(strconv.FormatFloat(1e10, 'e', -1, 64))    // 1e+10
fmt.Println(strconv.FormatFloat(0.1, 'g', -1, 64))     // 0.1`,
						Output: `3.14
1e+10
0.1
`,
					},
				},
			},
			{
				Title: "Bool",
				Examples: []Example{
					{
						Title: "ParseBool / FormatBool",
						Notes: `ParseBool accepts 1, t, T, TRUE, true, True and the false equivalents.`,
						Code: `b, _ := strconv.ParseBool("true")
fmt.Println(b, strconv.FormatBool(b))`,
						Output: `true true
`,
					},
				},
			},
			{
				Title: "Quoting",
				Examples: []Example{
					{
						Title: "Quote / Unquote — Go-syntax strings",
						Code: `fmt.Println(strconv.Quote("hi\tthere"))      // "hi\tthere"
s, _ := strconv.Unquote(` + "`\"hi\\tthere\"`" + `)
fmt.Println(s)                                // hi	there`,
						Output: `"hi\tthere"
hi	there
`,
					},
					{
						Title: "QuoteRune and AppendQuote — avoid allocations",
						Notes: "Append* variants write into an existing []byte — great inside hot loops.",
						Code: `buf := []byte("msg=")
buf = strconv.AppendQuote(buf, "hi\n")
fmt.Println(string(buf))`,
						Output: `msg="hi\n"
`,
					},
				},
			},
		},
	})
}
