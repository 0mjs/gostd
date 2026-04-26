package content

func init() {
	Register(&Package{
		Name:       "crypto/x509",
		ImportPath: "crypto/x509",
		Category:   "Crypto",
		Summary:    "Parse, create, and verify X.509 certificates. The foundation of TLS and PKI.",
		Sections: []Section{
			{
				Title: "Parsing",
				Examples: []Example{
					{
						Title: "Read a PEM cert file",
						Code: `data, _ := os.ReadFile("cert.pem")
block, _ := pem.Decode(data)
cert, err := x509.ParseCertificate(block.Bytes)
if err != nil { log.Fatal(err) }
fmt.Println(cert.Subject, cert.NotAfter)`,
					},
				},
			},
			{
				Title: "Cert pools and verification",
				Examples: []Example{
					{
						Title: "System cert pool",
						Code: `pool, _ := x509.SystemCertPool()
_, err := cert.Verify(x509.VerifyOptions{Roots: pool})`,
					},
				},
			},
			{
				Title: "Creating self-signed certs",
				Examples: []Example{
					{
						Title: "CreateCertificate",
						Code: `template := &x509.Certificate{
    SerialNumber: big.NewInt(1),
    Subject:      pkix.Name{CommonName: "localhost"},
    NotBefore:    time.Now(),
    NotAfter:     time.Now().AddDate(1, 0, 0),
    KeyUsage:     x509.KeyUsageDigitalSignature,
    ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
    DNSNames:     []string{"localhost"},
}
der, _ := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)`,
					},
				},
			},
		},
	})
}
