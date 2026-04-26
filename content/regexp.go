package content

func init() {
	Register(&Package{
		Name:       "regexp",
		ImportPath: "regexp",
		Category:   "Formatting & Strings",
		Summary:    "RE2 regular expressions. Guaranteed linear time — no catastrophic backtracking.",
		Sections: []Section{
			{
				Title: "Compile vs MustCompile",
				Description: "Compile returns (*Regexp, error). MustCompile panics on bad patterns — use it for constant patterns at program start.",
				Examples: []Example{
					{
						Title: "MustCompile — constant pattern",
						Code: `var emailRe = regexp.MustCompile(` + "`" + `^[^@\s]+@[^@\s]+\.[^@\s]+$` + "`" + `)

fmt.Println(emailRe.MatchString("a@b.co"))  // true
fmt.Println(emailRe.MatchString("nope"))    // false`,
						Output: `true
false
`,
					},
					{
						Title: "Compile — pattern from untrusted input",
						Code: `re, err := regexp.Compile(userPattern)
if err != nil {
    return fmt.Errorf("bad pattern: %w", err)
}`,
					},
				},
			},
			{
				Title: "Find family — 16 variants with a pattern",
				Description: "Method names follow: Find[All][String][Submatch][Index]. Think of it as four axes. All = return every match. String = work on strings rather than []byte. Submatch = include capture groups. Index = return offsets instead of matched text.",
				Examples: []Example{
					{
						Title: "FindString / FindAllString",
						Code: `re := regexp.MustCompile(` + "`" + `\d+` + "`" + `)
fmt.Println(re.FindString("abc 12 def 34"))      // "12"
fmt.Println(re.FindAllString("abc 12 def 34", -1)) // ["12" "34"]
fmt.Println(re.FindAllString("a1 b2 c3", 2))      // ["1" "2"] — cap at 2`,
						Output: `12
[12 34]
[1 2]
`,
					},
					{
						Title: "FindStringSubmatch — capture groups",
						Code: `re := regexp.MustCompile(` + "`" + `(\w+)=(\w+)` + "`" + `)
m := re.FindStringSubmatch("name=ada")
// m[0] = whole match, m[1..] = groups
fmt.Println(m[0], m[1], m[2])`,
						Output: `name=ada name ada
`,
					},
					{
						Title: "Named groups + SubexpIndex",
						Code: `re := regexp.MustCompile(` + "`" + `(?P<key>\w+)=(?P<val>\w+)` + "`" + `)
m := re.FindStringSubmatch("name=ada")
fmt.Println(m[re.SubexpIndex("key")], m[re.SubexpIndex("val")])`,
						Output: `name ada
`,
					},
					{
						Title: "FindStringIndex — positions instead of text",
						Code: `re := regexp.MustCompile(` + "`" + `\d+` + "`" + `)
loc := re.FindStringIndex("abc 12 def")
fmt.Println(loc)        // [4 6]
fmt.Println("abc 12 def"[loc[0]:loc[1]])`,
						Output: `[4 6]
12
`,
					},
				},
			},
			{
				Title: "Replace",
				Examples: []Example{
					{
						Title: "ReplaceAllString",
						Code: `re := regexp.MustCompile(` + "`" + `\s+` + "`" + `)
fmt.Println(re.ReplaceAllString("hi   there\tworld", " "))`,
						Output: `hi there world
`,
					},
					{
						Title: "ReplaceAllString with $1 references",
						Code: `re := regexp.MustCompile(` + "`" + `(\w+)@(\w+)` + "`" + `)
fmt.Println(re.ReplaceAllString("ada@example", "$2.$1"))`,
						Output: `example.ada
`,
					},
					{
						Title: "ReplaceAllStringFunc — custom per-match",
						Code: `re := regexp.MustCompile(` + "`" + `\w+` + "`" + `)
out := re.ReplaceAllStringFunc("hi there", strings.ToUpper)
fmt.Println(out)`,
						Output: `HI THERE
`,
					},
				},
			},
			{
				Title: "Split",
				Examples: []Example{
					{
						Title: "Split — splitter as a regex",
						Code: `re := regexp.MustCompile(` + "`" + `\s*,\s*` + "`" + `)
fmt.Println(re.Split("a , b,c ,  d", -1))`,
						Output: `[a b c d]
`,
					},
				},
			},
		},
	})
}
