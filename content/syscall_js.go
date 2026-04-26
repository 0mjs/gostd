package content

func init() {
	Register(&Package{
		Name:       "syscall/js",
		ImportPath: "syscall/js",
		Category:   "Misc",
		Summary:    "Call JavaScript from Go compiled to WebAssembly (GOOS=js, GOARCH=wasm). Build bridges to the DOM and browser APIs.",
		Sections: []Section{
			{
				Title: "Call into JS",
				Examples: []Example{
					{Title: "Global, Get, Call", Code: `doc := js.Global().Get("document")
body := doc.Get("body")
body.Set("innerHTML", "<h1>Hello from Go/Wasm</h1>")

js.Global().Get("console").Call("log", "hello")`},
				},
			},
			{
				Title: "Expose Go to JS",
				Examples: []Example{
					{Title: "FuncOf", Code: `add := js.FuncOf(func(this js.Value, args []js.Value) any {
    return args[0].Int() + args[1].Int()
})
defer add.Release()
js.Global().Set("goAdd", add)
// JS can now call goAdd(2, 3) and get 5.`},
				},
			},
			{
				Title: "Keep the program alive",
				Examples: []Example{
					{Title: "Block main", Code: `select {} // forever — otherwise the Go/Wasm program exits and releases Funcs`},
				},
			},
		},
	})
}
