package content

func init() {
	Register(&Package{
		Name:       "html/template",
		ImportPath: "html/template",
		Category:   "Templates",
		Summary:    "Same syntax as text/template, but context-aware auto-escaping for HTML, JS, CSS, URL attributes. Use this for web output.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Auto-escaping", Code: `t := template.Must(template.New("p").Parse("<p>{{.}}</p>"))
t.Execute(os.Stdout, "<script>alert(1)</script>")
// Output: <p>&lt;script&gt;alert(1)&lt;/script&gt;</p>`},
					{Title: "Trusted raw HTML", Code: `// Bypass escaping only for values you control:
t.Execute(os.Stdout, template.HTML("<b>bold</b>"))
// Other types: template.JS, template.CSS, template.URL, template.HTMLAttr`},
				},
			},
			{
				Title: "Context matters",
				Description: "Escaping is chosen based on where the value appears: attribute, JS literal, URL, etc.",
				Examples: []Example{
					{Title: "URL context", Code: `<a href="{{.URL}}">click</a>
{{/* template refuses unsafe schemes like javascript:... */}}`},
				},
			},
			{
				Title: "Parse + serve",
				Examples: []Example{
					{Title: "With http", Code: `t := template.Must(template.ParseFS(tmplFS, "tmpl/*.html"))
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    t.ExecuteTemplate(w, "index.html", data)
})`},
				},
			},
		},
	})
}
