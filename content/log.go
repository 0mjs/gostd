package content

func init() {
	Register(&Package{
		Name:       "log",
		ImportPath: "log",
		Category:   "Errors & Logging",
		Summary:    "Classic unstructured logger. Fine for small tools; prefer log/slog for structured logs in new code.",
		Sections: []Section{
			{
				Title: "Default logger",
				Examples: []Example{
					{Title: "Print / Printf / Println", Code: `log.Println("server started on :8080")
log.Printf("user %s logged in", name)`},
					{Title: "Fatal exits with status 1", Code: `log.Fatal("cannot open db:", err)     // log + os.Exit(1)
log.Fatalf("bad config: %v", err)`},
					{Title: "Panic logs and panics", Code: `log.Panic("invariant violated")`},
				},
			},
			{
				Title: "Custom logger",
				Examples: []Example{
					{Title: "New", Code: `l := log.New(os.Stderr, "api: ", log.LstdFlags|log.Lshortfile)
l.Println("hello")   // api: 2024/01/02 15:04:05 main.go:42: hello`},
					{Title: "SetFlags / SetPrefix on default", Code: `log.SetFlags(log.LstdFlags | log.Lmicroseconds)
log.SetPrefix("worker: ")`},
					{Title: "SetOutput", Code: `f, _ := os.Create("app.log")
log.SetOutput(f)`},
				},
			},
		},
	})
}
