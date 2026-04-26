package content

func init() {
	Register(&Package{
		Name:       "net/textproto",
		ImportPath: "net/textproto",
		Category:   "Networking",
		Summary:    "Implements generic text protocols in the SMTP/HTTP/NNTP style: numbered command/response, MIME headers, pipelining.",
		Sections: []Section{
			{
				Title: "When to reach for it",
				Description: "You're writing a client or server for a text protocol that's not HTTP. textproto handles the fiddly parts: canonicalized header names, reading folded header lines, sending/reading responses with continuation lines.",
			},
			{
				Title: "MIMEHeader",
				Examples: []Example{
					{
						Title: "CanonicalMIMEHeaderKey",
						Notes: "HTTP, SMTP, news all canonicalize header names the same way. http.Header is really a textproto.MIMEHeader.",
						Code: `textproto.CanonicalMIMEHeaderKey("content-type")  // "Content-Type"`,
					},
				},
			},
		},
	})
}
