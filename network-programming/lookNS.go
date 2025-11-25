package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a Domaine!")
		os.Exit(100)
	}

	domain := arguments[1]
	nameservers, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	for _, ns := range nameservers {
		fmt.Println(ns.Host)
	}

}
