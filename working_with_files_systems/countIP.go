package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	minusCOL := flag.Int("col", 1, "Column number")
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Printf("Usage: %s [filename1] [filename2]...[filename_n.].", filepath.Base(flags[0]))
		os.Exit(-1)
	}

	column := *minusCOL
	if column < 0 {
		fmt.Printf("Invalied Column number!")
		os.Exit(1)
	}

	myIPs := make(map[string]int)

	for _, filename := range flags {
		f, err := os.Open(filename)
		fmt.Println("\t\t", filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				continue
			}
			data := strings.Fields(line)
			ip := data[column-1]
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}

			_, ok := myIPs[ip]
			if ok {
				myIPs[ip] = myIPs[ip] + 1
			} else {
				myIPs[ip] = 1
			}
		}
	}

	for key, _ := range myIPs {
		fmt.Printf("%s %d\n", key, myIPs[key])
	}

}
