package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

var gUID int32 = 0
var INCLUDE bool = true

func userOfFIle(filename string) int32 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		return 100000
	}

	UID := fileInfo.Sys().(*syscall.Stat_t).Uid
	return int32(UID)
}

func walkFunction(path string, info os.FileInfo, err error) error {
	_, err = os.Lstat(path)
	if err != nil {
		return err
	}
	if userOfFIle(path) == gUID && INCLUDE {
		fmt.Println(path)
	} else if userOfFIle(path) != gUID && !(INCLUDE) {
		fmt.Println(path)
	}
	return err
}
func main() {
	minusNO := flag.Bool("no", true, "Include")
	minusPATH := flag.String("path", ".", "Path to Search")
	flag.Parse()
	flags := flag.Args()
	INCLUDE = *minusNO
	Path := *minusPATH

	if len(flags) == 0 {
		gUID = int32(os.Getuid())
	} else {
		u, err := user.Lookup(flags[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		temp, err := strconv.ParseInt(u.Uid, 10, 32)
		gUID = int32(temp)
	}

	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
	}

}
