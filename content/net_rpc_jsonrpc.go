package content

func init() {
	Register(&Package{
		Name:       "net/rpc/jsonrpc",
		ImportPath: "net/rpc/jsonrpc",
		Category:   "Networking",
		Summary:    "JSON-RPC 1.0 codec for net/rpc. Both net/rpc and this package are frozen — use gRPC or plain HTTP+JSON for new services.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Server", Code: `rpc.Register(new(Arith))
l, _ := net.Listen("tcp", ":1234")
for {
    conn, _ := l.Accept()
    go jsonrpc.ServeConn(conn)
}`},
					{Title: "Client", Code: `conn, _ := net.Dial("tcp", "localhost:1234")
c := jsonrpc.NewClient(conn)
var reply int
c.Call("Arith.Add", Args{3, 4}, &reply)`},
				},
			},
		},
	})
}
