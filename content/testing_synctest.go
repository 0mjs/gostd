package content

func init() {
	Register(&Package{
		Name:       "testing/synctest",
		ImportPath: "testing/synctest",
		Category:   "Testing",
		Summary:    "Test concurrent code with a fake clock and deterministic scheduling. Available since Go 1.24 (GOEXPERIMENT) / stable in 1.25+.",
		Sections: []Section{
			{
				Title: "Bubbles and fake time",
				Examples: []Example{
					{Title: "Test", Code: `func TestTimeout(t *testing.T) {
    synctest.Test(t, func(t *testing.T) {
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        start := time.Now()
        <-ctx.Done()
        if d := time.Since(start); d != time.Second {
            t.Errorf("wanted exactly 1s, got %v", d)
        }
    })
}`},
					{Title: "Wait for all goroutines idle", Code: `synctest.Wait() // inside a bubble — blocks until all bubble goroutines are durably blocked`},
				},
			},
		},
	})
}
