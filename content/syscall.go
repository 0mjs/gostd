package content

func init() {
	Register(&Package{
		Name:       "syscall",
		ImportPath: "syscall",
		Category:   "Misc",
		Summary:    "Low-level, platform-specific OS primitives. Frozen — prefer os, os/exec, and golang.org/x/sys for anything new.",
		Sections: []Section{
			{
				Title: "Common uses",
				Examples: []Example{
					{Title: "Set file permissions precisely", Code: `syscall.Umask(0o077) // affects os.Create / OpenFile defaults`},
					{Title: "Process attributes (Unix)", Code: `cmd := exec.Command("srv")
cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
cmd.Start()`},
					{Title: "Signals", Code: `syscall.Kill(pid, syscall.SIGTERM)`},
					{Title: "Errno", Code: `_, err := os.Open("x")
if errors.Is(err, syscall.ENOENT) { /* missing file */ }`},
				},
			},
			{
				Title: "Note",
				Description: "New OS features land in golang.org/x/sys, not here. Use this package only when you can't avoid it.",
			},
		},
	})
}
