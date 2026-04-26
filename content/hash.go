package content

func init() {
	Register(&Package{
		Name:       "hash",
		ImportPath: "hash",
		Category:   "Hashing",
		Summary:    "The Hash interface every hash algorithm in stdlib implements. A growable io.Writer that produces a fixed-size digest.",
		Sections: []Section{
			{
				Title: "The interface",
				Description: "Hash embeds io.Writer, adds Sum(b []byte) []byte that appends the current digest without resetting, Reset(), Size(), and BlockSize(). Hash32 and Hash64 add Sum32/Sum64 for short hashes.",
				Examples: []Example{
					{
						Title: "The pattern, same for every hash",
						Code: `h := sha256.New()            // or md5.New, fnv.New32a, crc32.NewIEEE, etc.
io.Copy(h, someReader)
digest := h.Sum(nil)         // fresh slice with the digest
fmt.Printf("%x\n", digest)`,
					},
				},
			},
		},
	})
}
