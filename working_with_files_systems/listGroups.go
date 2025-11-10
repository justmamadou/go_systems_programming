package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	arguments := os.Args
	var u *user.User
	var err error
	if len(arguments) == 1 {
		u, err = user.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	} else {
		username := arguments[1]
		u, err = user.Lookup(username)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}

	gids, _ := u.GroupIds()

	for _, gid := range gids {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			fmt.Print(err)
			os.Exit(-1)
		}
		fmt.Printf("(%s)%s\t", group.Gid, group.Name)
	}
}
