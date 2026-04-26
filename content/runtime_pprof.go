package content

func init() {
	Register(&Package{
		Name:       "runtime/pprof",
		ImportPath: "runtime/pprof",
		Category:   "Runtime & Debug",
		Summary:    "Write CPU, heap, block, mutex, goroutine profiles. Analyze with `go tool pprof`.",
		Sections: []Section{
			{
				Title: "CPU profile",
				Examples: []Example{
					{Title: "Start / Stop", Code: `f, _ := os.Create("cpu.pprof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()

// ... workload ...`},
				},
			},
			{
				Title: "Heap / goroutine / other profiles",
				Examples: []Example{
					{Title: "Lookup + WriteTo", Code: `f, _ := os.Create("heap.pprof")
pprof.Lookup("heap").WriteTo(f, 0)

// others: "goroutine", "allocs", "block", "mutex", "threadcreate"`},
				},
			},
			{
				Title: "Labels",
				Examples: []Example{
					{Title: "Tag goroutines for profiling", Code: `pprof.Do(ctx, pprof.Labels("endpoint", "/users"), func(ctx context.Context) {
    serve(ctx)
})`},
				},
			},
			{
				Title: "Analyze",
				Examples: []Example{
					{Title: "go tool pprof", Code: `go tool pprof cpu.pprof
(pprof) top
(pprof) web
(pprof) list funcName`},
				},
			},
		},
	})
}
