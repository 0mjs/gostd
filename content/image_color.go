package content

func init() {
	Register(&Package{
		Name:       "image/color",
		ImportPath: "image/color",
		Category:   "Image",
		Summary:    "Color models: RGBA, NRGBA, Gray, YCbCr, etc. Understand pre-multiplied vs non-premultiplied alpha.",
		Sections: []Section{
			{
				Title: "Common colors",
				Examples: []Example{
					{Title: "RGBA (pre-multiplied)", Code: `red := color.RGBA{R: 255, A: 255}       // opaque red
halfGreen := color.RGBA{G: 128, A: 128}  // ⚠ pre-multiplied`},
					{Title: "NRGBA (straight alpha)", Code: `halfGreen := color.NRGBA{G: 255, A: 128}`},
					{Title: "Gray / Gray16", Code: `g := color.Gray{Y: 128}`},
				},
			},
			{
				Title: "Convert between models",
				Examples: []Example{
					{Title: "Model.Convert", Code: `c := color.NRGBAModel.Convert(color.RGBA{R: 255, A: 128})
r, g, b, a := c.RGBA() // always uint32 0..0xffff`},
				},
			},
		},
	})
}
