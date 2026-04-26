package content

func init() {
	Register(&Package{
		Name:       "text/template",
		ImportPath: "text/template",
		Category:   "Templates",
		Summary:    "Generic text templating. Use html/template instead for anything rendered as HTML (auto-escapes).",
		Sections: []Section{
			{
				Title: "Parse and execute",
				Examples: []Example{
					{Title: "Simple", Code: `t := template.Must(template.New("greet").Parse("hello {{.Name}}\n"))
t.Execute(os.Stdout, map[string]string{"Name": "Ada"})`},
					{Title: "From files", Code: `t := template.Must(template.ParseFiles("layout.tmpl", "body.tmpl"))
t.ExecuteTemplate(os.Stdout, "layout.tmpl", data)`},
					{Title: "From embed.FS", Code: `//go:embed tmpl/*.tmpl
var tmplFS embed.FS
t := template.Must(template.ParseFS(tmplFS, "tmpl/*.tmpl"))`},
				},
			},
			{
				Title: "Actions",
				Examples: []Example{
					{Title: "Conditionals and ranges", Code: `{{if .LoggedIn}}hi {{.Name}}{{else}}sign in{{end}}
{{range .Items}}- {{.}}
{{else}}no items{{end}}
{{with .User}}{{.Email}}{{end}}`},
					{Title: "Pipelines and funcs", Code: `{{.Title | printf "%q"}}
{{len .Items}}
{{index . "key"}}`},
					{Title: "Variables", Code: `{{$name := .Name}}
{{range $i, $v := .Items}}{{$i}}={{$v}} {{end}}`},
				},
			},
			{
				Title: "Custom functions",
				Examples: []Example{
					{Title: "Funcs", Code: `funcs := template.FuncMap{
    "upper": strings.ToUpper,
    "join":  strings.Join,
}
t := template.Must(template.New("x").Funcs(funcs).Parse("{{upper .Name}}"))`},
				},
			},
			{
				Title: "Partials / composition",
				Examples: []Example{
					{Title: "define + template", Code: `{{define "header"}}<h1>{{.Title}}</h1>{{end}}
{{template "header" .}}

{{block "content" .}}default content{{end}}`},
				},
			},
		},
	})
}
