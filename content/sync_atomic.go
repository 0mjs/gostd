package content

func init() {
	Register(&Package{
		Name:       "sync/atomic",
		ImportPath: "sync/atomic",
		Category:   "Concurrency",
		Summary:    "Low-level atomic operations. Lock-free counters, flags, and pointers — faster than a Mutex when the state fits in a word.",
		Sections: []Section{
			{
				Title: "Typed wrappers (Go 1.19+) — prefer these",
				Description: "atomic.Int32, Int64, Uint32, Uint64, Uintptr, Bool, Pointer[T], Value. Methods like Load, Store, Add, CompareAndSwap, Swap.",
				Examples: []Example{
					{
						Title: "atomic.Int64 counter",
						Code: `var hits atomic.Int64

for i := 0; i < 5; i++ {
    go hits.Add(1)
}
time.Sleep(time.Millisecond)
fmt.Println(hits.Load())`,
					},
					{
						Title: "atomic.Bool flag",
						Code: `var closed atomic.Bool

if closed.CompareAndSwap(false, true) {
    // we're the one who closed it
    close(ch)
}`,
					},
					{
						Title: "atomic.Pointer[T] — lock-free pointer swap",
						Notes: "Perfect for hot-reloadable config or snapshot-style data structures.",
						Code: `var cfgP atomic.Pointer[Config]
cfgP.Store(&Config{...})

// readers
cfg := cfgP.Load()

// hot-reload
cfgP.Store(newCfg)`,
					},
					{
						Title: "atomic.Value — for arbitrary types (pre-generics)",
						Notes: "Still useful if you need something other than a pointer. The value must always be the same concrete type.",
						Code: `var v atomic.Value
v.Store("hello")
s := v.Load().(string)`,
					},
				},
			},
			{
				Title: "Raw functions — still here for interop",
				Description: "The original AddInt32, LoadUint64, etc. still work. Prefer the typed wrappers in new code — same performance, harder to misuse.",
			},
		},
	})
}
