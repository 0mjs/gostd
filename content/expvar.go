package content

func init() {
	Register(&Package{
		Name:       "expvar",
		ImportPath: "expvar",
		Category:   "Errors & Logging",
		Summary:    "Expose internal counters and stats as JSON at /debug/vars. Cheap, built-in observability.",
		Sections: []Section{
			{
				Title: "Publish variables",
				Examples: []Example{
					{Title: "Int / Float / String counters", Code: `var requests = expvar.NewInt("requests")
var errs = expvar.NewInt("errors")

func handle(w http.ResponseWriter, r *http.Request) {
    requests.Add(1)
}`},
					{Title: "Map of counters", Code: `var byStatus = expvar.NewMap("by_status")
byStatus.Add("200", 1)`},
					{Title: "Custom Func var", Code: `expvar.Publish("uptime", expvar.Func(func() any {
    return time.Since(start).Seconds()
}))`},
				},
			},
			{
				Title: "HTTP endpoint",
				Description: "Importing expvar registers /debug/vars on http.DefaultServeMux automatically.",
				Examples: []Example{
					{Title: "Import for side effect", Code: `import _ "expvar"
// then: go http.ListenAndServe(":6060", nil)
// curl localhost:6060/debug/vars`},
				},
			},
		},
	})
}
