package content

func init() {
	Register(&Package{
		Name:       "runtime/cgo",
		ImportPath: "runtime/cgo",
		Category:   "Runtime & Debug",
		Summary:    "Support package for cgo. Most programs import this transparently via `import \"C\"`. Its direct API is a handle system for passing Go values to C.",
		Sections: []Section{
			{
				Title: "Handle — pass a Go value through C",
				Description: "Safely round-trip a Go value across a C call that stores an opaque pointer.",
				Examples: []Example{
					{Title: "NewHandle / Value / Delete", Code: `// Package foo imports "C" and calls some C API that takes void*.
obj := &MyStruct{Name: "Ada"}
h := cgo.NewHandle(obj)

// Pass uintptr(h) to C as void*. Later, in a Go callback:
v := h.Value().(*MyStruct)

// Free when done — handles are never collected automatically:
h.Delete()`},
				},
			},
		},
	})
}
