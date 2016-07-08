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
	debug = log.New(os.Stdout,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	err = log.New(os.Stdout,
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
