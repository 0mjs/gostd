package content

// StartHere is an opinionated reading order for someone new to Go's stdlib.
// The premise: a beginner who learns these in this order can build almost any
// real program. Anything else specializes from this base.
type StartHereEntry struct {
	Pkg      string // import path, must exist in registry
	OneLine  string // why you'd reach for this *now*, not later
	NextStep string // a tiny prompt that gives the reader a concrete first task
}

var startHere = []StartHereEntry{
	{Pkg: "fmt", OneLine: "Print things, format strings, build error messages.", NextStep: "Print a struct three ways with %v, %+v, %#v."},
	{Pkg: "errors", OneLine: "Idiomatic error handling — sentinels, wrapping, Is/As.", NextStep: "Wrap an error with fmt.Errorf %w and check it with errors.Is."},
	{Pkg: "strings", OneLine: "Read, search, and reshape UTF-8 text.", NextStep: "Use strings.Builder to assemble a CSV row."},
	{Pkg: "strconv", OneLine: "Parse \"42\" into an int and back.", NextStep: "Read a number from the command line with Atoi."},
	{Pkg: "os", OneLine: "Files, env vars, args, exit codes.", NextStep: "Read a file with os.ReadFile and print its size."},
	{Pkg: "io", OneLine: "Streams: Reader/Writer interfaces glue everything.", NextStep: "Pipe a file into stdout with io.Copy."},
	{Pkg: "bufio", OneLine: "Read a file or stdin one line at a time.", NextStep: "Count lines in a file with bufio.Scanner."},
	{Pkg: "encoding/json", OneLine: "Encode/decode Go values as JSON.", NextStep: "Round-trip a struct: Marshal, then Unmarshal back."},
	{Pkg: "time", OneLine: "Wall time, durations, sleep, format/parse.", NextStep: "Format time.Now() as RFC3339; parse it back."},
	{Pkg: "context", OneLine: "Cancellation and deadlines through your call graph.", NextStep: "Run a slow function with a 1-second WithTimeout."},
	{Pkg: "net/http", OneLine: "HTTP client + server in one package.", NextStep: "Build a server that responds \"hello, {name}\" using path values."},
	{Pkg: "testing", OneLine: "Write tests, benchmarks, and table-driven cases.", NextStep: "Write a TestX with t.Run subtests over a table."},
	{Pkg: "slices", OneLine: "Generic slice helpers — sort, search, transform.", NextStep: "Sort a []User by Age with slices.SortFunc and cmp.Compare."},
	{Pkg: "log/slog", OneLine: "Structured logging in the stdlib.", NextStep: "Log a request with slog.Info(\"req\", \"method\", m, \"path\", p)."},
}

// StartHere returns the curriculum entries that have a registered package.
// Order is preserved.
func StartHere() []StartHereEntry {
	out := make([]StartHereEntry, 0, len(startHere))
	for _, e := range startHere {
		if _, ok := registry[e.Pkg]; ok {
			out = append(out, e)
		}
	}
	return out
}
