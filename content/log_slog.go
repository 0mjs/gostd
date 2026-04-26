package content

func init() {
	Register(&Package{
		Name:       "log/slog",
		ImportPath: "log/slog",
		Category:   "Errors & Logging",
		Summary:    "Structured logging (Go 1.21+). Leveled, attributable, and produces text or JSON for free.",
		Sections: []Section{
			{
				Title: "The default logger",
				Examples: []Example{
					{
						Title: "Info, Warn, Error — with key/value attrs",
						Code: `slog.Info("request served",
    "method", "GET",
    "path", "/api/users",
    "status", 200,
)`,
						Output: `2024/03/14 15:09:26 INFO request served method=GET path=/api/users status=200
`,
					},
					{
						Title: "Debug, Warn, Error",
						Notes: "Debug is silenced by default (level INFO). Set Level to change that.",
						Code: `slog.Debug("cache miss", "key", k)
slog.Warn("slow query", "ms", took.Milliseconds())
slog.Error("db write failed", "err", err)`,
					},
				},
			},
			{
				Title: "Handlers — where logs go and how they look",
				Description: "A Logger wraps a Handler. Swap the handler for JSON output or to filter levels.",
				Examples: []Example{
					{
						Title: "JSONHandler — machine-readable",
						Code: `logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("served", "status", 200, "ms", 42)`,
						Output: `{"time":"2024-03-14T15:09:26Z","level":"INFO","msg":"served","status":200,"ms":42}
`,
					},
					{
						Title: "Set the default logger for the whole program",
						Code: `slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
    Level: slog.LevelDebug,
})))`,
					},
				},
			},
			{
				Title: "Attributes and groups",
				Examples: []Example{
					{
						Title: "Typed attrs with slog.Any / slog.Int / slog.String",
						Notes: "Prefer typed attrs in hot paths: no reflection, no allocations for many common cases.",
						Code: `slog.LogAttrs(ctx, slog.LevelInfo, "order placed",
    slog.String("id", orderID),
    slog.Int("items", n),
    slog.Duration("took", d),
)`,
					},
					{
						Title: "With — derive a logger with pinned fields",
						Code: `reqLogger := slog.Default().With("req_id", id, "user", u)
reqLogger.Info("auth ok")
reqLogger.Info("db read", "table", "users")`,
					},
					{
						Title: "Group — nested attributes",
						Code: `logger.Info("user signed up",
    slog.Group("user",
        slog.String("email", email),
        slog.String("plan", "pro"),
    ),
)`,
					},
				},
			},
		},
	})
}
