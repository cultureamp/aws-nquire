package logging

import (
	"log"
	"os"
	"strings"
)

var (
	debug *log.Logger
	info  *log.Logger
	err   *log.Logger
)

func init() {
	debug = log.New(os.Stderr,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	info = log.New(os.Stderr,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	err = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Debug(m string) {
	debugFlag := os.Getenv("AWS_NQUIRE_DEBUG")
	if strings.EqualFold(debugFlag, "true") {
		debug.Println(m)
	}
}

func Info(m string) {
	info.Println(m)
}

func Error(m string) {
	err.Println(m)
}
