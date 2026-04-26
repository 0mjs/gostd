package content

func init() {
	Register(&Package{
		Name:       "text/tabwriter",
		ImportPath: "text/tabwriter",
		Category:   "Formatting & Strings",
		Summary:    "Aligned columns in plain text. Feed it tab-separated lines; it pads so columns line up.",
		Sections: []Section{
			{
				Title: "Writer — stream tab-separated into aligned columns",
				Examples: []Example{
					{
						Title: "Basic usage — always call Flush",
						Code: `w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
fmt.Fprintln(w, "NAME\tAGE\tROLE")
fmt.Fprintln(w, "Ada\t36\tengineer")
fmt.Fprintln(w, "Grace\t85\tadmiral")
w.Flush()`,
						Output: `NAME   AGE  ROLE
Ada    36   engineer
Grace  85   admiral
`,
					},
					{
						Title: "Flags — right-align, debug separators",
						Code: `// minwidth=0, tabwidth=8, padding=1, padchar='.', flags=tabwriter.AlignRight
w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '.', tabwriter.AlignRight)`,
					},
				},
			},
		},
	})
}
