package content

func init() {
	Register(&Package{
		Name:       "image",
		ImportPath: "image",
		Category:   "Image",
		Summary:    "Core image types: Image interface, Rectangle, Point, and concrete types like RGBA / Gray / NRGBA.",
		Sections: []Section{
			{
				Title: "Create an image",
				Examples: []Example{
					{Title: "NewRGBA", Code: `img := image.NewRGBA(image.Rect(0, 0, 200, 100))
img.Set(10, 10, color.RGBA{R: 255, A: 255})`},
					{Title: "NewGray / NewNRGBA", Code: `gray := image.NewGray(image.Rect(0, 0, 64, 64))
nrgba := image.NewNRGBA(image.Rect(0, 0, 64, 64))`},
				},
			},
			{
				Title: "Decode (format-agnostic)",
				Description: "Import the format packages for their side effects to auto-register decoders.",
				Examples: []Example{
					{Title: "Decode any registered format", Code: `import (
    _ "image/png"
    _ "image/jpeg"
    _ "image/gif"
)

f, _ := os.Open("input")
defer f.Close()
img, format, err := image.Decode(f)
fmt.Println(format) // "png" / "jpeg" / "gif"`},
					{Title: "DecodeConfig (no pixels)", Code: `cfg, _, _ := image.DecodeConfig(f)
fmt.Println(cfg.Width, cfg.Height)`},
				},
			},
			{
				Title: "Geometry",
				Examples: []Example{
					{Title: "Rectangle / Point", Code: `r := image.Rect(0, 0, 100, 50)
r.Dx()      // width
r.Dy()      // height
r.Intersect(image.Rect(50, 0, 150, 100))`},
				},
			},
		},
	})
}
