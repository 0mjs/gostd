package content

import "sort"

// Recipe is a single "How do I...?" entry — a concrete task paired with the
// minimum-viable Go that does it, and a link to the package it draws from.
type Recipe struct {
	Title string   // the "How do I X?" question, phrased as a task
	Code  string   // 1–8 lines that solve it
	Notes string   // optional "watch out for" context
	Pkgs  []string // import paths used; first one is the primary anchor
	Group string   // category for grouping on the recipes page
}

// Pkg returns the primary package for the recipe (the one we link to).
func (r Recipe) Pkg() string {
	if len(r.Pkgs) == 0 {
		return ""
	}
	return r.Pkgs[0]
}

var recipes = []Recipe{
	// ── Strings & numbers ──────────────────────────────────────────────
	{Group: "Strings & numbers", Title: "Check if a string contains a substring",
		Code: `if strings.Contains(s, "needle") {
    // ...
}`,
		Pkgs: []string{"strings"}},
	{Group: "Strings & numbers", Title: "Split on a separator",
		Code: `parts := strings.Split("a,b,c", ",")
// []string{"a", "b", "c"}`,
		Notes: "For whitespace use strings.Fields.",
		Pkgs:  []string{"strings"}},
	{Group: "Strings & numbers", Title: "Join a slice of strings",
		Code: `s := strings.Join([]string{"a", "b", "c"}, ", ")
// "a, b, c"`,
		Pkgs: []string{"strings"}},
	{Group: "Strings & numbers", Title: "Trim whitespace",
		Code: `s = strings.TrimSpace(s)`,
		Pkgs: []string{"strings"}},
	{Group: "Strings & numbers", Title: "Build a long string efficiently",
		Code: `var b strings.Builder
for i := 0; i < 1000; i++ {
    fmt.Fprintf(&b, "row=%d ", i)
}
result := b.String()`,
		Notes: "strings.Builder avoids the O(n²) of repeated `s += ...`.",
		Pkgs:  []string{"strings"}},
	{Group: "Strings & numbers", Title: "Convert a string to int",
		Code: `n, err := strconv.Atoi("42")
if err != nil {
    return fmt.Errorf("parse count: %w", err)
}`,
		Pkgs: []string{"strconv"}},
	{Group: "Strings & numbers", Title: "Convert int to string",
		Code: `s := strconv.Itoa(42)
// "42"`,
		Pkgs: []string{"strconv"}},
	{Group: "Strings & numbers", Title: "Format a float with N decimal places",
		Code: `fmt.Sprintf("%.2f", 3.14159) // "3.14"`,
		Pkgs: []string{"fmt"}},

	// ── Files & I/O ────────────────────────────────────────────────────
	{Group: "Files & I/O", Title: "Read a whole file into memory",
		Code: `data, err := os.ReadFile("config.yaml")
if err != nil {
    return err
}`,
		Notes: "Fine for small files. For huge files, stream with os.Open + bufio.",
		Pkgs:  []string{"os"}},
	{Group: "Files & I/O", Title: "Write a whole file (create or overwrite)",
		Code: `err := os.WriteFile("out.txt", data, 0644)`,
		Pkgs: []string{"os"}},
	{Group: "Files & I/O", Title: "Read a file line by line",
		Code: `f, err := os.Open("data.txt")
if err != nil {
    return err
}
defer f.Close()

sc := bufio.NewScanner(f)
for sc.Scan() {
    line := sc.Text()
    process(line)
}
if err := sc.Err(); err != nil {
    return err
}`,
		Pkgs: []string{"bufio", "os"}},
	{Group: "Files & I/O", Title: "Check if a file exists",
		Code: `_, err := os.Stat(path)
switch {
case err == nil:
    // exists
case errors.Is(err, fs.ErrNotExist):
    // does not exist
default:
    return err // some other error (permissions, I/O)
}`,
		Pkgs: []string{"os", "errors", "io/fs"}},
	{Group: "Files & I/O", Title: "Walk a directory tree",
		Code: `err := filepath.WalkDir("dir", func(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    if !d.IsDir() && strings.HasSuffix(path, ".go") {
        fmt.Println(path)
    }
    return nil
})`,
		Pkgs: []string{"path/filepath"}},
	{Group: "Files & I/O", Title: "Make a temp file or directory",
		Code: `f, err := os.CreateTemp("", "myapp-*.json")
if err != nil {
    return err
}
defer os.Remove(f.Name())
defer f.Close()

dir, err := os.MkdirTemp("", "myapp-*")
defer os.RemoveAll(dir)`,
		Pkgs: []string{"os"}},
	{Group: "Files & I/O", Title: "Copy a stream from reader to writer",
		Code: `_, err := io.Copy(dst, src)`,
		Pkgs: []string{"io"}},
	{Group: "Files & I/O", Title: "Read all stdin",
		Code: `data, err := io.ReadAll(os.Stdin)`,
		Pkgs: []string{"io", "os"}},

	// ── HTTP ───────────────────────────────────────────────────────────
	{Group: "HTTP", Title: "Make an HTTP GET",
		Code: `resp, err := http.Get("https://example.com")
if err != nil {
    return err
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)`,
		Notes: "http.DefaultClient has no timeout. For anything outside scripts, build your own &http.Client{Timeout: ...}.",
		Pkgs:  []string{"net/http", "io"}},
	{Group: "HTTP", Title: "POST a JSON body",
		Code: `body, _ := json.Marshal(payload)
resp, err := http.Post(url, "application/json", bytes.NewReader(body))
if err != nil {
    return err
}
defer resp.Body.Close()`,
		Pkgs: []string{"net/http", "encoding/json", "bytes"}},
	{Group: "HTTP", Title: "HTTP request with timeout and headers",
		Code: `client := &http.Client{Timeout: 10 * time.Second}
req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
if err != nil {
    return err
}
req.Header.Set("Authorization", "Bearer "+token)
resp, err := client.Do(req)`,
		Pkgs: []string{"net/http", "time", "context"}},
	{Group: "HTTP", Title: "Tiny HTTP server",
		Code: `mux := http.NewServeMux()
mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello, %s\n", r.PathValue("name"))
})
log.Fatal(http.ListenAndServe(":8080", mux))`,
		Notes: "Path patterns and PathValue are Go 1.22+.",
		Pkgs:  []string{"net/http"}},
	{Group: "HTTP", Title: "Parse a URL and its query parameters",
		Code: `u, err := url.Parse("https://api.x/?id=42&since=2026")
if err != nil {
    return err
}
id := u.Query().Get("id")`,
		Pkgs: []string{"net/url"}},

	// ── JSON ───────────────────────────────────────────────────────────
	{Group: "JSON", Title: "Encode a struct to JSON bytes",
		Code: `type User struct {
    Name string ` + "`json:\"name\"`" + `
    Age  int    ` + "`json:\"age,omitempty\"`" + `
}
b, err := json.Marshal(User{Name: "Ada", Age: 36})
// {"name":"Ada","age":36}`,
		Pkgs: []string{"encoding/json"}},
	{Group: "JSON", Title: "Pretty-print JSON",
		Code: `b, err := json.MarshalIndent(v, "", "  ")`,
		Pkgs: []string{"encoding/json"}},
	{Group: "JSON", Title: "Decode JSON into a struct",
		Code: `var u User
if err := json.Unmarshal(data, &u); err != nil {
    return err
}`,
		Pkgs: []string{"encoding/json"}},
	{Group: "JSON", Title: "Decode arbitrary JSON",
		Code: `var v any
if err := json.Unmarshal(data, &v); err != nil {
    return err
}
// v is map[string]any / []any / string / float64 / bool / nil`,
		Pkgs: []string{"encoding/json"}},
	{Group: "JSON", Title: "Stream-decode JSON from an HTTP body",
		Code: `defer resp.Body.Close()
var u User
if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
    return err
}`,
		Notes: "Avoids loading the whole body into memory before parsing.",
		Pkgs:  []string{"encoding/json"}},

	// ── Time ───────────────────────────────────────────────────────────
	{Group: "Time", Title: "Sleep for a duration",
		Code: `time.Sleep(2 * time.Second)`,
		Pkgs: []string{"time"}},
	{Group: "Time", Title: "Time how long something takes",
		Code: `start := time.Now()
doWork()
fmt.Println("took", time.Since(start))`,
		Pkgs: []string{"time"}},
	{Group: "Time", Title: "Format and parse RFC3339",
		Code: `s := time.Now().Format(time.RFC3339)
t, err := time.Parse(time.RFC3339, s)`,
		Pkgs: []string{"time"}},
	{Group: "Time", Title: "Run a function on an interval",
		Code: `tk := time.NewTicker(5 * time.Second)
defer tk.Stop()
for {
    select {
    case <-tk.C:
        beat()
    case <-ctx.Done():
        return
    }
}`,
		Pkgs: []string{"time", "context"}},

	// ── Concurrency ────────────────────────────────────────────────────
	{Group: "Concurrency", Title: "Wait for a group of goroutines",
		Code: `var wg sync.WaitGroup
for _, item := range items {
    wg.Add(1)
    go func(it Item) {
        defer wg.Done()
        process(it)
    }(item)
}
wg.Wait()`,
		Notes: "Capture loop variables explicitly. (In Go 1.22+ this is no longer required, but explicit is still clearer.)",
		Pkgs:  []string{"sync"}},
	{Group: "Concurrency", Title: "Cancel work with context",
		Code: `ctx, cancel := context.WithCancel(parent)
defer cancel()

go func() {
    if userPressedStop() {
        cancel()
    }
}()

select {
case result := <-doWork(ctx):
    use(result)
case <-ctx.Done():
    return ctx.Err()
}`,
		Pkgs: []string{"context"}},
	{Group: "Concurrency", Title: "Run with a hard timeout",
		Code: `ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := http.DefaultClient.Do(req)`,
		Pkgs: []string{"context", "net/http"}},
	{Group: "Concurrency", Title: "Limit concurrency with a semaphore channel",
		Code: `sem := make(chan struct{}, 8) // 8 in flight at once
for _, job := range jobs {
    sem <- struct{}{}
    go func(j Job) {
        defer func() { <-sem }()
        run(j)
    }(job)
}`,
		Pkgs: []string{"sync"}},
	{Group: "Concurrency", Title: "One-time lazy initialization",
		Code: `var (
    once sync.Once
    db   *sql.DB
)
func getDB() *sql.DB {
    once.Do(func() { db = openDB() })
    return db
}`,
		Pkgs: []string{"sync"}},

	// ── Slices & maps ──────────────────────────────────────────────────
	{Group: "Slices & maps", Title: "Sort a slice",
		Code: `slices.Sort(nums)`,
		Pkgs: []string{"slices"}},
	{Group: "Slices & maps", Title: "Sort a slice by a struct field",
		Code: `slices.SortFunc(users, func(a, b User) int {
    return cmp.Compare(a.Age, b.Age)
})`,
		Pkgs: []string{"slices", "cmp"}},
	{Group: "Slices & maps", Title: "Deduplicate a slice",
		Code: `slices.Sort(s)
s = slices.Compact(s)`,
		Notes: "Compact removes adjacent duplicates only; sort first.",
		Pkgs:  []string{"slices"}},
	{Group: "Slices & maps", Title: "Get sorted keys of a map",
		Code: `keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
slices.Sort(keys)`,
		Pkgs: []string{"slices", "maps"}},
	{Group: "Slices & maps", Title: "Iterate a map in key order",
		Code: `for _, k := range slices.Sorted(maps.Keys(m)) {
    fmt.Println(k, m[k])
}`,
		Notes: "Go 1.23+ — uses iter.Seq from maps.Keys.",
		Pkgs:  []string{"maps", "slices"}},

	// ── Errors ─────────────────────────────────────────────────────────
	{Group: "Errors", Title: "Wrap an error with context",
		Code: `if err := load(id); err != nil {
    return fmt.Errorf("load user %s: %w", id, err)
}`,
		Pkgs: []string{"fmt", "errors"}},
	{Group: "Errors", Title: "Check for a sentinel error",
		Code: `var ErrNoUser = errors.New("no user")

if errors.Is(err, ErrNoUser) {
    // 404 path
}`,
		Pkgs: []string{"errors"}},
	{Group: "Errors", Title: "Extract a typed error",
		Code: `var pe *fs.PathError
if errors.As(err, &pe) {
    log.Printf("path %q failed: %v", pe.Path, pe.Err)
}`,
		Pkgs: []string{"errors", "io/fs"}},

	// ── CLI & runtime ──────────────────────────────────────────────────
	{Group: "CLI & runtime", Title: "Parse command-line flags",
		Code: `port := flag.Int("port", 8080, "listen port")
verbose := flag.Bool("v", false, "verbose")
flag.Parse()
fmt.Println(*port, *verbose, flag.Args())`,
		Pkgs: []string{"flag"}},
	{Group: "CLI & runtime", Title: "Get an env var with a default",
		Code: `port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
// or, Go 1.22+:
port = cmp.Or(os.Getenv("PORT"), "8080")`,
		Pkgs: []string{"os", "cmp"}},
	{Group: "CLI & runtime", Title: "Catch ctrl-C cleanly",
		Code: `ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()

// pass ctx into your work; <-ctx.Done() fires on ^C`,
		Pkgs: []string{"os/signal", "context"}},
	{Group: "CLI & runtime", Title: "Exit with a non-zero status",
		Code: `if err := run(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}`,
		Notes: "os.Exit skips deferred functions; prefer returning errors from main.",
		Pkgs:  []string{"os", "fmt"}},

	// ── Hashing & crypto ───────────────────────────────────────────────
	{Group: "Hashing & crypto", Title: "SHA-256 a string or bytes",
		Code: `sum := sha256.Sum256([]byte("hello"))
hex := hex.EncodeToString(sum[:])`,
		Pkgs: []string{"crypto/sha256", "encoding/hex"}},
	{Group: "Hashing & crypto", Title: "Generate cryptographically random bytes",
		Code: `buf := make([]byte, 32)
if _, err := rand.Read(buf); err != nil {
    return err
}`,
		Notes: "crypto/rand, never math/rand, for anything that protects secrets.",
		Pkgs:  []string{"crypto/rand"}},
	{Group: "Hashing & crypto", Title: "Random integer for non-security purposes",
		Code: `n := rand.IntN(100) // 0..99`,
		Notes: "Go 1.22+ math/rand/v2 — auto-seeded, no Seed call needed.",
		Pkgs:  []string{"math/rand/v2"}},

	// ── Testing ────────────────────────────────────────────────────────
	{Group: "Testing", Title: "Write a unit test",
		Code: `func TestAdd(t *testing.T) {
    got := add(1, 2)
    if got != 3 {
        t.Fatalf("add(1,2) = %d, want 3", got)
    }
}`,
		Pkgs: []string{"testing"}},
	{Group: "Testing", Title: "Table-driven test",
		Code: `func TestAdd(t *testing.T) {
    cases := []struct {
        name string
        a, b int
        want int
    }{
        {"zero", 0, 0, 0},
        {"pos", 2, 3, 5},
        {"neg", -1, 1, 0},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            if got := add(c.a, c.b); got != c.want {
                t.Fatalf("add(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
            }
        })
    }
}`,
		Pkgs: []string{"testing"}},

	// ── Logging ────────────────────────────────────────────────────────
	{Group: "Logging", Title: "Structured log line with slog",
		Code: `slog.Info("request",
    "method", r.Method,
    "path", r.URL.Path,
    "status", code,
    "ms", elapsed.Milliseconds(),
)`,
		Notes: "Default handler writes text. For JSON: slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil))).",
		Pkgs:  []string{"log/slog"}},
}

