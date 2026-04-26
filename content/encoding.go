package content

func init() {
	Register(&Package{
		Name:       "encoding",
		ImportPath: "encoding",
		Category:   "Encoding",
		Summary:    "The shared interfaces every encoding package honors: BinaryMarshaler, TextMarshaler, and their Unmarshaler pairs.",
		Sections: []Section{
			{
				Title: "The four interfaces",
				Description: "Implement these once and your type slots into many encoders.",
				Examples: []Example{
					{
						Title: "encoding.TextMarshaler / TextUnmarshaler",
						Notes: "JSON, XML, YAML, env parsers all check for TextMarshaler. One implementation covers them all.",
						Code: `type Color struct{ R, G, B uint8 }

func (c Color) MarshalText() ([]byte, error) {
    return []byte(fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)), nil
}
func (c *Color) UnmarshalText(text []byte) error {
    _, err := fmt.Sscanf(string(text), "#%02x%02x%02x", &c.R, &c.G, &c.B)
    return err
}

b, _ := json.Marshal(Color{255, 0, 128})  // "\"#ff0080\""`,
					},
					{
						Title: "encoding.BinaryMarshaler / BinaryUnmarshaler",
						Notes: "Used by encoding/gob and some caches. Less common — reach for it when you need compact on-disk or on-wire formats.",
					},
				},
			},
		},
	})
}
