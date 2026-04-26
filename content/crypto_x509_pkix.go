package content

func init() {
	Register(&Package{
		Name:       "crypto/x509/pkix",
		ImportPath: "crypto/x509/pkix",
		Category:   "Crypto",
		Summary:    "ASN.1 structures used inside X.509: Name, RDNSequence, CertificateList. Mostly used alongside crypto/x509 types.",
		Sections: []Section{
			{
				Title: "Name — subject/issuer fields",
				Examples: []Example{
					{
						Title: "Build a subject",
						Code: `name := pkix.Name{
    CommonName:   "example.com",
    Organization: []string{"Example Inc"},
    Country:      []string{"US"},
}`,
					},
				},
			},
		},
	})
}
