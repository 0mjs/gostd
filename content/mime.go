package content

func init() {
	Register(&Package{
		Name:       "mime",
		ImportPath: "mime",
		Category:   "Networking",
		Summary:    "MIME types: look up by extension, parse Content-Type parameters, and encode/decode RFC 2047 encoded-words.",
		Sections: []Section{
			{
				Title: "Type lookup",
				Examples: []Example{
					{
						Title: "TypeByExtension",
						Code: `fmt.Println(mime.TypeByExtension(".pdf"))   // application/pdf
fmt.Println(mime.TypeByExtension(".go"))    // text/plain; charset=utf-8`,
					},
					{
						Title: "ExtensionsByType",
						Code: `exts, _ := mime.ExtensionsByType("image/jpeg")
fmt.Println(exts)   // [.jfif .jpe .jpeg .jpg]`,
					},
					{
						Title: "AddExtensionType — register a custom mapping",
						Code: `mime.AddExtensionType(".myapp", "application/vnd.myapp+json")`,
					},
				},
			},
			{
				Title: "Parsing Content-Type",
				Examples: []Example{
					{
						Title: "ParseMediaType",
						Code: `mt, params, _ := mime.ParseMediaType("multipart/form-data; boundary=xyz")
fmt.Println(mt, params["boundary"])`,
					},
					{
						Title: "FormatMediaType — build one back",
						Code: `s := mime.FormatMediaType("text/plain", map[string]string{"charset": "utf-8"})
// "text/plain; charset=utf-8"`,
					},
				},
			},
		},
	})
}
