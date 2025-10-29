package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func excludeNames(name string, exclude string) bool {
	if exclude == "" {
		return false
	}
	if filepath.Base(name) == exclude {
		return true
	}
	return false
}

func excludeExtensions(name string, extension string) bool {
	if extension == "" {
		return false
	}
	basename := filepath.Base(name)
	s := strings.Split(basename, ".")
	length := len(s)
	basenameExtension := s[length-1]
	if basenameExtension == extension {
		return true
	}
	return false
}

func main() {

	minusS := flag.Bool("s", false, "This is for printing socket files")
	minusP := flag.Bool("p", false, "TThis is for printing pipes")
	minusSL := flag.Bool("sl", false, "This is for printing symbolic links")
	minusD := flag.Bool("d", false, "This is for printing directories")
	minusF := flag.Bool("f", false, "This is for printing files")
	minusX := flag.String("x", "", "Files")
	minusEXT := flag.String("ext", "", "Extensions")

	flag.Parse()
	flags := flag.Args()

	printAll := false
	if *minusS && *minusP && *minusSL && *minusD && *minusF {
		printAll = true
	}
	if !(*minusS || *minusP || *minusSL || *minusD || *minusF) {
		printAll = true
	}

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
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
		if excludeNames(path, *minusX) {
			return nil
		}
		if excludeExtensions(path, *minusEXT) {
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
		fileInfo, _ = os.Lstat(path)
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if *minusSL {
				fmt.Println(path)
				return nil
			}
		}
		if fileInfo.Mode()&os.ModeNamedPipe != 0 {
			if *minusP {
				fmt.Println(path)
				return nil
			}
		}
		if fileInfo.Mode()&os.ModeSocket != 0 {
			if *minusS {
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
