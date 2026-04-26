package content

func init() {
	Register(&Package{
		Name:       "image/draw",
		ImportPath: "image/draw",
		Category:   "Image",
		Summary:    "Compose images: copy, blend, mask. The building block for simple image manipulation.",
		Sections: []Section{
			{
				Title: "Draw operations",
				Examples: []Example{
					{Title: "Copy (Src) vs blend (Over)", Code: `draw.Draw(dst, dst.Bounds(), src, image.Point{}, draw.Src)  // replace
draw.Draw(dst, dst.Bounds(), src, image.Point{}, draw.Over) // alpha blend`},
					{Title: "Solid fill", Code: `draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)`},
					{Title: "Masked draw", Code: `draw.DrawMask(dst, r, src, image.Point{}, mask, image.Point{}, draw.Over)`},
				},
			},
		},
	})
}
