package content

func init() {
	Register(&Package{
		Name:       "os/user",
		ImportPath: "os/user",
		Category:   "CLI & Runtime",
		Summary:    "Look up the current user, users by name/uid, and their groups.",
		Sections: []Section{
			{
				Title: "Current user and home dir",
				Examples: []Example{
					{
						Title: "user.Current",
						Code: `u, _ := user.Current()
fmt.Println(u.Username, u.Uid, u.HomeDir)`,
					},
					{
						Title: "Home dir — also available via os.UserHomeDir",
						Notes: "os.UserHomeDir is often what you really want — it consults $HOME and falls back correctly.",
						Code: `home, _ := os.UserHomeDir()
fmt.Println(home)`,
					},
				},
			},
			{
				Title: "Lookup",
				Examples: []Example{
					{
						Title: "Lookup by name or ID",
						Code: `u, _ := user.Lookup("ada")
u, _  = user.LookupId("1000")`,
					},
					{
						Title: "GroupIds — groups a user belongs to",
						Code: `gids, _ := u.GroupIds()
for _, gid := range gids {
    g, _ := user.LookupGroupId(gid)
    fmt.Println(g.Name)
}`,
					},
				},
			},
		},
	})
}
