package content

func init() {
	Register(&Package{
		Name:       "image/jpeg",
		ImportPath: "image/jpeg",
		Category:   "Image",
		Summary:    "Encode and decode JPEG. Lossy. Import _ \"image/jpeg\" to register it.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Encode with quality", Code: `f, _ := os.Create("out.jpg")
defer f.Close()
jpeg.Encode(f, img, &jpeg.Options{Quality: 85}) // 1..100`},
					{Title: "Decode", Code: `f, _ := os.Open("in.jpg")
defer f.Close()
img, err := jpeg.Decode(f)`},
				},
			},
		},
	})
}
