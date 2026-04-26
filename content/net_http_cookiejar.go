package content

func init() {
	Register(&Package{
		Name:       "net/http/cookiejar",
		ImportPath: "net/http/cookiejar",
		Category:   "Networking",
		Summary:    "An in-memory HTTP cookie store that implements http.CookieJar. Stick it on an http.Client to persist cookies across requests.",
		Sections: []Section{
			{
				Title: "Enable cookies on a client",
				Examples: []Example{
					{
						Title: "Log in once, reuse the session",
						Code: `jar, _ := cookiejar.New(nil)
client := &http.Client{Jar: jar}

client.PostForm("https://example.com/login", url.Values{
    "user": {"ada"}, "password": {"hunter2"},
})
client.Get("https://example.com/account")   // cookies tagged automatically`,
					},
					{
						Title: "Public suffix list for security",
						Notes: "Pass a PublicSuffixList in the Options so cookies can't be set for 'co.uk' and similar.",
					},
				},
			},
		},
	})
}
