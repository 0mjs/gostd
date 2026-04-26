package content

func init() {
	Register(&Package{
		Name:       "net/mail",
		ImportPath: "net/mail",
		Category:   "Networking",
		Summary:    "Parse RFC 5322 addresses and message headers. Read emails, not send them (see net/smtp for that).",
		Sections: []Section{
			{
				Title: "Addresses",
				Examples: []Example{
					{
						Title: "ParseAddress / ParseAddressList",
						Code: `addr, _ := mail.ParseAddress("Ada Lovelace <ada@example.com>")
fmt.Println(addr.Name, addr.Address)

list, _ := mail.ParseAddressList("ada@x.com, grace@y.com")
for _, a := range list {
    fmt.Println(a.Address)
}`,
					},
				},
			},
			{
				Title: "Full messages",
				Examples: []Example{
					{
						Title: "ReadMessage",
						Code: `msg, _ := mail.ReadMessage(f)
fmt.Println(msg.Header.Get("From"))
fmt.Println(msg.Header.Get("Subject"))

body, _ := io.ReadAll(msg.Body)
fmt.Println(string(body))`,
					},
				},
			},
		},
	})
}
