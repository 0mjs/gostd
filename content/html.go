package content

func init() {
	Register(&Package{
		Name:       "html",
		ImportPath: "html",
		Category:   "Templates",
		Summary:    "Escape and unescape HTML text. Tiny package — just EscapeString / UnescapeString.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Escape", Code: `s := html.EscapeString(`+"`"+`<a href="x">Ben & Jerry</a>`+"`"+`)
// &lt;a href=&#34;x&#34;&gt;Ben &amp; Jerry&lt;/a&gt;`},
					{Title: "Unescape", Code: `html.UnescapeString("Ben &amp; Jerry") // Ben & Jerry`},
				},
			},
		},
	})
}
