package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a DNS Name!")
		os.Exit(100)
	}

	domaine := arguments[1]
	ips, err := net.LookupHost(domaine)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	for _, ip := range ips {
		fmt.Printf("%s\n", ip)
	}
}
