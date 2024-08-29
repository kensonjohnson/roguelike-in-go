package logger

import (
	"log"
	"os"
)

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugOn     bool
)

func init() {
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ltime|log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)
	WarnLogger = log.New(os.Stdout, "WARN: ", log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ltime|log.Lshortfile)
}

func SetDebug(flag bool) {
	DebugOn = flag
}
