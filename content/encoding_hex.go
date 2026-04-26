package content

func init() {
	Register(&Package{
		Name:       "encoding/hex",
		ImportPath: "encoding/hex",
		Category:   "Encoding",
		Summary:    "Hex encoding: []byte ↔ ASCII hex string.",
		Sections: []Section{
			{
				Title: "One-shot encoding",
				Examples: []Example{
					{
						Title: "EncodeToString / DecodeString",
						Code: `s := hex.EncodeToString([]byte{0xDE, 0xAD, 0xBE, 0xEF})
fmt.Println(s)   // "deadbeef"

b, _ := hex.DecodeString(s)
fmt.Printf("% x\n", b)`,
					},
					{
						Title: "Dump — canonical hex+ASCII layout",
						Notes: "Same formatting as 'xxd' or 'hexdump -C'.",
						Code: `fmt.Print(hex.Dump([]byte("Hello, Go!")))`,
						Output: `00000000  48 65 6c 6c 6f 2c 20 47  6f 21                    |Hello, Go!|
`,
					},
				},
			},
			{
				Title: "Streaming",
				Examples: []Example{
					{
						Title: "NewEncoder / NewDecoder",
						Code: `w := hex.NewEncoder(os.Stdout)
w.Write([]byte("hi"))
fmt.Println()   // "6869"`,
					},
				},
			},
		},
	})
}
