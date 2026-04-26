package content

func init() {
	Register(&Package{
		Name:       "net/url",
		ImportPath: "net/url",
		Category:   "Networking",
		Summary:    "Parse, build, and escape URLs. Handles query strings, paths, and userinfo correctly.",
		Sections: []Section{
			{
				Title: "Parse and rebuild",
				Examples: []Example{
					{
						Title: "url.Parse",
						Code: `u, _ := url.Parse("https://user:pw@example.com:8080/a/b?x=1&y=2#frag")
fmt.Println(u.Scheme, u.Host, u.Path, u.RawQuery, u.Fragment)
fmt.Println(u.User.Username())`,
						Output: `https example.com:8080 /a/b x=1&y=2 frag
user
`,
					},
					{
						Title: "Build from parts",
						Code: `u := &url.URL{
    Scheme: "https",
    Host:   "example.com",
    Path:   "/search",
    RawQuery: url.Values{"q": {"go maps"}, "n": {"20"}}.Encode(),
}
fmt.Println(u.String())`,
						Output: `https://example.com/search?n=20&q=go+maps
`,
					},
					{
						Title: "Resolve a relative reference",
						Notes: "ResolveReference does what a browser does when following a relative link.",
						Code: `base, _ := url.Parse("https://a.com/docs/")
rel, _ := url.Parse("../x.html")
fmt.Println(base.ResolveReference(rel))`,
						Output: `https://a.com/x.html
`,
					},
				},
			},
			{
				Title: "Query values",
				Examples: []Example{
					{
						Title: "url.Values — repeated-key map",
						Code: `v := url.Values{}
v.Set("name", "ada")
v.Add("tag", "a")
v.Add("tag", "b")
fmt.Println(v.Encode())`,
						Output: `name=ada&tag=a&tag=b
`,
					},
				},
			},
			{
				Title: "Escaping",
				Examples: []Example{
					{
						Title: "QueryEscape / PathEscape",
						Notes: "They differ: QueryEscape encodes spaces as '+', PathEscape as '%20'. Use the right one for where the value will live.",
						Code: `url.QueryEscape("a b/c?")   // "a+b%2Fc%3F"
url.PathEscape("a b/c?")    // "a%20b%2Fc%3F"`,
					},
				},
			},
		},
	})
}
