package content

func init() {
	Register(&Package{
		Name:       "encoding/pem",
		ImportPath: "encoding/pem",
		Category:   "Encoding",
		Summary:    "The BEGIN/END block format for certificates and keys. Thin layer over base64 with a typed header.",
		Sections: []Section{
			{
				Title: "Decode a PEM block",
				Examples: []Example{
					{
						Title: "Decode — pull out a DER blob",
						Code: `block, _ := pem.Decode(data)
if block == nil || block.Type != "CERTIFICATE" {
    return errors.New("not a certificate PEM block")
}
cert, err := x509.ParseCertificate(block.Bytes)`,
					},
					{
						Title: "Encode — write a PEM block",
						Code: `pem.Encode(os.Stdout, &pem.Block{
    Type:  "CERTIFICATE",
    Bytes: der,
})`,
					},
				},
			},
		},
	})
}
