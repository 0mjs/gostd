package content

func init() {
	Register(&Package{
		Name:       "flag",
		ImportPath: "flag",
		Category:   "CLI & Runtime",
		Summary:    "Command-line flag parsing. Built in, zero-dependency, handles the 90% case.",
		Sections: []Section{
			{
				Title: "Basic flags",
				Examples: []Example{
					{
						Title: "String, Int, Bool",
						Notes: "Each returns a pointer. After flag.Parse(), dereference to read the value.",
						Code: `var (
    addr    = flag.String("addr", ":8080", "listen address")
    timeout = flag.Duration("timeout", 30*time.Second, "request timeout")
    debug   = flag.Bool("debug", false, "enable debug logging")
)

func main() {
    flag.Parse()
    fmt.Println(*addr, *timeout, *debug)
}`,
					},
					{
						Title: "Var forms — bind to existing variables",
						Code: `var port int
flag.IntVar(&port, "port", 8080, "port to listen on")`,
					},
				},
			},
			{
				Title: "Positional arguments",
				Examples: []Example{
					{
						Title: "flag.Args and flag.Arg",
						Notes: "Everything after the flags lives in flag.Args(). flag.NArg() is the count.",
						Code: `flag.Parse()
for i, a := range flag.Args() {
    fmt.Printf("positional %d: %s\n", i, a)
}`,
					},
				},
			},
			{
				Title: "Subcommands with FlagSet",
				Examples: []Example{
					{
						Title: "One FlagSet per subcommand",
						Notes: "flag has no native subcommand concept — use one FlagSet per command and dispatch on os.Args[1].",
						Code: `serve := flag.NewFlagSet("serve", flag.ExitOnError)
addr := serve.String("addr", ":8080", "")

migrate := flag.NewFlagSet("migrate", flag.ExitOnError)
target := migrate.String("to", "latest", "")

switch os.Args[1] {
case "serve":
    serve.Parse(os.Args[2:])
    runServe(*addr)
case "migrate":
    migrate.Parse(os.Args[2:])
    runMigrate(*target)
}`,
					},
				},
			},
			{
				Title: "Custom Value types",
				Examples: []Example{
					{
						Title: "flag.Value interface — repeatable -tag flags",
						Code: `type stringList []string
func (s *stringList) String() string     { return strings.Join(*s, ",") }
func (s *stringList) Set(v string) error { *s = append(*s, v); return nil }

var tags stringList
flag.Var(&tags, "tag", "repeatable")

// $ app -tag a -tag b -tag c`,
					},
				},
			},
		},
	})
}
