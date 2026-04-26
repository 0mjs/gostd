package content

func init() {
	Register(&Package{
		Name:       "runtime/trace",
		ImportPath: "runtime/trace",
		Category:   "Runtime & Debug",
		Summary:    "Execution tracer — fine-grained timeline of goroutines, syscalls, GC, and user-defined regions. Analyze with `go tool trace`.",
		Sections: []Section{
			{
				Title: "Record a trace",
				Examples: []Example{
					{Title: "Start / Stop", Code: `f, _ := os.Create("trace.out")
trace.Start(f)
defer trace.Stop()

// ... workload ...`},
				},
			},
			{
				Title: "User annotations",
				Examples: []Example{
					{Title: "Regions and tasks", Code: `ctx, task := trace.NewTask(ctx, "request")
defer task.End()

trace.WithRegion(ctx, "db.query", func() {
    db.Query(...)
})

trace.Log(ctx, "user_id", "42")`},
				},
			},
			{
				Title: "Analyze",
				Examples: []Example{
					{Title: "go tool trace", Code: `go tool trace trace.out
# opens browser with goroutine, network, GC views`},
				},
			},
		},
	})
}
