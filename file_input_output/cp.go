package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var BUFFERSIZE int64

func Copy(source, destination string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", source)
	}

	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = os.Stat(destination)
	if err == nil {
		return fmt.Errorf("File %s already exists.", destination)
	}

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := dst.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("usage: %s source destination BUFFERSIZE\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	source := os.Args[1]
	destination := os.Args[2]
	BUFFERSIZE, _ = strconv.ParseInt(os.Args[3], 10, 64)

	fmt.Printf("Copying %s to %s\n", source, destination)
	err := Copy(source, destination, BUFFERSIZE)
	if err != nil {
		fmt.Printf("File copying failed: %q\n", err)
	}
}
