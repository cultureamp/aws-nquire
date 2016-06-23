package logging

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	debug *log.Logger
	info  *log.Logger
)

func init() {
	fmt.Println("init being called")

	debug = log.New(os.Stdout,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Debug(m string) {
	debugFlag := os.Getenv("NQUIRE_DEBUG")
	if strings.EqualFold(debugFlag, "true") {
		debug.Println(m)
	}
}

func Info(m string) {
	info.Println(m)
}
