package content

func init() {
	Register(&Package{
		Name:       "net/http",
		ImportPath: "net/http",
		Category:   "Networking",
		Summary:    "HTTP server and client. Batteries included — no framework required.",
		Sections: []Section{
			{
				Title: "Client — making requests",
				Examples: []Example{
					{
						Title: "http.Get — the simplest thing",
						Notes: "http.Get, Post, and Head use http.DefaultClient. Fine for scripts; avoid in production — no timeout.",
						Code: `resp, err := http.Get("https://example.com")
if err != nil { log.Fatal(err) }
defer resp.Body.Close()

b, _ := io.ReadAll(resp.Body)
fmt.Println(resp.StatusCode, len(b))`,
					},
					{
						Title: "Always use a client with a timeout",
						Notes: "The zero-value http.Client has NO timeout and will hang forever on a bad network. Always set Timeout.",
						Code: `client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Get(url)`,
					},
					{
						Title: "NewRequestWithContext — cancellable requests",
						Code: `req, _ := http.NewRequestWithContext(ctx, "POST", url, body)
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer "+token)

resp, err := client.Do(req)`,
					},
					{
						Title: "Always drain and close the body",
						Notes: "If you Close without draining, the underlying TCP connection can't be reused. io.Copy(io.Discard, body) drains it cheaply.",
						Code: `defer func() {
    io.Copy(io.Discard, resp.Body)
    resp.Body.Close()
}()`,
					},
				},
			},
			{
				Title: "Server — handling requests",
				Examples: []Example{
					{
						Title: "Minimal server",
						Code: `func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "hello world")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}`,
					},
					{
						Title: "ServeMux patterns (Go 1.22+)",
						Notes: "ServeMux now supports method matching and wildcards. For most apps you don't need a third-party router anymore.",
						Code: `mux := http.NewServeMux()
mux.HandleFunc("GET /users/{id}", getUser)
mux.HandleFunc("POST /users", createUser)
mux.HandleFunc("GET /static/", serveStatic)   // subtree
mux.HandleFunc("GET /{$}", home)              // exact "/"

http.ListenAndServe(":8080", mux)`,
					},
					{
						Title: "Path parameters",
						Code: `func getUser(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    ...
}`,
					},
					{
						Title: "Production server — explicit timeouts and graceful shutdown",
						Code: `srv := &http.Server{
    Addr:              ":8080",
    Handler:           mux,
    ReadHeaderTimeout: 5 * time.Second,
    ReadTimeout:       15 * time.Second,
    WriteTimeout:      30 * time.Second,
    IdleTimeout:       60 * time.Second,
}

ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()

go srv.ListenAndServe()
<-ctx.Done()

shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
srv.Shutdown(shutdownCtx)`,
					},
				},
			},
			{
				Title: "Reading and writing bodies",
				Examples: []Example{
					{
						Title: "Decode JSON request body",
						Code: `func createUser(w http.ResponseWriter, r *http.Request) {
    var u User
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()
    if err := dec.Decode(&u); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ...
}`,
					},
					{
						Title: "Write JSON response",
						Code: `w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(u)`,
					},
					{
						Title: "Form values and query params",
						Code: `r.ParseForm()            // populates r.Form from URL + body
q := r.URL.Query()       // just query params
name := q.Get("name")
tags := q["tag"]         // repeated ?tag=a&tag=b`,
					},
				},
			},
			{
				Title: "Middleware",
				Examples: []Example{
					{
						Title: "Wrap a handler",
						Notes: "Middleware is just a func(Handler) Handler. Chain by composition.",
						Code: `func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        slog.Info("served", "method", r.Method, "path", r.URL.Path, "took", time.Since(start))
    })
}

http.ListenAndServe(":8080", logging(mux))`,
					},
				},
			},
			{
				Title: "Testing with httptest",
				Examples: []Example{
					{
						Title: "Spin up a real server in a test",
						Code: `srv := httptest.NewServer(http.HandlerFunc(myHandler))
defer srv.Close()

resp, err := http.Get(srv.URL + "/hello")`,
					},
					{
						Title: "Call a handler directly without listening",
						Code: `rec := httptest.NewRecorder()
req := httptest.NewRequest("GET", "/", nil)
myHandler(rec, req)

fmt.Println(rec.Code, rec.Body.String())`,
					},
				},
			},
		},
	})
}
