package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Crit("Crit: Logging in Go!")

	sysLog, err = syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "some program!")
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Emerg("Emerg: Logging in Go!")

	fmt.Fprintf(sysLog, "log.Print: Logging in Go!")

}
