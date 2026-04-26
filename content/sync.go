package content

func init() {
	Register(&Package{
		Name:       "sync",
		ImportPath: "sync",
		Category:   "Concurrency",
		Summary:    "Low-level synchronization primitives: mutexes, WaitGroups, Once, Pool, Map, Cond.",
		Sections: []Section{
			{
				Title: "Mutex and RWMutex",
				Description: "Guard shared data. Don't share by communicating? You're allowed — use channels OR a mutex, whichever fits.",
				Examples: []Example{
					{
						Title: "Mutex — one-at-a-time access",
						Notes: "Zero value is an unlocked mutex. Never copy a Mutex after first use — put it in a pointer or on a struct kept by pointer.",
						Code: `type Counter struct {
    mu sync.Mutex
    n  int
}
func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}`,
					},
					{
						Title: "RWMutex — many readers, one writer",
						Notes: "Reach for RWMutex only when reads vastly outnumber writes. For contention-light cases a plain Mutex is faster.",
						Code: `var mu sync.RWMutex

// reader side
mu.RLock()
v := cache[k]
mu.RUnlock()

// writer side
mu.Lock()
cache[k] = v
mu.Unlock()`,
					},
				},
			},
			{
				Title: "WaitGroup — wait for a set of goroutines",
				Examples: []Example{
					{
						Title: "Classic fan-out/wait",
						Notes: "Call Add BEFORE starting the goroutine. The common bug is Add inside the goroutine (race).",
						Code: `var wg sync.WaitGroup
for _, u := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        fetch(u)
    }(u)
}
wg.Wait()`,
					},
					{
						Title: "WaitGroup.Go (Go 1.25+)",
						Notes: "Go 1.25 added wg.Go(func()) which does Add+go+Done for you. Cleaner and harder to misuse.",
						Code: `var wg sync.WaitGroup
for _, u := range urls {
    wg.Go(func() { fetch(u) })
}
wg.Wait()`,
					},
				},
			},
			{
				Title: "Once — run exactly once",
				Examples: []Example{
					{
						Title: "Lazy initialization",
						Code: `var (
    once sync.Once
    cfg  *Config
)
func get() *Config {
    once.Do(func() { cfg = load() })
    return cfg
}`,
					},
					{
						Title: "OnceFunc / OnceValue (1.21+)",
						Notes: "Newer typed wrappers — often clearer than a manual sync.Once + package-level var.",
						Code: `var loadConfig = sync.OnceValue(func() *Config {
    return load()
})

cfg := loadConfig()  // computed once, cached forever`,
					},
				},
			},
			{
				Title: "Pool — reuse allocations",
				Description: "A free list of reusable objects. Great for large buffers that get created and thrown away in hot paths.",
				Examples: []Example{
					{
						Title: "Pool of []byte buffers",
						Notes: "Items in the pool may be GC'd at any time — don't rely on them sticking around. Never reuse content between Put and Get.",
						Code: `var bufPool = sync.Pool{
    New: func() any { return make([]byte, 0, 4096) },
}

buf := bufPool.Get().([]byte)
defer bufPool.Put(buf[:0])

buf = append(buf, "hello"...)`,
					},
				},
			},
			{
				Title: "Map — concurrent map",
				Description: "Avoid by default. A normal map + sync.RWMutex is simpler and usually faster. sync.Map shines only for two specific patterns: (1) keys written once and read many times; (2) disjoint keys per goroutine.",
				Examples: []Example{
					{
						Title: "Basic usage",
						Code: `var m sync.Map
m.Store("a", 1)
v, ok := m.Load("a")
m.Range(func(k, v any) bool {
    fmt.Println(k, v)
    return true   // return false to stop
})`,
					},
				},
			},
			{
				Title: "Atomic — lock-free primitives",
				Description: "sync/atomic provides atomic loads, stores, adds, CAS. Use for simple counters where a Mutex would be overkill.",
				Examples: []Example{
					{
						Title: "atomic.Int64 — typed, easy to use",
						Code: `var n atomic.Int64
go func() { n.Add(1) }()
go func() { n.Add(1) }()
// ...
fmt.Println(n.Load())`,
					},
				},
			},
		},
	})
}
