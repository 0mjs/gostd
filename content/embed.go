package content

func init() {
	Register(&Package{
		Name:       "embed",
		ImportPath: "embed",
		Category:   "I/O & Files",
		Summary:    "Embed files and directories into your binary at build time. No external assets at runtime.",
		Sections: []Section{
			{
				Title: "The three forms",
				Examples: []Example{
					{
						Title: "Embed a single file as a string",
						Code: `import _ "embed"

//go:embed version.txt
var version string

fmt.Println(version)`,
					},
					{
						Title: "Embed as []byte — binary assets",
						Code: `//go:embed logo.png
var logo []byte`,
					},
					{
						Title: "Embed a tree as embed.FS — io/fs.FS compatible",
						Notes: "Pass to http.FileServerFS, template.ParseFS, fs.Sub, fs.WalkDir — anywhere that takes an fs.FS.",
						Code: `import "embed"

//go:embed templates static
var assets embed.FS

tmpl := template.Must(template.ParseFS(assets, "templates/*.html"))

static, _ := fs.Sub(assets, "static")
mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(static)))`,
					},
				},
			},
			{
				Title: "Patterns and gotchas",
				Examples: []Example{
					{
						Title: "Multiple patterns",
						Code: `//go:embed *.sql migrations/*
var schema embed.FS`,
					},
					{
						Title: "The directive must be on a package-level var",
						Notes: "The line just above the var must be the //go:embed directive. No blank line between them.",
					},
					{
						Title: "Hidden files excluded by default",
						Notes: "Files/dirs starting with . or _ are ignored. Use all: prefix to include them: //go:embed all:assets.",
					},
				},
			},
		},
	})
}
