package content

func init() {
	Register(&Package{
		Name:       "crypto/subtle",
		ImportPath: "crypto/subtle",
		Category:   "Crypto",
		Summary:    "Constant-time operations. Use these when a timing leak would compromise secrets.",
		Sections: []Section{
			{
				Title: "ConstantTimeCompare",
				Examples: []Example{
					{
						Title: "Compare two secrets",
						Notes: "Returns 1 if equal, 0 otherwise. Only safe for equal-length inputs — check len first.",
						Code: `func eqSecret(a, b []byte) bool {
    return subtle.ConstantTimeCompare(a, b) == 1
}`,
					},
					{
						Title: "Also available: ConstantTimeSelect, ConstantTimeByteEq, XORBytes",
					},
				},
			},
		},
	})
}
