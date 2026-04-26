package content

func init() {
	Register(&Package{
		Name:       "net/http/httptrace",
		ImportPath: "net/http/httptrace",
		Category:   "Networking",
		Summary:    "Observe the lifecycle of an HTTP request: DNS lookup, connect, TLS handshake, first byte.",
		Sections: []Section{
			{
				Title: "Hook into the stages",
				Examples: []Example{
					{
						Title: "Measure each phase",
						Code: `trace := &httptrace.ClientTrace{
    DNSStart:       func(i httptrace.DNSStartInfo)       { fmt.Println("DNS start", i.Host) },
    DNSDone:        func(i httptrace.DNSDoneInfo)        { fmt.Println("DNS done") },
    ConnectStart:   func(network, addr string)           { fmt.Println("connecting to", addr) },
    ConnectDone:    func(network, addr string, err error){ fmt.Println("connected") },
    GotFirstResponseByte: func()                         { fmt.Println("first byte") },
}
req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

http.DefaultTransport.RoundTrip(req)`,
					},
				},
			},
		},
	})
}
