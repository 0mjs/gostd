package content

func init() {
	Register(&Package{
		Name:       "image/png",
		ImportPath: "image/png",
		Category:   "Image",
		Summary:    "Encode and decode PNG. Loss-less. Import _ \"image/png\" to register it with image.Decode.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Encode", Code: `f, _ := os.Create("out.png")
defer f.Close()
png.Encode(f, img)`},
					{Title: "Decode", Code: `f, _ := os.Open("in.png")
defer f.Close()
img, err := png.Decode(f)`},
					{Title: "Tune compression", Code: `e := &png.Encoder{CompressionLevel: png.BestCompression}
e.Encode(f, img)`},
				},
			},
		},
	})
}
