package content

func init() {
	Register(&Package{
		Name:       "os/exec",
		ImportPath: "os/exec",
		Category:   "CLI & Runtime",
		Summary:    "Run external commands. Connect their stdin/stdout to pipes or buffers, pass env and context, capture output.",
		Sections: []Section{
			{
				Title: "Simple runs",
				Examples: []Example{
					{
						Title: "Output — run and capture stdout",
						Code: `out, err := exec.Command("git", "rev-parse", "HEAD").Output()
if err != nil { log.Fatal(err) }
fmt.Println(strings.TrimSpace(string(out)))`,
					},
					{
						Title: "CombinedOutput — stdout + stderr together",
						Notes: "Good for 'just show me what happened' error reporting.",
						Code: `out, err := exec.Command("go", "test", "./...").CombinedOutput()
if err != nil {
    fmt.Fprintln(os.Stderr, string(out))
    os.Exit(1)
}`,
					},
					{
						Title: "Run — don't capture, just wait for exit",
						Code: `cmd := exec.Command("make", "build")
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr
if err := cmd.Run(); err != nil { log.Fatal(err) }`,
					},
				},
			},
			{
				Title: "Context and timeouts",
				Examples: []Example{
					{
						Title: "CommandContext — kill on cancel",
						Notes: "Signals go to the child when ctx is canceled. Use this instead of spawning and forgetting.",
						Code: `ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

cmd := exec.CommandContext(ctx, "curl", url)
if err := cmd.Run(); err != nil { log.Fatal(err) }`,
					},
				},
			},
			{
				Title: "Streaming I/O",
				Examples: []Example{
					{
						Title: "StdinPipe and StdoutPipe",
						Code: `cmd := exec.Command("sort")
stdin, _ := cmd.StdinPipe()
stdout, _ := cmd.StdoutPipe()

cmd.Start()
go func() {
    defer stdin.Close()
    fmt.Fprintln(stdin, "c\nb\na")
}()
b, _ := io.ReadAll(stdout)
cmd.Wait()
fmt.Print(string(b))`,
					},
				},
			},
			{
				Title: "Environment and working directory",
				Examples: []Example{
					{
						Title: "Custom env and cwd",
						Code: `cmd := exec.Command("bin/tool")
cmd.Dir = "/srv/app"
cmd.Env = append(os.Environ(), "FOO=bar")`,
					},
					{
						Title: "LookPath — resolve a command in PATH",
						Code: `bin, err := exec.LookPath("git")
if err != nil { log.Fatal("git not installed") }
fmt.Println(bin)`,
					},
				},
			},
			{
				Title: "Detecting exit codes",
				Examples: []Example{
					{
						Title: "ExitError",
						Code: `if err := cmd.Run(); err != nil {
    var ee *exec.ExitError
    if errors.As(err, &ee) {
        fmt.Println("exited with", ee.ExitCode())
    }
}`,
					},
				},
			},
		},
	})
}
