package content

func init() {
	Register(&Package{
		Name:       "net/http/httptest",
		ImportPath: "net/http/httptest",
		Category:   "Networking",
		Summary:    "Test helpers for HTTP handlers and servers. Fake requests, record responses, spin up local servers.",
		Sections: []Section{
			{
				Title: "Calling a handler without a real network",
				Examples: []Example{
					{
						Title: "Recorder — record what the handler wrote",
						Code: `rec := httptest.NewRecorder()
req := httptest.NewRequest("GET", "/hello?name=ada", nil)

myHandler(rec, req)

if rec.Code != 200 {
    t.Errorf("got %d", rec.Code)
}
if !strings.Contains(rec.Body.String(), "ada") {
    t.Error("missing ada")
}`,
					},
				},
			},
			{
				Title: "Spinning up a real local server",
				Examples: []Example{
					{
						Title: "NewServer — integration tests",
						Notes: "Gives you a live URL you can pass to clients. Close when done.",
						Code: `srv := httptest.NewServer(http.HandlerFunc(myHandler))
defer srv.Close()

resp, _ := http.Get(srv.URL + "/api/things")`,
					},
					{
						Title: "NewTLSServer — HTTPS variant",
						Notes: "srv.Client() returns a client that trusts the server's self-signed cert.",
						Code: `srv := httptest.NewTLSServer(handler)
defer srv.Close()
resp, _ := srv.Client().Get(srv.URL)`,
					},
				},
			},
		},
	})
}
