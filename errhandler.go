package errorhandler

import (
	"fmt"
	"log"
	"os"
)

func LoggErrorTo(pathfile string, errOuter error) {
	f, err := os.OpenFile(pathfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "--- error: ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags | log.Lshortfile)

	iLog.Println(errOuter)
}

func LoggError(errOuter error) {
	f, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "--- error: ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags | log.Lshortfile)

	iLog.Println(errOuter)
}