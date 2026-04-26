package content

func init() {
	Register(&Package{
		Name:       "os",
		ImportPath: "os",
		Category:   "I/O & Files",
		Summary:    "Portable OS interface: files, env, args, signals, processes.",
		Sections: []Section{
			{
				Title: "Args, environment, exit",
				Examples: []Example{
					{
						Title: "os.Args — command-line arguments",
						Notes: "os.Args[0] is the program name. For richer parsing use the flag package.",
						Code: `fmt.Println(os.Args)`,
					},
					{
						Title: "Getenv, LookupEnv, Setenv",
						Notes: "LookupEnv distinguishes 'unset' from 'empty'. Prefer it when empty has meaning.",
						Code: `os.Setenv("NAME", "ada")
v := os.Getenv("NAME")              // "ada"
v, ok := os.LookupEnv("MISSING")    // "", false`,
					},
					{
						Title: "os.Exit — stop without running deferred functions",
						Notes: "Exit skips all defers. Prefer returning an error from main's helper to let defers run.",
						Code: `if err := run(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}`,
					},
				},
			},
			{
				Title: "Files — open, create, read, write",
				Examples: []Example{
					{
						Title: "ReadFile and WriteFile — the one-shot helpers",
						Notes: "Use these when the whole file fits in memory. Nothing to close, nothing to flush.",
						Code: `b, err := os.ReadFile("hello.txt")
if err != nil { log.Fatal(err) }
fmt.Println(string(b))

err = os.WriteFile("out.txt", []byte("hi"), 0644)`,
					},
					{
						Title: "os.Open — read-only, streaming",
						Code: `f, err := os.Open("big.log")
if err != nil { log.Fatal(err) }
defer f.Close()

sc := bufio.NewScanner(f)
for sc.Scan() {
    // process line
}`,
					},
					{
						Title: "os.Create — write, truncating any existing file",
						Code: `f, err := os.Create("out.log")
if err != nil { log.Fatal(err) }
defer f.Close()
fmt.Fprintln(f, "hello")`,
					},
					{
						Title: "os.OpenFile — full control over flags and mode",
						Notes: "Combine flags with |. O_APPEND|O_CREATE|O_WRONLY is the classic 'append-only log' pattern.",
						Code: `f, err := os.OpenFile(
    "app.log",
    os.O_APPEND|os.O_CREATE|os.O_WRONLY,
    0644,
)
if err != nil { log.Fatal(err) }
defer f.Close()`,
					},
				},
			},
			{
				Title: "Stat, existence checks",
				Examples: []Example{
					{
						Title: "os.Stat and FileInfo",
						Code: `fi, err := os.Stat("hello.txt")
if err != nil { log.Fatal(err) }
fmt.Println(fi.Name(), fi.Size(), fi.IsDir(), fi.Mode())`,
					},
					{
						Title: "Check for not-exists",
						Notes: "errors.Is(err, os.ErrNotExist) is the idiomatic check — it also matches *PathError wrappers.",
						Code: `_, err := os.Stat("missing")
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("nope")
}`,
					},
				},
			},
			{
				Title: "Directories",
				Examples: []Example{
					{
						Title: "Mkdir vs MkdirAll",
						Notes: "Mkdir errors if parents are missing. MkdirAll is idempotent and creates any missing parents.",
						Code: `os.Mkdir("one", 0755)                   // fails if "one" exists
os.MkdirAll("a/b/c", 0755)               // creates whole chain, no error if present`,
					},
					{
						Title: "ReadDir — list a directory",
						Code: `entries, _ := os.ReadDir(".")
for _, e := range entries {
    fmt.Println(e.Name(), e.IsDir())
}`,
					},
					{
						Title: "Remove vs RemoveAll",
						Notes: "Remove works on a single file or empty dir. RemoveAll recurses. Use with extreme care.",
						Code: `os.Remove("file.txt")
os.RemoveAll("temp/")`,
					},
					{
						Title: "TempDir / CreateTemp / MkdirTemp",
						Code: `dir, _ := os.MkdirTemp("", "myapp-*")
defer os.RemoveAll(dir)

f, _ := os.CreateTemp(dir, "data-*.json")
defer f.Close()
fmt.Println(f.Name())`,
					},
				},
			},
			{
				Title: "Signals",
				Examples: []Example{
					{
						Title: "Notify on SIGINT/SIGTERM for graceful shutdown",
						Code: `ctx, stop := signal.NotifyContext(context.Background(),
    os.Interrupt, syscall.SIGTERM)
defer stop()

<-ctx.Done()
fmt.Println("shutting down...")`,
					},
				},
			},
		},
	})
}
