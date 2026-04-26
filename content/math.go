package content

func init() {
	Register(&Package{
		Name:       "math",
		ImportPath: "math",
		Category:   "Math",
		Summary:    "Basic math: constants, elementary functions, float classification. Operates on float64.",
		Sections: []Section{
			{
				Title: "Constants",
				Examples: []Example{
					{
						Title: "Pi, E, and limit constants",
						Code: `fmt.Println(math.Pi, math.E)
fmt.Println(math.MaxFloat64, math.SmallestNonzeroFloat64)
fmt.Println(math.MaxInt, math.MinInt)          // platform int
fmt.Println(math.MaxInt64, math.MinInt64)`,
					},
				},
			},
			{
				Title: "Elementary functions",
				Examples: []Example{
					{
						Title: "Abs, Floor, Ceil, Round, Trunc",
						Code: `fmt.Println(math.Abs(-3.2))    // 3.2
fmt.Println(math.Floor(3.9))   // 3
fmt.Println(math.Ceil(3.1))    // 4
fmt.Println(math.Round(2.5))   // 3 — away from zero
fmt.Println(math.Trunc(-2.9))  // -2`,
					},
					{
						Title: "Pow, Sqrt, Cbrt, Exp, Log",
						Code: `math.Pow(2, 10)   // 1024
math.Sqrt(16)     // 4
math.Log(math.E)  // 1
math.Log10(1000)  // 3`,
					},
					{
						Title: "Trig — Sin, Cos, Tan, Atan2",
						Notes: "Atan2(y, x) gives the angle of a point from the origin — respects quadrant.",
						Code: `math.Sin(math.Pi/2)   // 1
math.Atan2(1, 1)      // π/4`,
					},
					{
						Title: "Min / Max / Mod",
						Notes: "Note: Go's built-in min/max (since 1.21) work on ordered types. math.Min/Max exist but have float-specific NaN semantics.",
						Code: `math.Max(1.5, 2.0)   // 2
math.Mod(7, 3)       // 1 — float remainder`,
					},
				},
			},
			{
				Title: "Float classification",
				Examples: []Example{
					{
						Title: "IsNaN, IsInf, Inf, NaN",
						Notes: "NaN is never equal to anything, including itself. Use IsNaN.",
						Code: `x := math.NaN()
fmt.Println(x == x)        // false
fmt.Println(math.IsNaN(x)) // true

y := math.Inf(+1)
fmt.Println(math.IsInf(y, +1))  // true`,
					},
				},
			},
		},
	})
}
