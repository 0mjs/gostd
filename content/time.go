package content

func init() {
	Register(&Package{
		Name:       "time",
		ImportPath: "time",
		Category:   "Time & Context",
		Summary:    "Time, durations, timers, tickers, parsing and formatting. The single source of truth for 'now'.",
		Sections: []Section{
			{
				Title: "Now, durations, arithmetic",
				Examples: []Example{
					{
						Title: "time.Now and monotonic clock",
						Notes: "Now() embeds both a wall-clock reading and a monotonic reading. Subtracting two Times from the same process uses the monotonic part — immune to clock changes.",
						Code: `now := time.Now()
fmt.Println(now)
fmt.Println(now.Unix())          // seconds since 1970
fmt.Println(now.UnixMilli())     // ms since 1970`,
					},
					{
						Title: "Duration constants",
						Notes: "time.Duration is an int64 of nanoseconds. The constants (time.Second, time.Millisecond, ...) let you express durations readably.",
						Code: `d := 2*time.Hour + 30*time.Minute
fmt.Println(d)                    // "2h30m0s"
fmt.Println(d.Minutes())          // 150
fmt.Println(time.Minute / time.Second) // 60`,
						Output: `2h30m0s
150
60
`,
					},
					{
						Title: "Add, Sub, Before/After/Equal",
						Code: `now := time.Now()
later := now.Add(1 * time.Hour)
diff := later.Sub(now)               // 1h0m0s
fmt.Println(later.After(now))        // true
fmt.Println(diff)`,
					},
					{
						Title: "Since and Until — shortcuts for common cases",
						Notes: "time.Since(t) == time.Now().Sub(t). time.Until(t) == t.Sub(time.Now()).",
						Code: `start := time.Now()
expensiveWork()
fmt.Println("took", time.Since(start))`,
					},
				},
			},
			{
				Title: "Formatting and parsing",
				Description: "Go uses a reference date Mon Jan 2 15:04:05 MST 2006 (01/02 03:04:05PM '06 -0700) instead of format directives. Memorize it once.",
				Examples: []Example{
					{
						Title: "Format with a layout — reference date",
						Code: `now := time.Date(2024, 3, 14, 15, 9, 26, 0, time.UTC)
fmt.Println(now.Format("2006-01-02 15:04:05"))
fmt.Println(now.Format(time.RFC3339))`,
						Output: `2024-03-14 15:09:26
2024-03-14T15:09:26Z
`,
					},
					{
						Title: "Parse — must match the layout exactly",
						Code: `t, err := time.Parse("2006-01-02", "2024-03-14")
fmt.Println(t, err)`,
						Output: `2024-03-14 00:00:00 +0000 UTC <nil>
`,
					},
					{
						Title: "Common built-in layouts",
						Code: `// time.RFC3339 = "2006-01-02T15:04:05Z07:00"
// time.DateOnly = "2006-01-02"
// time.TimeOnly = "15:04:05"
// time.DateTime = "2006-01-02 15:04:05"
// time.Kitchen  = "3:04PM"`,
					},
				},
			},
			{
				Title: "Timers, tickers, Sleep",
				Examples: []Example{
					{
						Title: "time.Sleep — block for a duration",
						Code: `time.Sleep(200 * time.Millisecond)`,
					},
					{
						Title: "time.After — one-shot channel fire",
						Notes: "Useful in selects for timeouts. For cancellable timeouts, prefer context.WithTimeout.",
						Code: `select {
case msg := <-ch:
    fmt.Println("got", msg)
case <-time.After(2 * time.Second):
    fmt.Println("timeout")
}`,
					},
					{
						Title: "time.NewTimer vs time.NewTicker",
						Notes: "Timer fires once on C. Ticker fires repeatedly on C until Stop. Always Stop tickers you're done with or they leak goroutines.",
						Code: `t := time.NewTicker(500 * time.Millisecond)
defer t.Stop()
for i := 0; i < 3; i++ {
    <-t.C
    fmt.Println("tick", i)
}`,
					},
					{
						Title: "AfterFunc — run a callback later",
						Code: `timer := time.AfterFunc(1*time.Second, func() {
    fmt.Println("fired")
})
// timer.Stop() would cancel before it fires`,
					},
				},
			},
			{
				Title: "Zones and clocks",
				Examples: []Example{
					{
						Title: "In a specific time zone",
						Code: `loc, _ := time.LoadLocation("Europe/London")
t := time.Now().In(loc)
fmt.Println(t.Format(time.RFC3339))`,
					},
					{
						Title: "Truncate and Round — snap to intervals",
						Code: `t := time.Now()
fmt.Println(t.Truncate(time.Hour))  // start of the current hour
fmt.Println(t.Round(time.Minute))`,
					},
				},
			},
		},
	})
}
