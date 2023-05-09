package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "mGo.log"
var LOGFILE2 = "http.go"

func main() {
	var err error

	// LoggErrorTo(LOGFILE, err)
	LoggError(err)

}

func LoggErrorTo(pathfile string, err error) {
	f, err := os.OpenFile(pathfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "--- error: ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags | log.Lshortfile)

	iLog.Println(err)
}

func LoggError(err error) {
	f, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "--- error: ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags | log.Lshortfile)

	iLog.Println(err)
}
