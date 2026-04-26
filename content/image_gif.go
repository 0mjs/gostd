package content

func init() {
	Register(&Package{
		Name:       "image/gif",
		ImportPath: "image/gif",
		Category:   "Image",
		Summary:    "Encode and decode GIF (including animated). Palette-based.",
		Sections: []Section{
			{
				Title: "Single-frame",
				Examples: []Example{
					{Title: "Encode / Decode", Code: `gif.Encode(w, img, nil)
img, _ := gif.Decode(r)`},
				},
			},
			{
				Title: "Animated",
				Examples: []Example{
					{Title: "EncodeAll", Code: `g := &gif.GIF{
    Image: []*image.Paletted{frame1, frame2},
    Delay: []int{10, 10}, // 100ths of a second
    LoopCount: 0,          // infinite
}
gif.EncodeAll(w, g)`},
					{Title: "DecodeAll", Code: `g, _ := gif.DecodeAll(r)
for i, frame := range g.Image {
    fmt.Println(i, g.Delay[i], frame.Bounds())
}`},
				},
			},
		},
	})
}
