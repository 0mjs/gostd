package content

func init() {
	Register(&Package{
		Name:       "runtime/metrics",
		ImportPath: "runtime/metrics",
		Category:   "Runtime & Debug",
		Summary:    "Structured, versioned access to runtime metrics (heap, GC pauses, scheduler latency, etc). Preferred over runtime.MemStats for monitoring.",
		Sections: []Section{
			{
				Title: "Read metrics",
				Examples: []Example{
					{Title: "Sample a set of metrics", Code: `samples := []metrics.Sample{
    {Name: "/gc/pauses:seconds"},
    {Name: "/memory/classes/heap/objects:bytes"},
    {Name: "/sched/goroutines:goroutines"},
}
metrics.Read(samples)
for _, s := range samples {
    fmt.Println(s.Name, s.Value.Kind())
}`},
					{Title: "Discover available metrics", Code: `for _, d := range metrics.All() {
    fmt.Println(d.Name, "-", d.Description)
}`},
				},
			},
		},
	})
}