// AllRecipes returns every registered recipe.
func AllRecipes() []Recipe { return recipes }

// RecipeGroup is a category for the recipes index page.
type RecipeGroup struct {
	Name    string
	Recipes []Recipe
}

// recipeGroupOrder controls how groups appear on the recipes page.
var recipeGroupOrder = map[string]int{
	"Strings & numbers":  1,
	"Files & I/O":        2,
	"HTTP":               3,
	"JSON":               4,
	"Time":               5,
	"Concurrency":        6,
	"Slices & maps":      7,
	"Errors":             8,
	"CLI & runtime":      9,
	"Hashing & crypto":   10,
	"Testing":            11,
	"Logging":            12,
}

// RecipeGroups groups AllRecipes by Group, in canonical order.
func RecipeGroups() []RecipeGroup {
	byName := map[string]*RecipeGroup{}
	var groups []*RecipeGroup
	for _, r := range recipes {
		g, ok := byName[r.Group]
		if !ok {
			g = &RecipeGroup{Name: r.Group}
			byName[r.Group] = g
			groups = append(groups, g)
		}
		g.Recipes = append(g.Recipes, r)
	}
	sort.SliceStable(groups, func(i, j int) bool {
		oi, oj := recipeGroupOrder[groups[i].Name], recipeGroupOrder[groups[j].Name]
		if oi == 0 {
			oi = 99
		}
		if oj == 0 {
			oj = 99
		}
		if oi != oj {
			return oi < oj
		}
		return groups[i].Name < groups[j].Name
	})
	out := make([]RecipeGroup, len(groups))
	for i, g := range groups {
		out[i] = *g
	}
	return out
}
