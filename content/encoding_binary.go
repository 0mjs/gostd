package content

func init() {
	Register(&Package{
		Name:       "encoding/binary",
		ImportPath: "encoding/binary",
		Category:   "Encoding",
		Summary:    "Read and write fixed-size binary numbers. Big-endian, little-endian, and variable-length (varint).",
		Sections: []Section{
			{
				Title: "Byte orders",
				Description: "binary.BigEndian and binary.LittleEndian are values implementing ByteOrder. Use BigEndian for network protocols (aka 'network byte order'), LittleEndian for most on-disk formats.",
			},
			{
				Title: "Fixed-width integers",
				Examples: []Example{
					{
						Title: "PutUint32 / Uint32 — []byte round-trip",
						Code: `buf := make([]byte, 4)
binary.BigEndian.PutUint32(buf, 0xDEADBEEF)
fmt.Printf("% x\n", buf)               // de ad be ef
fmt.Printf("%x\n", binary.BigEndian.Uint32(buf))`,
					},
					{
						Title: "Read / Write — streaming over a Reader/Writer",
						Notes: "Encodes any fixed-size type: numbers, arrays, and structs of fixed-size fields (no slices, no strings).",
						Code: `type Header struct {
    Magic uint32
    Size  uint64
}

var h Header
err := binary.Read(r, binary.BigEndian, &h)`,
					},
				},
			},
			{
				Title: "Variable-length (varint)",
				Description: "Compact encoding where small numbers take fewer bytes. Used by protobuf and many wire formats.",
				Examples: []Example{
					{
						Title: "Uvarint and Varint",
						Code: `buf := make([]byte, binary.MaxVarintLen64)
n := binary.PutUvarint(buf, 12345)
fmt.Printf("% x\n", buf[:n])

v, _ := binary.Uvarint(buf)
fmt.Println(v)`,
					},
				},
			},
			{
				Title: "Append variants (1.19+)",
				Examples: []Example{
					{
						Title: "AppendUvarint, AppendUint32 — allocation-friendly",
						Code: `b := []byte{}
b = binary.BigEndian.AppendUint32(b, 42)
b = binary.AppendUvarint(b, 12345)`,
					},
				},
			},
		},
	})
}
