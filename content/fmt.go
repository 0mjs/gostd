package content

func init() {
	Register(&Package{
		Name:       "fmt",
		ImportPath: "fmt",
		Category:   "Formatting & Strings",
		Summary:    "Formatted I/O: Print, Printf, Sprintf, Scan, Errorf. Go's go-to for building and parsing text.",
		Sections: []Section{
			{
				Title: "The three print families",
				Description: "Each family comes in three flavors: no suffix (no newline), ln (adds spaces between args and a newline), and f (format string). They also come with prefixes: nothing (stdout), S (return string), F (write to any io.Writer), and Errorf (returns an error).",
				Examples: []Example{
					{
						Title: "Println — spaces between args, newline at end",
						Code: `fmt.Println("hello", "world", 42)
fmt.Println() // blank line`,
						Output: `hello world 42

`,
					},
					{
						Title: "Print — no spaces between args unless they're both strings",
						Notes: "Print inserts a space between two non-string operands. It does NOT add a newline.",
						Code: `fmt.Print("a", "b")      // no space: ab
fmt.Print(1, 2)          // space:   1 2
fmt.Print("x", 7, "y\n") // mixed:  x7 y`,
						Output: `ab1 2x7 y`,
					},
					{
						Title: "Printf — explicit format string",
						Code: `fmt.Printf("%s is %d years old\n", "Ada", 36)`,
						Output: `Ada is 36 years old
`,
					},
					{
						Title: "Sprintf — format into a string",
						Notes: "Use Sprintf when you want the result as a string (e.g., for logging, building messages).",
						Code: `msg := fmt.Sprintf("user=%s id=%d", "ada", 7)
fmt.Println(msg)`,
						Output: `user=ada id=7
`,
					},
					{
						Title: "Fprintf — write to any io.Writer",
						Notes: "F-variants take an io.Writer first. Use with os.Stderr, files, bytes.Buffer, http.ResponseWriter, etc.",
						Code: `fmt.Fprintf(os.Stderr, "warning: %s\n", "disk almost full")

var buf bytes.Buffer
fmt.Fprintf(&buf, "line %d\n", 1)
fmt.Print(buf.String())`,
						Output: `line 1
`,
					},
					{
						Title: "Errorf — build an error",
						Notes: "Use %w to wrap another error so errors.Is and errors.As work.",
						Code: `err := fmt.Errorf("open %q: %w", "/tmp/x", os.ErrNotExist)
fmt.Println(err)
fmt.Println(errors.Is(err, os.ErrNotExist))`,
						Output: `open "/tmp/x": file does not exist
true
`,
					},
				},
			},
			{
				Title: "The common verbs",
				Description: "Verbs begin with % and describe how to format an argument.",
				Examples: []Example{
					{
						Title: "%v, %+v, %#v — the general verbs",
						Notes: "%v is the default. %+v shows struct field names. %#v prints Go syntax (great for debugging).",
						Code: `type Point struct{ X, Y int }
p := Point{1, 2}
fmt.Printf("%v\n", p)
fmt.Printf("%+v\n", p)
fmt.Printf("%#v\n", p)`,
						Output: `{1 2}
{X:1 Y:2}
main.Point{X:1, Y:2}
`,
					},
					{
						Title: "%T — the type of a value",
						Code: `var x any = 3.14
fmt.Printf("%T\n", x)`,
						Output: `float64
`,
					},
					{
						Title: "%d, %b, %o, %x, %X — integers in different bases",
						Code: `fmt.Printf("%d %b %o %x %X\n", 255, 255, 255, 255, 255)`,
						Output: `255 11111111 377 ff FF
`,
					},
					{
						Title: "%s, %q, %x — strings",
						Notes: "%s is raw. %q adds Go-quoted escapes. %x encodes as hex.",
						Code: `s := "hi\tthere"
fmt.Printf("%s\n", s)
fmt.Printf("%q\n", s)
fmt.Printf("%x\n", s)`,
						Output: `hi	there
"hi\tthere"
68690974686572650
`,
					},
					{
						Title: "%f, %e, %g — floats",
						Notes: "%f = decimal, %e = scientific, %g = whichever is shorter.",
						Code: `fmt.Printf("%f\n", 1234.5678)
fmt.Printf("%e\n", 1234.5678)
fmt.Printf("%g\n", 1234.5678)`,
						Output: `1234.567800
1.234568e+03
1234.5678
`,
					},
					{
						Title: "%t, %c, %U, %p — booleans, runes, Unicode, pointers",
						Code: `fmt.Printf("%t\n", true)
fmt.Printf("%c %U\n", 'G', 'G')
x := 42
fmt.Printf("%p\n", &x)`,
						Output: `true
G U+0047
0xc000012345
`,
					},
				},
			},
			{
				Title: "Width, precision, padding, flags",
				Description: "Between the % and the verb you can put: flags (+, -, #, 0, space), width, and .precision.",
				Examples: []Example{
					{
						Title: "Width — minimum column width",
						Notes: "Positive width right-aligns, negative left-aligns.",
						Code: `fmt.Printf("[%5d]\n", 42)   // right-aligned
fmt.Printf("[%-5d]\n", 42)  // left-aligned
fmt.Printf("[%05d]\n", 42)  // zero-padded`,
						Output: `[   42]
[42   ]
[00042]
`,
					},
					{
						Title: "Precision on floats and strings",
						Code: `fmt.Printf("%.2f\n", 3.14159)   // 2 digits after point
fmt.Printf("%8.2f\n", 3.14159)  // width 8, precision 2
fmt.Printf("%.3s\n", "hello")   // truncate string to 3`,
						Output: `3.14
    3.14
hel
`,
					},
					{
						Title: "%+d and % d — signed with leading sign or space",
						Code: `fmt.Printf("%+d %+d\n", 7, -7)
fmt.Printf("% d % d\n", 7, -7)`,
						Output: `+7 -7
 7 -7
`,
					},
				},
			},
			{
				Title: "Scan / Sscan / Fscan",
				Description: "Parsing back. Scan reads from stdin, Sscan from a string, Fscan from a Reader.",
				Examples: []Example{
					{
						Title: "Sscanf — parse a formatted string",
						Code: `var name string
var age int
n, err := fmt.Sscanf("Ada 36", "%s %d", &name, &age)
fmt.Println(n, err, name, age)`,
						Output: `2 <nil> Ada 36
`,
					},
					{
						Title: "Sscan — whitespace-separated",
						Code: `var a, b, c int
fmt.Sscan("1 2 3", &a, &b, &c)
fmt.Println(a + b + c)`,
						Output: `6
`,
					},
				},
			},
			{
				Title: "Stringer interface",
				Description: "Any type with a String() string method formats itself for %s and %v.",
				Examples: []Example{
					{
						Title: "Custom formatting via Stringer",
						Code: `type Celsius float64
func (c Celsius) String() string {
    return fmt.Sprintf("%.1f°C", float64(c))
}
fmt.Println(Celsius(22.5))`,
						Output: `22.5°C
`,
					},
					{
						Title: "Error vs Stringer — error wins for %v",
						Notes: "If a type implements both error and Stringer, the error method is used.",
						Code: `var err error = fmt.Errorf("boom")
fmt.Printf("%v\n", err)  // "boom"
fmt.Printf("%s\n", err)  // "boom"`,
						Output: `boom
boom
`,
					},
				},
			},
		},
	})
}
