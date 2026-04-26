package content

func init() {
	Register(&Package{
		Name:       "net",
		ImportPath: "net",
		Category:   "Networking",
		Summary:    "The low-level networking package: TCP, UDP, UNIX sockets, DNS resolution. net/http and net/url sit on top of this.",
		Sections: []Section{
			{
				Title: "TCP server and client",
				Examples: []Example{
					{
						Title: "TCP echo server",
						Code: `ln, err := net.Listen("tcp", ":8080")
if err != nil { log.Fatal(err) }
defer ln.Close()

for {
    conn, err := ln.Accept()
    if err != nil { log.Print(err); continue }
    go func(c net.Conn) {
        defer c.Close()
        io.Copy(c, c)
    }(conn)
}`,
					},
					{
						Title: "Dial — make an outgoing connection",
						Code: `conn, err := net.Dial("tcp", "example.com:80")
if err != nil { log.Fatal(err) }
defer conn.Close()

fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
io.Copy(os.Stdout, conn)`,
					},
					{
						Title: "Dialer — timeouts and keep-alive",
						Notes: "net.Dial gives you a default dialer. Construct your own for timeouts, local address binding, or context cancellation.",
						Code: `d := &net.Dialer{Timeout: 5 * time.Second, KeepAlive: 30 * time.Second}
conn, err := d.DialContext(ctx, "tcp", "example.com:443")`,
					},
				},
			},
			{
				Title: "UDP",
				Examples: []Example{
					{
						Title: "Listen and read packets",
						Code: `pc, _ := net.ListenPacket("udp", ":9000")
defer pc.Close()

buf := make([]byte, 1500)
for {
    n, addr, _ := pc.ReadFrom(buf)
    fmt.Printf("%d bytes from %s: %s\n", n, addr, buf[:n])
}`,
					},
				},
			},
			{
				Title: "DNS",
				Examples: []Example{
					{
						Title: "LookupHost, LookupIP, LookupMX",
						Code: `ips, _ := net.LookupIP("example.com")
for _, ip := range ips {
    fmt.Println(ip)
}`,
					},
				},
			},
			{
				Title: "Addresses and ports",
				Examples: []Example{
					{
						Title: "SplitHostPort and JoinHostPort",
						Code: `host, port, _ := net.SplitHostPort("example.com:8080")
addr := net.JoinHostPort("::1", "8080")   // handles IPv6 brackets correctly`,
					},
				},
			},
		},
	})
}
