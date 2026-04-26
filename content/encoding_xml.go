package content

func init() {
	Register(&Package{
		Name:       "encoding/xml",
		ImportPath: "encoding/xml",
		Category:   "Encoding",
		Summary:    "XML encode/decode via reflection and struct tags. Same model as encoding/json.",
		Sections: []Section{
			{
				Title: "Marshal and Unmarshal",
				Examples: []Example{
					{
						Title: "Struct tags control element/attr mapping",
						Code: "type Book struct {\n    XMLName xml.Name `xml:\"book\"`\n    Title   string   `xml:\"title\"`\n    Pages   int      `xml:\"pages,attr\"`\n}\n\nb, _ := xml.MarshalIndent(Book{Title: \"Go\", Pages: 300}, \"\", \"  \")\nfmt.Println(string(b))\n",
						Output: `<book pages="300">
  <title>Go</title>
</book>
`,
					},
					{
						Title: "chardata, innerxml, omitempty",
						Code: "type P struct {\n    XMLName xml.Name `xml:\"p\"`\n    Text    string   `xml:\",chardata\"`\n}\n",
					},
				},
			},
			{
				Title: "Streaming with Decoder",
				Examples: []Example{
					{
						Title: "Token loop — walk element-by-element",
						Notes: "Use for huge XML that doesn't fit in memory.",
						Code: `dec := xml.NewDecoder(r)
for {
    tok, err := dec.Token()
    if err == io.EOF { break }
    switch t := tok.(type) {
    case xml.StartElement:
        if t.Name.Local == "item" {
            var it Item
            dec.DecodeElement(&it, &t)
        }
    }
}`,
					},
				},
			},
		},
	})
}
