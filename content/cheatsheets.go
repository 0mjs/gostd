package content

// Cheatsheets attach an at-a-glance "TL;DR + most-used patterns" panel to the
// most popular curated packages. They render at the very top of a package page
// — above sections — so an experienced reader can recall the API in a glance,
// and a beginner can copy a working line without wading through prose.
//
// Each cheat row is a real task ("Read a whole file") paired with the smallest
// Go expression that does it. Not exhaustive, intentionally — a cheatsheet
// that lists every signature stops being a cheatsheet.

func init() {
	RegisterCheatsheet("fmt", `Format text and values. Reach for fmt to print, build error messages, render structs while debugging, or scan simple input.`, []CheatRow{
		{"Print a value", `fmt.Println(x)`},
		{"Print to stderr", `fmt.Fprintln(os.Stderr, "warning:", msg)`},
		{"Format into a string", `s := fmt.Sprintf("user=%d", id)`},
		{"Build a wrapped error", `err := fmt.Errorf("open %q: %w", path, err)`},
		{"Show a struct with field names", `fmt.Printf("%+v\n", point)`},
		{"Show Go-syntax (debug)", `fmt.Printf("%#v\n", point)`},
		{"Type of a value", `fmt.Printf("%T\n", x)`},
		{"Hex / binary", `fmt.Printf("%x %b\n", n, n)`},
	})

	RegisterCheatsheet("strings", `Read, search, and reshape UTF-8 text. For mutable byte buffers see bytes; for splitting on patterns see regexp.`, []CheatRow{
		{"Contains a substring", `strings.Contains(s, "needle")`},
		{"Replace all occurrences", `strings.ReplaceAll(s, "old", "new")`},
		{"Split on a separator", `parts := strings.Split(line, ",")`},
		{"Join a slice", `strings.Join(parts, ", ")`},
		{"Trim whitespace", `strings.TrimSpace(s)`},
		{"Lower / upper case", `strings.ToLower(s)`},
		{"Has prefix / suffix", `strings.HasPrefix(s, "http://")`},
		{"Build a long string fast", `var b strings.Builder
b.WriteString("hello, ")
b.WriteString(name)
result := b.String()`},
	})

	RegisterCheatsheet("strconv", `Convert numbers and bools to/from strings. Use this rather than fmt.Sscan for one-off parsing — strconv is faster and gives clearer errors.`, []CheatRow{
		{"String → int", `n, err := strconv.Atoi("42")`},
		{"Int → string", `s := strconv.Itoa(42)`},
		{"String → float", `f, err := strconv.ParseFloat("3.14", 64)`},
		{"Float → string", `s := strconv.FormatFloat(3.14, 'f', -1, 64)`},
		{"String → bool", `b, err := strconv.ParseBool("true")`},
		{"Quote / unquote", `strconv.Quote(s); strconv.Unquote(s)`},
	})

	RegisterCheatsheet("errors", `Build, wrap, and inspect errors. Pair with fmt.Errorf %w to add context while keeping the underlying error inspectable.`, []CheatRow{
		{"Sentinel error", `var ErrNoUser = errors.New("no user")`},
		{"Wrap with context", `return fmt.Errorf("load %s: %w", id, err)`},
		{"Check kind", `if errors.Is(err, ErrNoUser) { ... }`},
		{"Extract typed error", `var pe *fs.PathError
if errors.As(err, &pe) { use(pe.Path) }`},
		{"Combine multiple", `return errors.Join(e1, e2, e3)`},
	})

	RegisterCheatsheet("time", `Wall-clock time, monotonic durations, sleeps, timers, and parsing/formatting using a reference layout.`, []CheatRow{
		{"Now", `t := time.Now()`},
		{"Sleep", `time.Sleep(2 * time.Second)`},
		{"How long since", `elapsed := time.Since(start)`},
		{"Add / subtract", `t.Add(24 * time.Hour)`},
		{"Format ISO-8601 / RFC3339", `t.Format(time.RFC3339)`},
		{"Parse a date", `t, err := time.Parse(time.RFC3339, "2026-04-26T10:00:00Z")`},
		{"Periodic ticker", `tk := time.NewTicker(time.Second)
defer tk.Stop()
for now := range tk.C { use(now) }`},
		{"One-shot timer", `time.AfterFunc(d, func() { ... })`},
	})

	RegisterCheatsheet("context", `Carry deadlines, cancellation, and request-scoped values across goroutines and API boundaries. Pass ctx as the first argument; never store it in a struct.`, []CheatRow{
		{"Root", `ctx := context.Background()`},
		{"Cancel manually", `ctx, cancel := context.WithCancel(parent)
defer cancel()`},
		{"Time-bound", `ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()`},
		{"Check cancellation", `select {
case <-ctx.Done():
    return ctx.Err()
default:
}`},
		{"Carry a value (sparingly)", `ctx = context.WithValue(parent, traceIDKey{}, id)`},
	})

	RegisterCheatsheet("slices", `Generic slice helpers (Go 1.21+). Sort, search, mutate, and convert slices without writing the loop yourself.`, []CheatRow{
		{"Sort ascending", `slices.Sort(nums)`},
		{"Sort by custom key", `slices.SortFunc(users, func(a, b User) int {
    return cmp.Compare(a.Age, b.Age)
})`},
		{"Contains / Index", `slices.Contains(s, v); slices.Index(s, v)`},
		{"Reverse in place", `slices.Reverse(s)`},
		{"Remove dup adjacent (sort first)", `slices.Sort(s)
s = slices.Compact(s)`},
		{"Insert / delete", `s = slices.Insert(s, i, v)
s = slices.Delete(s, i, j)`},
		{"Min / max", `slices.Min(s); slices.Max(s)`},
		{"Clone (avoid sharing backing array)", `cp := slices.Clone(s)`},
	})

	RegisterCheatsheet("io", `Stream-shaped I/O — Reader / Writer interfaces and the helpers that compose them. Most of the rest of the stdlib is built on these.`, []CheatRow{
		{"Read everything from a Reader", `b, err := io.ReadAll(r)`},
		{"Pipe a Reader to a Writer", `_, err := io.Copy(dst, src)`},
		{"Throw bytes away", `io.Copy(io.Discard, r)`},
		{"Cap at N bytes", `r = io.LimitReader(r, 1<<20)`},
		{"Tee (read + write through)", `r = io.TeeReader(r, audit)`},
		{"Chain readers", `r := io.MultiReader(header, body)`},
		{"Multiple writers at once", `w := io.MultiWriter(file, os.Stdout)`},
	})

	RegisterCheatsheet("os", `Talk to the operating system: open files, read env, args, exit, signals, processes.`, []CheatRow{
		{"Read a whole file", `b, err := os.ReadFile("path")`},
		{"Write a whole file", `err := os.WriteFile("path", data, 0644)`},
		{"Open for streaming", `f, err := os.Open("path")
if err != nil { return err }
defer f.Close()`},
		{"Create a new file", `f, err := os.Create("path")`},
		{"Does it exist?", `_, err := os.Stat(path)
missing := errors.Is(err, fs.ErrNotExist)`},
		{"Env var with default", `v := cmp.Or(os.Getenv("PORT"), "8080")`},
		{"CLI args", `args := os.Args[1:]`},
		{"Exit (skips deferred funcs!)", `os.Exit(1)`},
	})

	RegisterCheatsheet("encoding/json", `Encode/decode Go values to JSON. Match wire format with struct tags. For huge documents, prefer the streaming Encoder/Decoder.`, []CheatRow{
		{"Encode to bytes", `b, err := json.Marshal(v)`},
		{"Pretty print", `b, err := json.MarshalIndent(v, "", "  ")`},
		{"Decode from bytes", `err := json.Unmarshal(b, &v)`},
		{"Stream encode to Writer", `err := json.NewEncoder(w).Encode(v)`},
		{"Stream decode from Reader", `err := json.NewDecoder(r).Decode(&v)`},
		{"Optional field tag", `Field string ` + "`" + `json:"field,omitempty"` + "`"},
		{"Skip unknown fields safely", `dec := json.NewDecoder(r)
dec.DisallowUnknownFields()
err := dec.Decode(&v)`},
	})

	RegisterCheatsheet("net/http", `HTTP client + server in the same package. The zero-value http.Client.Get works for quick scripts; for production wire your own Client with a Timeout.`, []CheatRow{
		{"Quick GET", `resp, err := http.Get(url)
if err != nil { return err }
defer resp.Body.Close()
b, _ := io.ReadAll(resp.Body)`},
		{"Client with timeout", `c := &http.Client{Timeout: 10 * time.Second}
resp, err := c.Get(url)`},
		{"POST JSON", `body, _ := json.Marshal(payload)
resp, err := http.Post(url, "application/json", bytes.NewReader(body))`},
		{"Build a request with headers", `req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
req.Header.Set("Authorization", "Bearer "+token)
resp, err := http.DefaultClient.Do(req)`},
		{"Tiny server", `http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
})
log.Fatal(http.ListenAndServe(":8080", nil))`},
		{"Path values (Go 1.22+)", `mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintln(w, id)
})`},
	})

	RegisterCheatsheet("bufio", `Buffer a Reader/Writer for cheap line-at-a-time I/O. Scanner is the right tool for "read a file/stdin line by line".`, []CheatRow{
		{"Read stdin line by line", `sc := bufio.NewScanner(os.Stdin)
for sc.Scan() {
    line := sc.Text()
}
if err := sc.Err(); err != nil { ... }`},
		{"Read a file line by line", `f, err := os.Open("data.txt")
if err != nil { return err }
defer f.Close()
sc := bufio.NewScanner(f)
for sc.Scan() { handle(sc.Text()) }`},
		{"Lift Scanner line limit (default 64KB)", `sc.Buffer(make([]byte, 0, 64*1024), 1<<20)`},
		{"Buffered writer (don't forget Flush)", `w := bufio.NewWriter(f)
defer w.Flush()
w.WriteString("hi\n")`},
	})

	RegisterCheatsheet("sync", `Goroutine coordination primitives. Reach for channels first; reach for sync when channels would obscure intent (counters, lazy init, simple shared state).`, []CheatRow{
		{"Wait for N goroutines", `var wg sync.WaitGroup
for _, item := range items {
    wg.Add(1)
    go func(it Item) {
        defer wg.Done()
        process(it)
    }(item)
}
wg.Wait()`},
		{"One-time init", `var once sync.Once
once.Do(setup)`},
		{"Mutex around shared state", `var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()`},
		{"Read-mostly map", `var mu sync.RWMutex
mu.RLock(); v := m[k]; mu.RUnlock()`},
		{"Pool for reusable buffers", `var bufPool = sync.Pool{
    New: func() any { return new(bytes.Buffer) },
}`},
	})

	RegisterCheatsheet("regexp", `RE2-flavored regular expressions — linear-time, no backreferences. For simple substring matches use strings.Contains; for splitting on whitespace strings.Fields. Compile patterns once at package init.`, []CheatRow{
		{"Compile (panics on bad pattern)", `var hex = regexp.MustCompile(`+"`"+`^[0-9a-f]+$`+"`"+`)`},
		{"Match boolean", `if hex.MatchString(s) { ... }`},
		{"Find first match", `m := re.FindString(s)`},
		{"Find all matches", `all := re.FindAllString(s, -1)`},
		{"Capture groups", `m := re.FindStringSubmatch(s)
// m[0] full, m[1] first group...`},
		{"Replace with template", `out := re.ReplaceAllString(s, "$1-$2")`},
	})
}

