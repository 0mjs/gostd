package content

func init() {
	Register(&Package{
		Name:       "encoding/csv",
		ImportPath: "encoding/csv",
		Category:   "Encoding",
		Summary:    "Read and write RFC 4180 CSV. Streaming, configurable delimiter, quote-aware.",
		Sections: []Section{
			{
				Title: "Reading",
				Examples: []Example{
					{
						Title: "ReadAll — small files",
						Code: `r := csv.NewReader(strings.NewReader("name,age\nAda,36\nGrace,85\n"))
records, err := r.ReadAll()
if err != nil { log.Fatal(err) }
for _, row := range records {
    fmt.Println(row)
}`,
						Output: `[name age]
[Ada 36]
[Grace 85]
`,
					},
					{
						Title: "Read loop — streaming",
						Notes: "Use when the file is large. Returns io.EOF at the end, one row at a time.",
						Code: `r := csv.NewReader(f)
for {
    rec, err := r.Read()
    if err == io.EOF { break }
    if err != nil { log.Fatal(err) }
    fmt.Println(rec)
}`,
					},
					{
						Title: "Customize delimiter and fields",
						Code: `r := csv.NewReader(f)
r.Comma = ';'                  // semicolon-separated
r.Comment = '#'                // skip comment lines
r.FieldsPerRecord = -1         // allow ragged rows
r.TrimLeadingSpace = true`,
					},
				},
			},
			{
				Title: "Writing",
				Examples: []Example{
					{
						Title: "Writer — always Flush",
						Code: `w := csv.NewWriter(os.Stdout)
w.Write([]string{"name", "age"})
w.Write([]string{"Ada", "36"})
w.Flush()
if err := w.Error(); err != nil { log.Fatal(err) }`,
					},
				},
			},
		},
	})
}
