package content

func init() {
	Register(&Package{
		Name:       "structs",
		ImportPath: "structs",
		Category:   "Misc",
		Summary:    "Marker types you embed in your own structs to opt into special behavior. Currently just HostLayout (Go 1.23+).",
		Sections: []Section{
			{
				Title: "HostLayout",
				Description: "Tell the compiler a struct matches the host platform's C layout — needed for syscall/cgo interop, not everyday code.",
				Examples: []Example{
					{Title: "Opt-in host layout", Code: `type timespec struct {
    _ structs.HostLayout
    Sec  int64
    Nsec int64
}`},
				},
			},
		},
	})
}
