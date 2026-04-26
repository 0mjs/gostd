package content

func init() {
	Register(&Package{
		Name:       "errors",
		ImportPath: "errors",
		Category:   "Errors & Logging",
		Summary:    "Error inspection, wrapping, and sentinel creation. The counterpart to fmt.Errorf(\"%w\", ...).",
		Sections: []Section{
			{
				Title: "Creating errors",
				Examples: []Example{
					{
						Title: "errors.New — plain sentinel",
						Code: `var ErrNotFound = errors.New("not found")

if _, err := lookup(k); errors.Is(err, ErrNotFound) {
    // handle
}`,
					},
					{
						Title: "fmt.Errorf — formatted and wrapping",
						Notes: "Use %w to wrap. %v formats without wrapping. You can wrap multiple errors with several %w verbs (1.20+).",
						Code: `err := fmt.Errorf("open %s: %w", path, os.ErrNotExist)
fmt.Println(err)

// Multiple wrapping (Go 1.20+)
err = fmt.Errorf("%w; %w", io.EOF, errors.New("extra"))`,
					},
					{
						Title: "errors.Join — combine independent errors",
						Notes: "Useful when validating many fields or fanning out to several goroutines.",
						Code: `err := errors.Join(
    errors.New("name required"),
    errors.New("age invalid"),
)
fmt.Println(err)`,
					},
				},
			},
			{
				Title: "Inspecting errors",
				Examples: []Example{
					{
						Title: "errors.Is — sentinel comparison through wrappers",
						Notes: "Use Is instead of == when checking for known sentinel errors — works even if they were wrapped with %w.",
						Code: `if errors.Is(err, os.ErrNotExist) {
    return newFile()
}`,
					},
					{
						Title: "errors.As — extract a specific error type",
						Notes: "As assigns into its target if any error in the chain matches that type.",
						Code: `var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("path was", pathErr.Path)
}`,
					},
					{
						Title: "errors.Unwrap — one layer off",
						Notes: "Usually you want Is/As instead. Unwrap is for custom error walkers.",
					},
				},
			},
			{
				Title: "Custom error types",
				Examples: []Example{
					{
						Title: "Implement Error() and optionally Unwrap()",
						Code: `type HTTPError struct {
    Code int
    Msg  string
    Err  error
}
func (e *HTTPError) Error() string { return fmt.Sprintf("%d: %s", e.Code, e.Msg) }
func (e *HTTPError) Unwrap() error { return e.Err }`,
					},
				},
			},
		},
	})
}
