package content

func init() {
	Register(&Package{
		Name:       "math/big",
		ImportPath: "math/big",
		Category:   "Math",
		Summary:    "Arbitrary-precision integers (Int), rationals (Rat), and floats (Float). No overflow, at the cost of allocation.",
		Sections: []Section{
			{
				Title: "Int — arbitrary-precision integers",
				Description: "Go operators don't work — you call methods that write into a receiver. The result is always the receiver, which is returned for chaining.",
				Examples: []Example{
					{
						Title: "Factorial — canonical big.Int example",
						Code: `fact := big.NewInt(1)
for i := int64(1); i <= 100; i++ {
    fact.Mul(fact, big.NewInt(i))
}
fmt.Println(fact.String())`,
					},
					{
						Title: "Modular exponentiation for crypto-ish math",
						Code: `result := new(big.Int).Exp(base, exp, mod)`,
					},
					{
						Title: "Parse from a string",
						Code: `n, ok := new(big.Int).SetString("123456789012345678901234567890", 10)
fmt.Println(n, ok)`,
					},
				},
			},
			{
				Title: "Rat — arbitrary-precision rationals",
				Examples: []Example{
					{
						Title: "Exact fractions",
						Code: `r := big.NewRat(1, 3)
r.Add(r, big.NewRat(1, 6))
fmt.Println(r.String())    // 1/2`,
					},
				},
			},
			{
				Title: "Float — arbitrary-precision floating-point",
				Examples: []Example{
					{
						Title: "High-precision math",
						Code: `f := new(big.Float).SetPrec(200).SetFloat64(1.0)
f.Quo(f, big.NewFloat(7))
fmt.Println(f.Text('f', 50))`,
					},
				},
			},
		},
	})
}
