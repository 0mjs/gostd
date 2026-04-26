package content

func init() {
	Register(&Package{
		Name:       "log/syslog",
		ImportPath: "log/syslog",
		Category:   "Errors & Logging",
		Summary:    "Write to system log via the BSD syslog protocol. Unix-only. Frozen — new code should use log/slog + journald / your platform's logger.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Dial + write", Code: `w, err := syslog.Dial("", "", syslog.LOG_INFO|syslog.LOG_LOCAL0, "myapp")
if err != nil { log.Fatal(err) }
w.Info("server started")
w.Err("connection failed")`},
					{Title: "Wire to the standard log package", Code: `w, _ := syslog.New(syslog.LOG_INFO, "myapp")
log.SetOutput(w)`},
				},
			},
		},
	})
}
