package content

func init() {
	Register(&Package{
		Name:       "go/constant",
		ImportPath: "go/constant",
		Category:   "Go Tooling",
		Summary:    "Arbitrary-precision values used by go/types for untyped constants (ints, rationals, strings, booleans).",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Build and operate", Code: `a := constant.MakeInt64(7)
b := constant.MakeInt64(3)
sum := constant.BinaryOp(a, token.ADD, b) // 10
fmt.Println(constant.Int64Val(sum))`},
				},
			},
		},
	})
}
