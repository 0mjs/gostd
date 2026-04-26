package content

func init() {
	Register(&Package{
		Name:       "testing/iotest",
		ImportPath: "testing/iotest",
		Category:   "Testing",
		Summary:    "Readers and writers that simulate slow, broken, or one-byte-at-a-time I/O. Useful for stress-testing parsers.",
		Sections: []Section{
			{
				Title: "Simulated readers",
				Examples: []Example{
					{Title: "OneByteReader", Code: `r := iotest.OneByteReader(strings.NewReader("hello"))
// forces your code through many small reads`},
					{Title: "HalfReader / DataErrReader / ErrReader", Code: `r1 := iotest.HalfReader(src)                // returns half at a time
r2 := iotest.DataErrReader(src)             // returns final data+err together
r3 := iotest.ErrReader(io.ErrUnexpectedEOF) // always errors`},
					{Title: "TimeoutReader", Code: `r := iotest.TimeoutReader(src)
// second read returns ErrTimeout — exercise retry logic`},
				},
			},
			{
				Title: "Validate a Reader",
				Examples: []Example{
					{Title: "TestReader", Code: `if err := iotest.TestReader(myReader, wantBytes); err != nil {
    t.Fatal(err)
}`},
				},
			},
		},
	})
}
