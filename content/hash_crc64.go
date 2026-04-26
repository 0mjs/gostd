package content

func init() {
	Register(&Package{
		Name:       "hash/crc64",
		ImportPath: "hash/crc64",
		Category:   "Hashing",
		Summary:    "CRC-64 checksums. Same shape as crc32 with a 64-bit digest.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "Checksum",
						Code: `t := crc64.MakeTable(crc64.ECMA)
sum := crc64.Checksum([]byte("hello"), t)
fmt.Printf("%016x\n", sum)`,
					},
				},
			},
		},
	})
}
