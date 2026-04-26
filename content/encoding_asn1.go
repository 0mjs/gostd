package content

func init() {
	Register(&Package{
		Name:       "encoding/asn1",
		ImportPath: "encoding/asn1",
		Category:   "Encoding",
		Summary:    "ASN.1 DER encoding — the format certificates, private keys, and X.509 structures live in. Reach for this when interoperating with PKI.",
		Sections: []Section{
			{
				Title: "Marshal / Unmarshal",
				Examples: []Example{
					{
						Title: "Encode a simple struct",
						Code: `type Point struct{ X, Y int }
b, _ := asn1.Marshal(Point{3, 4})
fmt.Printf("% x\n", b)

var p Point
asn1.Unmarshal(b, &p)
fmt.Println(p)`,
					},
					{
						Title: "Struct tags for tagged fields",
						Code: "type Cert struct {\n    Version    int      `asn1:\"optional,explicit,default:1,tag:0\"`\n    Subject    string\n}\n",
					},
				},
			},
		},
	})
}
