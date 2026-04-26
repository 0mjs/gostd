package content

func init() {
	Register(&Package{
		Name:       "reflect",
		ImportPath: "reflect",
		Category:   "Reflection & Unsafe",
		Summary:    "Inspect and manipulate arbitrary values at runtime. Slow and complex — prefer interfaces and generics when you can.",
		Sections: []Section{
			{
				Title: "Type and Value",
				Examples: []Example{
					{Title: "TypeOf / ValueOf", Code: `t := reflect.TypeOf(42)      // int
v := reflect.ValueOf("hi")    // string value
fmt.Println(t.Kind(), v.Kind()) // int string`},
					{Title: "Kind vs Type", Code: `type MyInt int
reflect.TypeOf(MyInt(1)).Kind() // reflect.Int
reflect.TypeOf(MyInt(1)).Name() // "MyInt"`},
				},
			},
			{
				Title: "Inspect structs",
				Examples: []Example{
					{Title: "Walk fields + tags", Code: "type User struct {\n    Name string `json:\"name\"`\n    Age  int    `json:\"age\"`\n}\nt := reflect.TypeOf(User{})\nfor i := 0; i < t.NumField(); i++ {\n    f := t.Field(i)\n    fmt.Println(f.Name, f.Type, f.Tag.Get(\"json\"))\n}"},
				},
			},
			{
				Title: "Mutate with reflection",
				Examples: []Example{
					{Title: "Settable values need pointers", Code: `x := 1
v := reflect.ValueOf(&x).Elem() // Elem dereferences
v.SetInt(42)
// x == 42`},
				},
			},
			{
				Title: "Dynamic calls",
				Examples: []Example{
					{Title: "Call a method by name", Code: `m := reflect.ValueOf(obj).MethodByName("Greet")
out := m.Call([]reflect.Value{reflect.ValueOf("world")})
fmt.Println(out[0].String())`},
				},
			},
			{
				Title: "DeepEqual",
				Examples: []Example{
					{Title: "Recursive equality", Code: `reflect.DeepEqual([]int{1,2}, []int{1,2}) // true
reflect.DeepEqual(map[string]int{"a":1}, map[string]int{"a":1}) // true`},
				},
			},
		},
	})
}
