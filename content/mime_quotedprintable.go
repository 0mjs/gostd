package content

func init() {
	Register(&Package{
		Name:       "mime/quotedprintable",
		ImportPath: "mime/quotedprintable",
		Category:   "Networking",
		Summary:    "Quoted-printable encoding, from email bodies. =3D for =, =20 for space, soft line breaks.",
		Sections: []Section{
			{
				Title: "Encode and decode",
				Examples: []Example{
					{
						Title: "Decode a QP body",
						Code: `r := quotedprintable.NewReader(strings.NewReader("Hello=2C=20world!"))
b, _ := io.ReadAll(r)
fmt.Println(string(b))   // "Hello, world!"`,
					},
					{
						Title: "Encode",
						Code: `var buf bytes.Buffer
w := quotedprintable.NewWriter(&buf)
w.Write([]byte("Héllo, world!"))
w.Close()
fmt.Println(buf.String())`,
					},
				},
			},
		},
	})
}
