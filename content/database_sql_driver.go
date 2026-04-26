package content

func init() {
	Register(&Package{
		Name:       "database/sql/driver",
		ImportPath: "database/sql/driver",
		Category:   "Database",
		Summary:    "Interfaces implemented by database drivers. You only touch this if you're writing a driver — or implementing Valuer/Scanner on custom types.",
		Sections: []Section{
			{
				Title: "Custom value types",
				Description: "Make your own type round-trip through database/sql by implementing Valuer and Scanner.",
				Examples: []Example{
					{Title: "Valuer and Scanner", Code: `type JSON map[string]any

func (j JSON) Value() (driver.Value, error) {
    return json.Marshal(j)
}

func (j *JSON) Scan(src any) error {
    switch v := src.(type) {
    case []byte:
        return json.Unmarshal(v, j)
    case string:
        return json.Unmarshal([]byte(v), j)
    case nil:
        *j = nil; return nil
    }
    return fmt.Errorf("cannot scan %T into JSON", src)
}`},
				},
			},
			{
				Title: "Writing a driver (rare)",
				Examples: []Example{
					{Title: "Register a driver", Code: `func init() {
    sql.Register("mydriver", &MyDriver{})
}
// Implement driver.Driver, driver.Connector, driver.Conn, etc.`},
				},
			},
		},
	})
}
