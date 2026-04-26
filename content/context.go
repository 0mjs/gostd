package content

func init() {
	Register(&Package{
		Name:       "context",
		ImportPath: "context",
		Category:   "Time & Context",
		Summary:    "Cancellation, deadlines, and request-scoped values that propagate across API boundaries and goroutines.",
		Sections: []Section{
			{
				Title: "The four constructors",
				Description: "Every request starts from Background() or TODO(). You then wrap it to add cancellation, deadlines, timeouts, or values.",
				Examples: []Example{
					{
						Title: "Background vs TODO",
						Notes: "Background is the root for normal code. TODO signals 'I don't know what to use here yet' — useful during refactors. They behave identically; the name documents intent.",
						Code: `ctx := context.Background()  // main, init, tests
// ctx := context.TODO()   // placeholder`,
					},
					{
						Title: "WithCancel — caller cancels when done",
						Code: `ctx, cancel := context.WithCancel(context.Background())
defer cancel()          // ALWAYS defer cancel to release resources

go doWork(ctx)
time.Sleep(100 * time.Millisecond)
cancel()                // tell doWork to stop`,
					},
					{
						Title: "WithTimeout and WithDeadline",
						Notes: "Timeout is relative, Deadline is absolute. Both return a cancel that you must defer, even if the deadline fires on its own.",
						Code: `ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()

ctx, cancel = context.WithDeadline(ctx, time.Now().Add(5*time.Second))
defer cancel()`,
					},
					{
						Title: "WithValue — propagate request-scoped data",
						Notes: "Use sparingly. Meant for things like request IDs, auth info, logger — NOT function arguments in disguise. Use a private key type to avoid collisions.",
						Code: `type reqIDKey struct{}

ctx = context.WithValue(ctx, reqIDKey{}, "req-123")
id, _ := ctx.Value(reqIDKey{}).(string)
fmt.Println(id)`,
					},
				},
			},
			{
				Title: "Observing cancellation",
				Examples: []Example{
					{
						Title: "select on ctx.Done()",
						Code: `func doWork(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case job := <-jobs:
            process(job)
        }
    }
}`,
					},
					{
						Title: "ctx.Err — why did we stop?",
						Notes: "Returns nil while alive. After Done, returns context.Canceled or context.DeadlineExceeded.",
						Code: `if err := ctx.Err(); err != nil {
    return fmt.Errorf("abort: %w", err)
}`,
					},
				},
			},
			{
				Title: "Patterns",
				Examples: []Example{
					{
						Title: "Always pass ctx as the first argument",
						Notes: "The idiom: func Do(ctx context.Context, ...). Never store ctx inside a struct; accept it on each call.",
						Code: `func Fetch(ctx context.Context, url string) ([]byte, error) { ... }`,
					},
					{
						Title: "HTTP request with context",
						Code: `req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := http.DefaultClient.Do(req)`,
					},
					{
						Title: "signal.NotifyContext — cancel on SIGINT",
						Code: `ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()

if err := server.Shutdown(ctx); err != nil { ... }`,
					},
				},
			},
		},
	})
}
