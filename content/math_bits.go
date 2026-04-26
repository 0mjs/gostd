package content

func init() {
	Register(&Package{
		Name:       "math/bits",
		ImportPath: "math/bits",
		Category:   "Math",
		Summary:    "Low-level bit counting, shifting, leading/trailing zeros. Uses CPU instructions where available — very fast.",
		Sections: []Section{
			{
				Title: "Counting",
				Examples: []Example{
					{
						Title: "OnesCount, LeadingZeros, TrailingZeros",
						Code: `bits.OnesCount64(0b10110101)   // 5
bits.LeadingZeros8(1)          // 7
bits.TrailingZeros32(8)        // 3`,
					},
					{
						Title: "Len — minimum bits needed to represent n",
						Code: `bits.Len(0)    // 0
bits.Len(1)    // 1
bits.Len(255)  // 8`,
					},
				},
			},
			{
				Title: "Rotations and byte swap",
				Examples: []Example{
					{
						Title: "RotateLeft, ReverseBytes",
						Code: `bits.RotateLeft32(1, 4)           // 16
bits.ReverseBytes32(0x11223344)   // 0x44332211`,
					},
				},
			},
			{
				Title: "Math with overflow awareness",
				Description: "Add64, Sub64, Mul64, Div64 return carry/borrow/hi halves. Use when implementing big-int primitives or checksums.",
				Examples: []Example{
					{
						Title: "128-bit multiply from two 64-bit values",
						Code: `hi, lo := bits.Mul64(math.MaxUint64, 2)
fmt.Println(hi, lo)`,
					},
				},
			},
		},
	})
}
