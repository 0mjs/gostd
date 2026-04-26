package content

func init() {
	Register(&Package{
		Name:       "net/http/httputil",
		ImportPath: "net/http/httputil",
		Category:   "Networking",
		Summary:    "Utilities on top of net/http: reverse proxy, dump request/response, chunked encoding.",
		Sections: []Section{
			{
				Title: "ReverseProxy — transparent HTTP proxying",
				Examples: []Example{
					{
						Title: "NewSingleHostReverseProxy",
						Notes: "A few lines gets you a real reverse proxy. Wrap with middleware for logging, auth, or header rewriting.",
						Code: `target, _ := url.Parse("http://backend:8080")
proxy := httputil.NewSingleHostReverseProxy(target)

http.Handle("/", proxy)
http.ListenAndServe(":80", nil)`,
					},
					{
						Title: "Customize with Rewrite (1.20+)",
						Code: `proxy := &httputil.ReverseProxy{
    Rewrite: func(r *httputil.ProxyRequest) {
        r.SetURL(target)
        r.Out.Header.Set("X-Forwarded-Host", r.In.Host)
    },
}`,
					},
				},
			},
			{
				Title: "Dumping requests and responses",
				Examples: []Example{
					{
						Title: "DumpRequest / DumpResponse — debug the wire format",
						Code: `b, _ := httputil.DumpRequest(r, true)
fmt.Println(string(b))`,
					},
				},
			},
		},
	})
}
