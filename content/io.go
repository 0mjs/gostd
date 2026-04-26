package content

func init() {
	Register(&Package{
		Name:       "io",
		ImportPath: "io",
		Category:   "I/O & Files",
		Summary:    "The Reader/Writer interfaces every streaming API is built on — and the helpers to copy, limit, tee, and combine them.",
		Sections: []Section{
			{
				Title: "The core interfaces",
				Description: "Reader and Writer are arguably the most important interfaces in Go. Everything streamable implements them.",
				Examples: []Example{
					{
						Title: "Reader: Read(p []byte) (n int, err error)",
						Notes: "Returns how many bytes were filled into p. May return n > 0 AND err = io.EOF on the last read — always handle bytes before checking err.",
						Code: `r := strings.NewReader("hello")
buf := make([]byte, 3)
for {
    n, err := r.Read(buf)
    if n > 0 {
        fmt.Printf("got %d: %q\n", n, buf[:n])
    }
    if err == io.EOF {
        break
    }
}`,
						Output: `got 3: "hel"
got 2: "lo"
`,
					},
					{
						Title: "Writer: Write(p []byte) (n int, err error)",
						Code: `var b bytes.Buffer  // *bytes.Buffer implements io.Writer
fmt.Fprint(&b, "hi")
b.Write([]byte(" there"))
fmt.Println(b.String())`,
						Output: `hi there
`,
					},
				},
			},
			{
				Title: "Copy, CopyN, CopyBuffer",
				Description: "The bread-and-butter of streaming: move bytes from a Reader to a Writer.",
				Examples: []Example{
					{
						Title: "io.Copy — copy until EOF",
						Notes: "Uses WriterTo / ReaderFrom fast paths when available (e.g., *os.File). Returns bytes copied.",
						Code: `r := strings.NewReader("hello world")
n, _ := io.Copy(os.Stdout, r)
fmt.Printf("\n(%d bytes)\n", n)`,
						Output: `hello world
(11 bytes)
`,
					},
					{
						Title: "io.CopyN — copy exactly N bytes",
						Code: `r := strings.NewReader("hello world")
io.CopyN(os.Stdout, r, 5)`,
						Output: `hello`,
					},
					{
						Title: "io.CopyBuffer — bring your own buffer",
						Notes: "Avoids allocating a 32 KiB temporary buffer per call. Use when Copy is in a hot path.",
						Code: `buf := make([]byte, 4096)
io.CopyBuffer(dst, src, buf)`,
					},
				},
			},
			{
				Title: "Read helpers",
				Examples: []Example{
					{
						Title: "io.ReadAll — read everything into memory",
						Notes: "Convenient but only safe for bounded inputs. For untrusted sources, wrap in io.LimitReader.",
						Code: `b, _ := io.ReadAll(strings.NewReader("abc"))
fmt.Println(string(b))`,
						Output: `abc
`,
					},
					{
						Title: "io.ReadFull — fill a buffer exactly or error",
						Code: `r := strings.NewReader("hello")
buf := make([]byte, 5)
_, err := io.ReadFull(r, buf)
fmt.Println(string(buf), err)`,
						Output: `hello <nil>
`,
					},
					{
						Title: "io.LimitReader — cap the bytes you'll read",
						Notes: "Critical for accepting untrusted input: stops at N bytes even if the source has more.",
						Code: `r := io.LimitReader(strings.NewReader("abcdefg"), 3)
b, _ := io.ReadAll(r)
fmt.Println(string(b))`,
						Output: `abc
`,
					},
				},
			},
			{
				Title: "Combining and adapting streams",
				Examples: []Example{
					{
						Title: "io.MultiReader — concat Readers",
						Code: `r := io.MultiReader(
    strings.NewReader("hello "),
    strings.NewReader("world"),
)
io.Copy(os.Stdout, r)`,
						Output: `hello world`,
					},
					{
						Title: "io.MultiWriter — tee to many Writers",
						Code: `var buf bytes.Buffer
w := io.MultiWriter(os.Stdout, &buf)
fmt.Fprintln(w, "logged")
fmt.Println("captured:", buf.String())`,
						Output: `logged
captured: logged
`,
					},
					{
						Title: "io.TeeReader — observe a stream as it's read",
						Notes: "Every byte read from the returned Reader is also written to the given Writer. Great for hashing or logging as you stream.",
						Code: `var captured bytes.Buffer
src := strings.NewReader("hello")
r := io.TeeReader(src, &captured)
io.Copy(io.Discard, r)
fmt.Println(captured.String())`,
						Output: `hello
`,
					},
					{
						Title: "io.Pipe — in-memory synchronous pipe",
						Notes: "Write in one goroutine, Read in another. Useful when an API wants a Reader but you produce bytes procedurally.",
						Code: `pr, pw := io.Pipe()
go func() {
    defer pw.Close()
    fmt.Fprint(pw, "hello")
}()
io.Copy(os.Stdout, pr)`,
						Output: `hello`,
					},
				},
			},
			{
				Title: "Sinks and sentinels",
				Examples: []Example{
					{
						Title: "io.Discard — /dev/null for Writer",
						Notes: "Use when you need to read a whole stream but don't care about the bytes (e.g., draining an HTTP body before Close).",
						Code: `io.Copy(io.Discard, someReader)`,
					},
					{
						Title: "io.EOF",
						Notes: "The sentinel error Readers return when there's no more data. Not a real error — just end of stream. Compare with ==.",
					},
				},
			},
		},
	})
}
