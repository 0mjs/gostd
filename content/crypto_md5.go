package content

func init() {
	Register(&Package{
		Name:       "crypto/md5",
		ImportPath: "crypto/md5",
		Category:   "Crypto",
		Summary:    "MD5 hashing. BROKEN for security use. Only for non-security checksums like ETags or dedup.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{
						Title: "Sum",
						Code: `sum := md5.Sum([]byte("hello"))
fmt.Printf("%x\n", sum)`,
					},
					{
						Title: "When to still use MD5",
						Notes: "Legitimate uses: cache keys, ETags, file-level dedup where collisions are not adversarial. For literally anything security-related use SHA-256 or stronger.",
					},
				},
			},
		},
	})
}
