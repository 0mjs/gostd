package content

func init() {
	Register(&Package{
		Name:       "net/smtp",
		ImportPath: "net/smtp",
		Category:   "Networking",
		Summary:    "Simple SMTP client. Frozen — the maintainers recommend an external package for anything non-trivial.",
		Sections: []Section{
			{
				Title: "Send a single message",
				Examples: []Example{
					{
						Title: "SendMail with PLAIN auth",
						Code: `auth := smtp.PlainAuth("", "user@example.com", "pw", "smtp.example.com")
err := smtp.SendMail(
    "smtp.example.com:587",
    auth,
    "from@example.com",
    []string{"to@example.com"},
    []byte("Subject: hi\r\n\r\nhello"),
)`,
					},
				},
			},
			{
				Title: "Caveat",
				Description: "For reliable delivery (DKIM, retries, MIME attachments) use a third-party library or a transactional email API.",
			},
		},
	})
}
