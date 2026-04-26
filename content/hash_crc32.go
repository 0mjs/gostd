package content

func init() {
	Register(&Package{
		Name:       "hash/crc32",
		ImportPath: "hash/crc32",
		Category:   "Hashing",
		Summary:    "CRC-32 checksums. Fast, 32-bit, great for integrity checks — not a secure hash.",
		Sections: []Section{
			{
				Title: "Compute a CRC-32",
				Examples: []Example{
					{
						Title: "ChecksumIEEE — the common polynomial",
						Code: `sum := crc32.ChecksumIEEE([]byte("hello"))
fmt.Printf("%08x\n", sum)`,
					},
					{
						Title: "Streaming via Hash32",
						Code: `h := crc32.NewIEEE()
io.Copy(h, f)
fmt.Printf("%08x\n", h.Sum32())`,
					},
					{
						Title: "Other polynomials",
						Notes: "Castagnoli (used by iSCSI and ext4) is faster on modern CPUs thanks to hardware support.",
						Code: `t := crc32.MakeTable(crc32.Castagnoli)
sum := crc32.Checksum(data, t)`,
					},
				},
			},
		},
	})
}
