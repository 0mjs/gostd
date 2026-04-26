package content

func init() {
	Register(&Package{
		Name:       "net/rpc",
		ImportPath: "net/rpc",
		Category:   "Networking",
		Summary:    "A Go-specific RPC framework. Frozen — not developed further. For new systems prefer gRPC or Twirp.",
		Sections: []Section{
			{
				Title: "What it looks like",
				Examples: []Example{
					{
						Title: "Registered method signature",
						Code: `type Args struct{ A, B int }

type Calc struct{}
func (c *Calc) Add(args *Args, reply *int) error {
    *reply = args.A + args.B
    return nil
}

rpc.Register(&Calc{})
rpc.HandleHTTP()
http.ListenAndServe(":8080", nil)`,
					},
					{
						Title: "Client",
						Code: `client, _ := rpc.DialHTTP("tcp", "localhost:8080")
var result int
client.Call("Calc.Add", Args{3, 4}, &result)`,
					},
				},
			},
		},
	})
}
