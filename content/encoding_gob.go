package content

func init() {
	Register(&Package{
		Name:       "encoding/gob",
		ImportPath: "encoding/gob",
		Category:   "Encoding",
		Summary:    "Go-native binary serialization. Self-describing, concise, Go-to-Go only. Not a cross-language format.",
		Sections: []Section{
			{
				Title: "Encode and decode over a stream",
				Examples: []Example{
					{
						Title: "Round-trip a struct via bytes.Buffer",
						Code: `type P struct{ X, Y int }

var buf bytes.Buffer
enc := gob.NewEncoder(&buf)
enc.Encode(P{3, 4})

var q P
dec := gob.NewDecoder(&buf)
dec.Decode(&q)
fmt.Println(q)`,
						Output: `{3 4}
`,
					},
					{
						Title: "Interface values — must Register concrete types",
						Notes: "gob needs to know the concrete type behind an interface before it can encode it.",
						Code: `gob.Register(&Circle{})
gob.Register(&Square{})

enc.Encode([]Shape{&Circle{R: 1}, &Square{S: 2}})`,
					},
				},
			},
			{
				Title: "When to use gob vs JSON vs protobuf",
				Description: "Use gob for Go-to-Go state snapshots (caches, disk persistence). Prefer JSON for humans and HTTP APIs. Use protobuf/flatbuffers when you cross languages or need a strict schema.",
			},
		},
	})
}
