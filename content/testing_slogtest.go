package content

func init() {
	Register(&Package{
		Name:       "testing/slogtest",
		ImportPath: "testing/slogtest",
		Category:   "Testing",
		Summary:    "Conformance tests for custom slog.Handler implementations.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "TestHandler", Code: `var buf bytes.Buffer
h := myhandler.New(&buf)
results := func() []map[string]any {
    var out []map[string]any
    for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
        if len(line) == 0 { continue }
        var m map[string]any
        json.Unmarshal(line, &m)
        out = append(out, m)
    }
    return out
}
if err := slogtest.TestHandler(h, results); err != nil {
    t.Fatal(err)
}`},
				},
			},
		},
	})
}
