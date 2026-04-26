package content

func init() {
	Register(&Package{
		Name:       "math/rand/v2",
		ImportPath: "math/rand/v2",
		Category:   "Math",
		Summary:    "Pseudo-random numbers (Go 1.22+). Auto-seeded, simpler API, and not safe for cryptographic use — reach for crypto/rand for that.",
		Sections: []Section{
			{
				Title: "v2 vs v1",
				Description: "math/rand/v2 is a cleaner redo. No more manual seeding, no global mutex bottleneck on concurrent use, and a tighter API. Prefer v2 for new code.",
			},
			{
				Title: "Random numbers",
				Examples: []Example{
					{
						Title: "Int, IntN, Float64",
						Notes: "IntN(n) returns a number in [0, n). Float64() returns [0.0, 1.0).",
						Code: `rand.IntN(6) + 1          // dice roll 1..6
rand.Float64()            // [0, 1)
rand.Int32()              // full range int32
rand.Uint64()`,
					},
					{
						Title: "N[T] — typed-range generic (great quality-of-life)",
						Code: `d := rand.N(6) + 1        // returns int, [0, 6)
t := rand.N(10 * time.Second)  // random duration up to 10s`,
					},
				},
			},
			{
				Title: "Shuffle and pick",
				Examples: []Example{
					{
						Title: "Shuffle a slice in place",
						Code: `xs := []int{1, 2, 3, 4, 5}
rand.Shuffle(len(xs), func(i, j int) {
    xs[i], xs[j] = xs[j], xs[i]
})`,
					},
					{
						Title: "Perm — a random permutation",
						Code: `order := rand.Perm(5)   // e.g., [3 0 4 1 2]`,
					},
				},
			},
			{
				Title: "Explicit sources for reproducibility",
				Examples: []Example{
					{
						Title: "Seed a Rand for deterministic output",
						Notes: "Use this in tests. The default global rand is seeded randomly per process — good for production, bad for assertions.",
						Code: `r := rand.New(rand.NewPCG(1, 2))
for i := 0; i < 3; i++ {
    fmt.Println(r.IntN(100))
}`,
					},
				},
			},
		},
	})
}
