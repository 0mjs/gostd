package content

func init() {
	Register(&Package{
		Name:       "image/color/palette",
		ImportPath: "image/color/palette",
		Category:   "Image",
		Summary:    "Pre-built palettes: WebSafe (216 colors) and Plan9 (256 colors). Useful for GIF encoding.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "WebSafe / Plan9", Code: `p := palette.Plan9       // 256-color palette
p2 := palette.WebSafe    // 216-color 6×6×6 cube
img := image.NewPaletted(image.Rect(0,0,100,100), p)`},
				},
			},
		},
	})
}
