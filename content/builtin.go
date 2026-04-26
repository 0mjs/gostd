package content

func init() {
	Register(&Package{
		Name:       "builtin",
		ImportPath: "builtin",
		Category:   "Misc",
		Summary:    "Documentation-only package for the predeclared identifiers (len, cap, append, make, new, delete, panic, recover, any, error, ...). Always in scope.",
		Sections: []Section{
			{
				Title: "Slices and maps",
				Examples: []Example{
					{Title: "make / new / len / cap", Code: `s := make([]int, 3, 8) // len 3, cap 8
m := make(map[string]int, 16)
p := new(int)             // *int pointing to zero value
len(s); cap(s); len(m)`},
					{Title: "append / copy / clear / delete", Code: `s = append(s, 4, 5)
n := copy(dst, src)
clear(s)          // Go 1.21+: zero all elements
clear(m)          // empty map
delete(m, "key")`},
				},
			},
			{
				Title: "Channels and control flow",
				Examples: []Example{
					{Title: "close / panic / recover", Code: `close(ch)
panic("boom")
defer func() { if r := recover(); r != nil { /* handled */ } }()`},
				},
			},
			{
				Title: "Numeric helpers",
				Examples: []Example{
					{Title: "min / max (Go 1.21+)", Code: `min(3, 1, 2)     // 1
max(3.0, 1.0)    // 3.0`},
				},
			},
			{
				Title: "Predeclared types",
				Examples: []Example{
					{Title: "Aliases", Code: `any          // alias for interface{}
error        // interface{ Error() string }
byte         // alias for uint8
rune         // alias for int32
comparable   // type constraint`},
				},
			},
		},
	})
}
