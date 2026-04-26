package content

func init() {
	Register(&Package{
		Name:       "strings",
		ImportPath: "strings",
		Category:   "Formatting & Strings",
		Summary:    "Search, split, join, replace, case conversion. Works on UTF-8-encoded strings.",
		Sections: []Section{
			{
				Title: "Searching and testing",
				Examples: []Example{
					{
						Title: "Contains, HasPrefix, HasSuffix",
						Code: `strings.Contains("gopher", "ph")    // true
strings.HasPrefix("gopher", "go")   // true
strings.HasSuffix("gopher", "her")  // true
strings.ContainsAny("abc", "xyzb")  // true — any of the runes
strings.ContainsRune("abc", 'b')    // true`,
					},
					{
						Title: "Index family — position of first match, or -1",
						Notes: "Pair Index with slicing when you need to split on something custom.",
						Code: `strings.Index("gopher", "ph")       // 2
strings.LastIndex("banana", "a")    // 5
strings.IndexAny("abc", "xyb")      // 1 — first rune from set
strings.IndexByte("abc", 'c')       // 2`,
					},
					{
						Title: "Count",
						Code: `strings.Count("cheese", "e")  // 3
strings.Count("abc", "")      // 4 — len(s) + 1 for empty substring`,
					},
					{
						Title: "EqualFold — case-insensitive equality",
						Notes: "Prefer EqualFold over lowercasing both sides — it's cheaper and Unicode-correct.",
						Code: `strings.EqualFold("Go", "GO")  // true`,
					},
				},
			},
			{
				Title: "Split and join",
				Examples: []Example{
					{
						Title: "Split, SplitN, SplitAfter",
						Notes: "Split drops the separator; SplitAfter keeps it. SplitN caps the number of pieces.",
						Code: `strings.Split("a,b,c", ",")          // ["a" "b" "c"]
strings.SplitN("a,b,c", ",", 2)      // ["a" "b,c"]
strings.SplitAfter("a,b,c", ",")     // ["a," "b," "c"]
strings.Fields("  hello   world  ")  // ["hello" "world"] — any whitespace`,
					},
					{
						Title: "Join",
						Code: `strings.Join([]string{"a", "b", "c"}, ", ")  // "a, b, c"`,
					},
					{
						Title: "Cut — idiomatic single split",
						Notes: "Cut is the modern (1.18+) way to split on a separator once. Cleaner than SplitN(2).",
						Code: `k, v, ok := strings.Cut("name=ada", "=")
fmt.Println(k, v, ok)
k, v, ok = strings.Cut("bare", "=")
fmt.Println(k, v, ok)  // "bare", "", false`,
						Output: `name ada true
bare  false
`,
					},
					{
						Title: "CutPrefix / CutSuffix (1.20+)",
						Code: `s, ok := strings.CutPrefix("go-mod", "go-")  // "mod", true
s, ok = strings.CutSuffix("app.log", ".log") // "app", true`,
					},
				},
			},
			{
				Title: "Transformation",
				Examples: []Example{
					{
						Title: "Case conversion",
						Code: `strings.ToUpper("Go")        // "GO"
strings.ToLower("Go")        // "go"
strings.Title("hello world") // DEPRECATED — use cases package
strings.ToValidUTF8("a\xffb", "?")  // "a?b"`,
					},
					{
						Title: "TrimSpace, Trim, TrimLeft/Right",
						Notes: "Trim removes any runes in the cutset, not a substring. Use TrimPrefix/TrimSuffix for exact matches.",
						Code: `strings.TrimSpace("  hi\n")        // "hi"
strings.Trim("--hi--", "-")        // "hi"
strings.TrimLeft("000123", "0")    // "123"
strings.TrimPrefix("go-mod", "go-") // "mod"
strings.TrimSuffix("a.log", ".log") // "a"`,
					},
					{
						Title: "Replace and ReplaceAll",
						Code: `strings.Replace("aaaa", "a", "b", 2)  // "bbaa"
strings.ReplaceAll("aaaa", "a", "b")  // "bbbb"`,
					},
					{
						Title: "Map — per-rune transformation",
						Code: `rot13 := func(r rune) rune {
    switch {
    case r >= 'A' && r <= 'Z':
        return 'A' + (r-'A'+13)%26
    case r >= 'a' && r <= 'z':
        return 'a' + (r-'a'+13)%26
    }
    return r
}
strings.Map(rot13, "Hello, Gophers!")  // "Uryyb, Tbcuref!"`,
					},
					{
						Title: "Repeat",
						Code: `strings.Repeat("ab", 3)  // "ababab"`,
					},
				},
			},
			{
				Title: "Building strings efficiently",
				Description: "Concatenation with + allocates each time. For loops of writes, use strings.Builder.",
				Examples: []Example{
					{
						Title: "strings.Builder — zero-copy final string",
						Notes: "Builder writes into a growing []byte and hands it out as a string without copying. Don't copy Builder values after first use.",
						Code: `var b strings.Builder
for i := 0; i < 3; i++ {
    fmt.Fprintf(&b, "line %d\n", i)
}
fmt.Print(b.String())`,
						Output: `line 0
line 1
line 2
`,
					},
					{
						Title: "Reader — read from a string",
						Notes: "strings.NewReader gives you an io.Reader/Seeker without a copy. Useful for APIs that take readers.",
						Code: `r := strings.NewReader("hello")
buf := make([]byte, 3)
n, _ := r.Read(buf)
fmt.Printf("%d %s\n", n, buf[:n])`,
						Output: `3 hel
`,
					},
				},
			},
			{
				Title: "Replacer — multiple replacements, one pass",
				Examples: []Example{
					{
						Title: "NewReplacer",
						Notes: "Much faster than chained ReplaceAll calls when you have many pairs. Safe to reuse.",
						Code: `r := strings.NewReplacer("<", "&lt;", ">", "&gt;", "&", "&amp;")
fmt.Println(r.Replace("<b>A & B</b>"))`,
						Output: `&lt;b&gt;A &amp; B&lt;/b&gt;
`,
					},
				},
			},
		},
	})
}
