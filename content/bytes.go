package content

func init() {
	Register(&Package{
		Name:       "bytes",
		ImportPath: "bytes",
		Category:   "Formatting & Strings",
		Summary:    "The strings package, but for []byte. Plus Buffer (a mutable growable byte buffer) and NewReader.",
		Sections: []Section{
			{
				Title: "Search and slice manipulation",
				Description: "Most functions mirror the strings package but take and return []byte.",
				Examples: []Example{
					{
						Title: "Contains, Index, HasPrefix",
						Code: `bytes.Contains([]byte("gopher"), []byte("ph"))   // true
bytes.Index([]byte("gopher"), []byte("ph"))      // 2
bytes.HasPrefix([]byte("gopher"), []byte("go"))  // true`,
					},
					{
						Title: "Split, Fields, Join",
						Code: `bytes.Split([]byte("a,b,c"), []byte(","))  // [[97] [98] [99]]
bytes.Fields([]byte("  hi  there  "))      // [[104 105] [116 104 101 114 101]]
bytes.Join([][]byte{{'a'}, {'b'}}, []byte("-")) // "a-b"`,
					},
					{
						Title: "ToUpper, TrimSpace, Replace",
						Code: `bytes.ToUpper([]byte("Go"))              // "GO"
bytes.TrimSpace([]byte("  hi\n"))        // "hi"
bytes.ReplaceAll([]byte("aaa"), []byte("a"), []byte("b")) // "bbb"`,
					},
					{
						Title: "bytes.Equal — fast byte-slice equality",
						Notes: "Faster than string(a) == string(b) because it skips the allocation. Safe with crypto/subtle's ConstantTimeCompare for secrets.",
						Code: `bytes.Equal([]byte("a"), []byte("a"))  // true`,
					},
				},
			},
			{
				Title: "bytes.Buffer — growable byte buffer",
				Description: "Implements io.Reader and io.Writer. Zero value is ready to use.",
				Examples: []Example{
					{
						Title: "Write and retrieve",
						Code: `var b bytes.Buffer
b.WriteString("hello")
b.WriteByte(' ')
fmt.Fprintf(&b, "%d", 42)
fmt.Println(b.String())   // "hello 42"
fmt.Println(b.Len())      // 8`,
						Output: `hello 42
8
`,
					},
					{
						Title: "Read from a buffer — it's also an io.Reader",
						Code: `b := bytes.NewBufferString("one\ntwo\n")
scanner := bufio.NewScanner(b)
for scanner.Scan() {
    fmt.Println("line:", scanner.Text())
}`,
						Output: `line: one
line: two
`,
					},
					{
						Title: "Reset + reuse — avoid reallocation",
						Notes: "Call Reset to clear the buffer for reuse without throwing the underlying slice away.",
						Code: `var b bytes.Buffer
b.WriteString("first")
b.Reset()
b.WriteString("second")
fmt.Println(b.String())`,
						Output: `second
`,
					},
				},
			},
			{
				Title: "bytes.Reader — treat a []byte as an io.Reader",
				Examples: []Example{
					{
						Title: "NewReader is cheap and supports Seek",
						Code: `r := bytes.NewReader([]byte("hello world"))
r.Seek(6, io.SeekStart)
buf := make([]byte, 5)
r.Read(buf)
fmt.Println(string(buf))`,
						Output: `world
`,
					},
				},
			},
		},
	})
}
