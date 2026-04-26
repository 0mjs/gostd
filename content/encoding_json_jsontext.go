package content

func init() {
	Register(&Package{
		Name:       "encoding/json/jsontext",
		ImportPath: "encoding/json/jsontext",
		Category:   "Encoding",
		Summary:    "Low-level JSON tokenizer/emitter that underpins encoding/json/v2. Stream JSON without building a tree, with precise control over formatting.",
		Sections: []Section{
			{
				Title: "Streaming decode",
				Examples: []Example{
					{Title: "Decoder", Code: `dec := jsontext.NewDecoder(r)
for {
    tok, err := dec.ReadToken()
    if err == io.EOF { break }
    if err != nil { log.Fatal(err) }
    fmt.Println(tok.Kind(), tok.String())
}`},
				},
			},
			{
				Title: "Streaming encode",
				Examples: []Example{
					{Title: "Encoder", Code: `enc := jsontext.NewEncoder(w)
enc.WriteToken(jsontext.BeginObject)
enc.WriteToken(jsontext.String("name"))
enc.WriteToken(jsontext.String("Ada"))
enc.WriteToken(jsontext.EndObject)`},
				},
			},
		},
	})
}
