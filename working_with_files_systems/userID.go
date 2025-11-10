package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println(os.Getuid())
		return
	}

	u := arguments[1]
	us, err := user.Lookup(u)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(us.Uid)
}
