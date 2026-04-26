package content

func init() {
	Register(&Package{
		Name:       "encoding/json/v2",
		ImportPath: "encoding/json/v2",
		Category:   "Encoding",
		Summary:    "Rewrite of encoding/json (experimental in Go 1.25, GOEXPERIMENT=jsonv2). Faster, stricter, configurable. Opt-in.",
		Sections: []Section{
			{
				Title: "Marshal / Unmarshal",
				Examples: []Example{
					{Title: "Basic usage", Code: `import jsonv2 "encoding/json/v2"

data, err := jsonv2.Marshal(user)
err = jsonv2.Unmarshal(data, &user)`},
				},
			},
			{
				Title: "Options",
				Description: "v2 uses functional options instead of flag-overloaded struct tags.",
				Examples: []Example{
					{Title: "Marshal with options", Code: `data, err := jsonv2.Marshal(v,
    jsonv2.Deterministic(true),
    jsonv2.FormatNilSliceAsNull(true),
)`},
				},
			},
			{
				Title: "When to use it",
				Examples: []Example{
					{Title: "Trade-offs", Code: `// Pros: clearer semantics, better performance, explicit options.
// Cons: experimental — APIs may change, and tooling (lint/codegen) may lag.
// Stick with encoding/json for production today; watch the proposal.`},
				},
			},
		},
	})
}
