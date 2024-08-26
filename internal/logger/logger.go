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
	DebugOn     *bool
)

func init() {
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ltime|log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)
	WarnLogger = log.New(os.Stdout, "WARN: ", log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ltime|log.Lshortfile)
}

func SetDebug(flag *bool) {
	DebugOn = flag
}

func Debug(v ...any) {
	if DebugOn != nil && *DebugOn {
		DebugLogger.Println(v...)
	}
}

func Debugf(format string, v ...any) {
	if DebugOn != nil && *DebugOn {
		DebugLogger.Printf(format, v...)
	}
}

func Info(v ...any) {
	InfoLogger.Println(v...)
}

func Infof(format string, v ...any) {
	InfoLogger.Printf(format, v...)
}

func Warn(v ...any) {
	WarnLogger.Println(v...)
}

func Warnf(format string, v ...any) {
	WarnLogger.Printf(format, v...)
}

func Fatal(v ...any) {
	ErrorLogger.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	ErrorLogger.Fatalf(format, v...)
}
