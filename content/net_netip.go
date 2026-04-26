package content

func init() {
	Register(&Package{
		Name:       "net/netip",
		ImportPath: "net/netip",
		Category:   "Networking",
		Summary:    "A modern, comparable, immutable IP address type. Zero allocations, usable as a map key. Prefer this over net.IP in new code.",
		Sections: []Section{
			{
				Title: "Why netip over net.IP?",
				Description: "net.IP is a []byte — not comparable, allocates, has representation ambiguities (v4 in v6). netip.Addr is a struct: comparable with ==, usable as a map key, fixed size, no allocations.",
			},
			{
				Title: "Addr — a single IP",
				Examples: []Example{
					{
						Title: "Parse and inspect",
						Code: `ip, err := netip.ParseAddr("192.0.2.1")
fmt.Println(ip.Is4(), ip.Is6())
fmt.Println(ip.IsPrivate(), ip.IsLoopback())

v6 := netip.MustParseAddr("2001:db8::1")
fmt.Println(v6.Is4(), v6.Is6())`,
					},
					{
						Title: "Use as a map key",
						Code: `seen := map[netip.Addr]int{}
seen[ip]++   // can't do this with net.IP`,
					},
				},
			},
			{
				Title: "AddrPort — Addr + port number",
				Examples: []Example{
					{
						Title: "Parse a host:port",
						Code: `ap, _ := netip.ParseAddrPort("192.0.2.1:443")
fmt.Println(ap.Addr(), ap.Port())`,
					},
				},
			},
			{
				Title: "Prefix — CIDR-style subnet",
				Examples: []Example{
					{
						Title: "Containment checks",
						Code: `p := netip.MustParsePrefix("10.0.0.0/8")
fmt.Println(p.Contains(netip.MustParseAddr("10.1.2.3")))  // true
fmt.Println(p.Contains(netip.MustParseAddr("192.0.2.1"))) // false`,
					},
				},
			},
		},
	})
}
