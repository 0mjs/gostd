package content

func init() {
	Register(&Package{
		Name:       "runtime/debug",
		ImportPath: "runtime/debug",
		Category:   "Runtime & Debug",
		Summary:    "GC tuning, memory limits, stack traces, and build info.",
		Sections: []Section{
			{
				Title: "GC and memory tuning",
				Examples: []Example{
					{Title: "SetGCPercent", Code: `old := debug.SetGCPercent(200) // less frequent GC
debug.SetGCPercent(-1)          // disable`},
					{Title: "SetMemoryLimit (Go 1.19+)", Code: `debug.SetMemoryLimit(1 << 30) // 1 GiB soft limit`},
					{Title: "FreeOSMemory", Code: `debug.FreeOSMemory() // hint to return memory to the OS`},
				},
			},
			{
				Title: "Stack traces",
				Examples: []Example{
					{Title: "Stack / PrintStack", Code: `defer func() {
    if r := recover(); r != nil {
        log.Printf("panic: %v\n%s", r, debug.Stack())
    }
}()`},
				},
			},
			{
				Title: "Build info",
				Examples: []Example{
					{Title: "ReadBuildInfo", Code: `if info, ok := debug.ReadBuildInfo(); ok {
    fmt.Println(info.Main.Path, info.Main.Version)
    for _, s := range info.Settings {
        fmt.Println(s.Key, s.Value) // vcs.revision, vcs.time, ...
    }
}`},
				},
			},
			{
				Title: "Panic control",
				Examples: []Example{
					{Title: "SetPanicOnFault / SetTraceback", Code: `debug.SetPanicOnFault(true)
debug.SetTraceback("all") // or "system", "crash"`},
				},
			},
		},
	})
}
