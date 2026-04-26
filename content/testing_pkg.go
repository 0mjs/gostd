package content

func init() {
	Register(&Package{
		Name:       "testing",
		ImportPath: "testing",
		Category:   "Testing",
		Summary:    "Unit tests, benchmarks, fuzz tests, examples. Run with `go test`. File suffix _test.go.",
		Sections: []Section{
			{
				Title: "Unit tests",
				Examples: []Example{
					{Title: "func TestXxx(t *testing.T)", Code: `func TestAdd(t *testing.T) {
    got := Add(2, 3)
    if got != 5 {
        t.Errorf("Add(2,3) = %d, want 5", got)
    }
}`},
					{Title: "Table-driven", Code: `func TestAbs(t *testing.T) {
    tests := []struct{
        name string
        in, want int
    }{
        {"positive", 3, 3},
        {"negative", -3, 3},
        {"zero", 0, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Abs(tt.in); got != tt.want {
                t.Errorf("Abs(%d) = %d, want %d", tt.in, got, tt.want)
            }
        })
    }
}`},
					{Title: "Fail vs Skip vs Fatal", Code: `t.Error("continue running after logging")
t.Fatal("abort this test now")
t.Skip("skipping on this platform")
t.Helper() // mark helper — errors report caller line`},
					{Title: "Cleanup / TempDir", Code: `dir := t.TempDir() // auto-removed
t.Cleanup(func() { close(ch) })`},
					{Title: "Parallel", Code: `func TestThing(t *testing.T) {
    t.Parallel()
    // ...
}`},
				},
			},
			{
				Title: "Benchmarks",
				Examples: []Example{
					{Title: "func BenchmarkXxx(b *testing.B)", Code: `func BenchmarkEncode(b *testing.B) {
    data := make([]byte, 1024)
    b.ResetTimer()
    for b.Loop() { // Go 1.24+; or: for i := 0; i < b.N; i++
        _ = encode(data)
    }
}`},
					{Title: "Report allocations", Code: `b.ReportAllocs()
b.ReportMetric(float64(ops)/b.Elapsed().Seconds(), "ops/s")`},
				},
			},
			{
				Title: "Fuzz tests",
				Examples: []Example{
					{Title: "func FuzzXxx(f *testing.F)", Code: `func FuzzReverse(f *testing.F) {
    f.Add("hello")
    f.Fuzz(func(t *testing.T, s string) {
        r := Reverse(Reverse(s))
        if r != s {
            t.Errorf("double reverse changed %q -> %q", s, r)
        }
    })
}`},
				},
			},
			{
				Title: "Examples (runnable docs)",
				Examples: []Example{
					{Title: "func ExampleXxx", Code: `func ExampleHello() {
    fmt.Println(Hello("world"))
    // Output: hello, world
}`},
				},
			},
			{
				Title: "Running tests",
				Examples: []Example{
					{Title: "CLI", Code: `go test ./...                  # all packages
go test -run TestAdd -v        # one test
go test -bench=. -benchmem     # benchmarks
go test -fuzz=FuzzReverse      # fuzz
go test -race                  # race detector
go test -cover -coverprofile=c.out`},
				},
			},
		},
	})
}
