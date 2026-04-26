package content

func init() {
	Register(&Package{
		Name:       "runtime",
		ImportPath: "runtime",
		Category:   "Runtime & Debug",
		Summary:    "Hooks into the Go runtime: GOMAXPROCS, goroutine count, GC triggering, caller info, finalizers.",
		Sections: []Section{
			{
				Title: "CPU and goroutines",
				Examples: []Example{
					{Title: "GOMAXPROCS", Code: `old := runtime.GOMAXPROCS(4)
runtime.GOMAXPROCS(runtime.NumCPU())`},
					{Title: "NumGoroutine", Code: `fmt.Println("goroutines:", runtime.NumGoroutine())`},
					{Title: "Gosched / Goexit", Code: `runtime.Gosched() // yield to scheduler
runtime.Goexit()   // terminate current goroutine (deferred funcs still run)`},
				},
			},
			{
				Title: "GC",
				Examples: []Example{
					{Title: "Force GC", Code: `runtime.GC()          // blocking full GC
debug.SetGCPercent(200) // tune GC trigger (debug package)`},
					{Title: "MemStats", Code: `var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Println(m.HeapAlloc, m.NumGC)`},
				},
			},
			{
				Title: "Caller info",
				Examples: []Example{
					{Title: "Caller / Callers", Code: `pc, file, line, ok := runtime.Caller(1)
if ok {
    fn := runtime.FuncForPC(pc)
    fmt.Printf("%s at %s:%d\n", fn.Name(), file, line)
}`},
				},
			},
			{
				Title: "Finalizers and cleanups",
				Examples: []Example{
					{Title: "SetFinalizer (legacy)", Code: `runtime.SetFinalizer(obj, func(o *T) { o.Close() })`},
					{Title: "AddCleanup (Go 1.24+)", Code: `runtime.AddCleanup(obj, func(h *Handle) { h.Free() }, handle)`},
				},
			},
			{
				Title: "Runtime info",
				Examples: []Example{
					{Title: "GOOS / GOARCH / Version", Code: `fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.Version())`},
				},
			},
		},
	})
}
