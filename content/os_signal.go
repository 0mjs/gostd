package content

func init() {
	Register(&Package{
		Name:       "os/signal",
		ImportPath: "os/signal",
		Category:   "CLI & Runtime",
		Summary:    "Receive OS signals (SIGINT, SIGTERM, etc.) via a channel or context.",
		Sections: []Section{
			{
				Title: "NotifyContext — the modern idiom",
				Examples: []Example{
					{
						Title: "Graceful shutdown",
						Notes: "Cancels ctx when the first matching signal arrives. Call stop() to restore default handling. Pairs perfectly with http.Server.Shutdown.",
						Code: `ctx, stop := signal.NotifyContext(context.Background(),
    os.Interrupt, syscall.SIGTERM)
defer stop()

go srv.ListenAndServe()
<-ctx.Done()

shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
srv.Shutdown(shutdownCtx)`,
					},
				},
			},
			{
				Title: "Notify — raw channel version",
				Examples: []Example{
					{
						Title: "signal.Notify",
						Notes: "Buffered channel of size 1 is the classic pattern — if signals arrive faster than you read, the oldest queued one wins.",
						Code: `ch := make(chan os.Signal, 1)
signal.Notify(ch, os.Interrupt)

<-ch
fmt.Println("bye")`,
					},
					{
						Title: "Stop — unregister",
						Code: `signal.Stop(ch)   // ch will no longer receive signals`,
					},
				},
			},
		},
	})
}
