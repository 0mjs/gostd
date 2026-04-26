package content

func init() {
	Register(&Package{
		Name:       "bufio",
		ImportPath: "bufio",
		Category:   "I/O & Files",
		Summary:    "Buffered I/O on top of io.Reader/Writer. Adds line-reading, large single reads, and a fluent Scanner.",
		Sections: []Section{
			{
				Title: "Why buffer?",
				Description: "Each os.File.Read is a syscall. bufio reads a big chunk ahead and hands you slices out of memory — cheap iteration, same Reader/Writer interface.",
			},
			{
				Title: "Scanner — the line-reading workhorse",
				Examples: []Example{
					{
						Title: "Scan lines (the default)",
						Code: `sc := bufio.NewScanner(strings.NewReader("one\ntwo\nthree\n"))
for sc.Scan() {
    fmt.Println(sc.Text())
}
if err := sc.Err(); err != nil {
    log.Fatal(err)
}`,
						Output: `one
two
three
`,
					},
					{
						Title: "Scan words or runes",
						Code: `sc := bufio.NewScanner(strings.NewReader("the  quick brown"))
sc.Split(bufio.ScanWords)
for sc.Scan() {
    fmt.Println(sc.Text())
}`,
						Output: `the
quick
brown
`,
					},
					{
						Title: "Increase the buffer for long lines",
						Notes: "Default max line is 64 KiB. Feed it your own buffer for longer lines.",
						Code: `sc := bufio.NewScanner(r)
buf := make([]byte, 0, 1024*1024)
sc.Buffer(buf, 10*1024*1024)  // up to 10 MiB per line`,
					},
				},
			},
			{
				Title: "Reader — more flexible than Scanner",
				Description: "Scanner is line-at-a-time. Reader lets you peek, unread, ReadString with any delimiter, and handle binary.",
				Examples: []Example{
					{
						Title: "ReadString — delimiter-based",
						Notes: "Returns the match including the delimiter. Returns err != nil at EOF even with data, so handle both.",
						Code: `r := bufio.NewReader(strings.NewReader("abc,def,"))
for {
    s, err := r.ReadString(',')
    if s != "" {
        fmt.Printf("%q\n", s)
    }
    if err != nil {
        break
    }
}`,
						Output: `"abc,"
"def,"
`,
					},
					{
						Title: "Peek — look without consuming",
						Code: `r := bufio.NewReader(strings.NewReader("hello"))
b, _ := r.Peek(2)
fmt.Println(string(b))
fmt.Println(string(mustRead(r, 5)))`,
						Output: `he
hello
`,
					},
					{
						Title: "ReadByte / UnreadByte",
						Code: `r := bufio.NewReader(strings.NewReader("abc"))
c, _ := r.ReadByte()
r.UnreadByte()
d, _ := r.ReadByte()
fmt.Printf("%c %c\n", c, d)  // 'a' 'a'`,
					},
				},
			},
			{
				Title: "Writer — always Flush",
				Description: "bufio.Writer buffers writes until the buffer is full or you Flush. Forgetting Flush loses data.",
				Examples: []Example{
					{
						Title: "Flush before close",
						Code: `f, _ := os.Create("out.txt")
defer f.Close()
w := bufio.NewWriter(f)
defer w.Flush()      // MUST Flush before f.Close runs
fmt.Fprintln(w, "hello")`,
					},
				},
			},
		},
	})
}
