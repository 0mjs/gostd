package content

func init() {
	Register(&Package{
		Name:       "crypto/tls",
		ImportPath: "crypto/tls",
		Category:   "Crypto",
		Summary:    "TLS 1.3 and 1.2 client and server. Sits under net/http for HTTPS, but you can use it directly over any net.Conn.",
		Sections: []Section{
			{
				Title: "Server — ListenAndServeTLS",
				Examples: []Example{
					{
						Title: "Minimal HTTPS",
						Code: `http.HandleFunc("/", handler)
log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))`,
					},
					{
						Title: "Explicit TLS config — set MinVersion",
						Notes: "Go's default TLS config is already strong, but pinning MinVersion: TLS 1.2 is a good defensive default.",
						Code: `srv := &http.Server{
    Addr: ":443",
    TLSConfig: &tls.Config{
        MinVersion: tls.VersionTLS12,
    },
}
srv.ListenAndServeTLS("cert.pem", "key.pem")`,
					},
				},
			},
			{
				Title: "Client — use a pool of CAs",
				Examples: []Example{
					{
						Title: "Talking to a server with a custom CA",
						Code: `caCert, _ := os.ReadFile("ca.pem")
pool := x509.NewCertPool()
pool.AppendCertsFromPEM(caCert)

client := &http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{RootCAs: pool},
    },
}`,
					},
					{
						Title: "Load a cert and key",
						Code: `cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
if err != nil { log.Fatal(err) }
cfg := &tls.Config{Certificates: []tls.Certificate{cert}}`,
					},
				},
			},
			{
				Title: "Low-level — TLS over any net.Conn",
				Examples: []Example{
					{
						Title: "Dial a raw TLS endpoint",
						Code: `conn, err := tls.Dial("tcp", "example.com:443", &tls.Config{ServerName: "example.com"})
if err != nil { log.Fatal(err) }
defer conn.Close()
conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))`,
					},
				},
			},
		},
	})
}
