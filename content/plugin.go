package content

func init() {
	Register(&Package{
		Name:       "plugin",
		ImportPath: "plugin",
		Category:   "Misc",
		Summary:    "Load Go shared objects (.so) at runtime. Linux/macOS only. Fragile — all plugins must match the host's exact toolchain and deps. Rarely the right tool.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Build a plugin", Code: `go build -buildmode=plugin -o hello.so hello.go`},
					{Title: "Open and look up symbols", Code: `p, err := plugin.Open("hello.so")
if err != nil { log.Fatal(err) }

sym, err := p.Lookup("Greet")
if err != nil { log.Fatal(err) }

greet := sym.(func(string) string)
fmt.Println(greet("world"))`},
				},
			},
		},
	})
}
