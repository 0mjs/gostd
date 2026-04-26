package content

func init() {
	Register(&Package{
		Name:       "math/rand",
		ImportPath: "math/rand",
		Category:   "Math",
		Summary:    "Legacy pseudo-random package. Prefer math/rand/v2 in new code. Keep this around only for unchanged legacy APIs.",
		Sections: []Section{
			{
				Title: "Why prefer v2?",
				Description: "math/rand has a global mutex (serializes goroutine calls), needed explicit seeding before Go 1.20, and has a bigger, messier API. math/rand/v2 fixes all three. See the math/rand/v2 page on this site.",
			},
		},
	})
}
