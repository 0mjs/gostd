package content

func init() {
	Register(&Package{
		Name:       "net/http/pprof",
		ImportPath: "net/http/pprof",
		Category:   "Networking",
		Summary:    "Serve runtime profiling data over HTTP at /debug/pprof/. Import for side effects to enable.",
		Sections: []Section{
			{
				Title: "Enable the handlers",
				Examples: []Example{
					{
						Title: "Blank import — handlers attach to http.DefaultServeMux",
						Notes: "If you use a custom mux, you'll need to register the handlers yourself. Never expose pprof publicly — keep it on localhost or behind auth.",
						Code: `import (
    "net/http"
    _ "net/http/pprof"
)

func main() {
    go http.ListenAndServe("localhost:6060", nil)
    // ...
}`,
					},
					{
						Title: "Profile with go tool pprof",
						Code: `# 30-second CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# heap snapshot
go tool pprof http://localhost:6060/debug/pprof/heap

# goroutines
curl http://localhost:6060/debug/pprof/goroutine?debug=2`,
					},
				},
			},
		},
	})
}
