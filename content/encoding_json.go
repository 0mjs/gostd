package content

func init() {
	Register(&Package{
		Name:       "encoding/json",
		ImportPath: "encoding/json",
		Category:   "Encoding",
		Summary:    "Encode and decode JSON. Works with structs via reflection and tags, or maps/any for dynamic shapes.",
		Sections: []Section{
			{
				Title: "Marshal and Unmarshal",
				Description: "Marshal goes value → []byte. Unmarshal goes []byte → pointer. Always pass a pointer to Unmarshal.",
				Examples: []Example{
					{
						Title: "Struct marshaling with tags",
						Notes: "json:\"name\" renames the field. omitempty skips zero values. A field must be exported (capitalized) to be seen by the encoder.",
						Code: "type User struct {\n    Name string `json:\"name\"`\n    Age  int    `json:\"age,omitempty\"`\n}\n\nb, _ := json.Marshal(User{Name: \"Ada\"})\nfmt.Println(string(b))   // {\"name\":\"Ada\"}\n",
						Output: `{"name":"Ada"}
`,
					},
					{
						Title: "MarshalIndent — pretty-printed",
						Code: `b, _ := json.MarshalIndent(User{Name: "Ada", Age: 36}, "", "  ")
fmt.Println(string(b))`,
						Output: `{
  "name": "Ada",
  "age": 36
}`,
					},
					{
						Title: "Unmarshal into a struct",
						Code: `data := []byte(` + "`" + `{"name":"Ada","age":36}` + "`" + `)
var u User
if err := json.Unmarshal(data, &u); err != nil {
    log.Fatal(err)
}
fmt.Println(u)`,
					},
					{
						Title: "Unmarshal into a map or any",
						Notes: "Use when the shape isn't known. Numbers decode as float64 unless you use json.Decoder.UseNumber().",
						Code: `var v any
json.Unmarshal([]byte(` + "`" + `{"a":1,"b":[2,3]}` + "`" + `), &v)
fmt.Printf("%#v\n", v)`,
					},
				},
			},
			{
				Title: "Streaming with Encoder / Decoder",
				Description: "Use these for JSON over the wire (HTTP, TCP, files). Decoder reads from an io.Reader; Encoder writes to an io.Writer.",
				Examples: []Example{
					{
						Title: "Decoder — stream a JSON array",
						Code: `dec := json.NewDecoder(resp.Body)
_, _ = dec.Token()   // read the opening [
for dec.More() {
    var u User
    if err := dec.Decode(&u); err != nil {
        log.Fatal(err)
    }
    fmt.Println(u)
}`,
					},
					{
						Title: "Encoder — write a JSON response",
						Notes: "Encoder adds a trailing newline after each value, which is what you usually want for HTTP JSON lines.",
						Code: `func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(User{Name: "Ada", Age: 36})
}`,
					},
					{
						Title: "DisallowUnknownFields — strict decoding",
						Notes: "Reject payloads with fields your struct doesn't declare. Great for config files and public APIs.",
						Code: `dec := json.NewDecoder(r)
dec.DisallowUnknownFields()
if err := dec.Decode(&cfg); err != nil {
    return err
}`,
					},
				},
			},
			{
				Title: "Custom encoding",
				Examples: []Example{
					{
						Title: "Implement json.Marshaler / Unmarshaler",
						Code: `type Celsius float64

func (c Celsius) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("\"%.1f°C\"", float64(c))), nil
}

b, _ := json.Marshal(Celsius(22.5))
fmt.Println(string(b))`,
						Output: `"22.5°C"`,
					},
					{
						Title: "RawMessage — defer decoding",
						Notes: "Use when the shape of a sub-object depends on another field (e.g. a 'type' discriminator).",
						Code: "type Envelope struct {\n    Kind    string          `json:\"kind\"`\n    Payload json.RawMessage `json:\"payload\"`\n}\n\nvar env Envelope\njson.Unmarshal(data, &env)\n\nswitch env.Kind {\ncase \"user\":\n    var u User\n    json.Unmarshal(env.Payload, &u)\ncase \"order\":\n    var o Order\n    json.Unmarshal(env.Payload, &o)\n}\n",
					},
				},
			},
			{
				Title: "Valid and compact",
				Examples: []Example{
					{
						Title: "json.Valid",
						Code: `json.Valid([]byte(` + "`" + `{"ok":true}` + "`" + `))    // true
json.Valid([]byte("{oops"))          // false`,
					},
					{
						Title: "json.Indent / json.Compact",
						Code: `var out bytes.Buffer
json.Indent(&out, raw, "", "  ")   // pretty-print
// json.Compact(&out, raw)         // remove whitespace`,
					},
				},
			},
		},
	})
}
