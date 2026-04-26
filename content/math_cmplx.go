package content

func init() {
	Register(&Package{
		Name:       "math/cmplx",
		ImportPath: "math/cmplx",
		Category:   "Math",
		Summary:    "Elementary functions on the complex128 type: Abs, Exp, Log, Sqrt, Phase, trig.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "Euler's identity in three lines",
						Code: `c := cmplx.Exp(1i * math.Pi)
fmt.Println(c)                    // ~(-1+0i)
fmt.Println(cmplx.Abs(c))         // 1`,
					},
					{
						Title: "Polar form",
						Code: `z := complex(3, 4)
fmt.Println(cmplx.Abs(z), cmplx.Phase(z))   // 5, atan2(4,3)`,
					},
				},
			},
		},
	})
}
