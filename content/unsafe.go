package content

func init() {
	Register(&Package{
		Name:       "unsafe",
		ImportPath: "unsafe",
		Category:   "Reflection & Unsafe",
		Summary:    "Escape hatch: bypass type safety for pointer arithmetic, zero-copy conversions, and layout tricks. Breaks portability if misused.",
		Sections: []Section{
			{
				Title: "Sizes and alignment",
				Examples: []Example{
					{Title: "Sizeof / Alignof / Offsetof", Code: `type T struct { A int32; B int64 }
var t T
unsafe.Sizeof(t)        // total bytes
unsafe.Alignof(t.B)     // alignment of field
unsafe.Offsetof(t.B)    // byte offset of field in struct`},
				},
			},
			{
				Title: "String and slice conversions",
				Description: "Go 1.20+ helpers for zero-copy interop.",
				Examples: []Example{
					{Title: "String <-> []byte, zero copy", Code: `// string from byte slice (read-only):
s := unsafe.String(unsafe.SliceData(b), len(b))

// byte pointer from string:
p := unsafe.StringData(s)
b := unsafe.Slice(p, len(s))

// IMPORTANT: the string data must not be mutated.`},
					{Title: "Build a slice from a C-style pointer", Code: `s := unsafe.Slice(ptr, n) // []T of length n backed by ptr`},
				},
			},
			{
				Title: "unsafe.Pointer rules",
				Description: "Only these conversions are valid — anything else is undefined.",
				Examples: []Example{
					{Title: "Allowed conversions", Code: `// 1. *T1 <-> unsafe.Pointer <-> *T2  (same memory)
// 2. unsafe.Pointer <-> uintptr        (arithmetic, but uintptr is NOT a GC root)
// 3. Use in syscall.Syscall or similar where required.`},
				},
			},
		},
	})
}
