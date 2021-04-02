package migraty

import (
	"log"
	"os"
)

var info *log.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var warning *log.Logger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
var error *log.Logger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)

func logInfo(v ...interface{}) {
	info.Println(v...)
}

func logWarning(v ...interface{}) {
	warning.Println(v...)
}

func logError(v ...interface{}) {
	error.Println(v...)
}
