package content

func init() {
	Register(&Package{
		Name:       "debug/buildinfo",
		ImportPath: "debug/buildinfo",
		Category:   "Runtime & Debug",
		Summary:    "Read build info embedded in a Go binary on disk (module path, version, VCS revision).",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "ReadFile", Code: `info, err := buildinfo.ReadFile("/usr/local/bin/mytool")
if err != nil { log.Fatal(err) }
fmt.Println(info.GoVersion, info.Main.Path, info.Main.Version)
for _, s := range info.Settings {
    if s.Key == "vcs.revision" { fmt.Println("rev:", s.Value) }
}`},
					{Title: "Read", Code: `f, _ := os.Open("mytool")
info, _ := buildinfo.Read(f)`},
				},
			},
		},
	})
}
