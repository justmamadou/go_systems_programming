package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	minusF := flag.Bool("f", false, "Regular file")
	minusD := flag.Bool("d", false, "Directory")
	minusL := flag.Bool("l", false, "Symbolik link")
	minusS := flag.Bool("s", false, "Socket")
	flag.Parse()
	flags := flag.Args()
	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	printAll := false
	if *minusD && *minusF && *minusL && *minusS {
		printAll = true
	}
	if !(*minusD || *minusF || *minusL || *minusS) {
		printAll = true
	}

	Path := flags[0]

	walkFunction := func(path string, info os.FileInfo, err error) error {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if printAll {
			fmt.Println(path)
			return nil
		}
		mode := fileInfo.Mode()
		if mode.IsRegular() && *minusF {
			fmt.Println(path)
			return nil
		}
		if mode.IsDir() && *minusD {
			fmt.Println(path)
			return nil
		}

		fileInfo, err = os.Lstat(path)
		if err != nil {
			return err
		}
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if *minusL {
				fmt.Println(path)
				return nil
			}
		}
		if fileInfo.Mode()&os.ModeSocket != 0{
			if *minusS{
				fmt.Println(path)
				return nil
			}
		}
		return nil
	}
	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
