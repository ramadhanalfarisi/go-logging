package go_logging

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

type logging struct {
	file  *os.File
	debug bool
}

var (
	date          = time.Now().Format("2006-01-02")
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	DebugLogger   *log.Logger
)

// To config logging function
func ConfigLogging(debug bool, folderpath string) logging {
	var f *os.File
	log_path := folderpath + date + ".txt"
	if _, err := os.Stat(log_path); os.IsNotExist(err) {
		var errcreate error
		f, errcreate = os.Create(log_path)
		if errcreate != nil {
			log.Fatal(errcreate)
		}
	} else {
		var erropen error
		f, erropen = os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if erropen != nil {
			log.Fatal(erropen)
		}
	}
	return logging{file: f, debug: debug}
}

// To log info
func (l logging) Info(msg string) {
	InfoLogger = log.New(l.file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger.Println(msg)
}

// To log warning
func (l logging) Warning(msg string) {
	WarningLogger = log.New(l.file, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger.Println(msg)
}

// To log debug
func (l logging) Debug(msg string) {
	DebugLogger = log.New(l.file, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	debug := l.debug
	if debug {
		DebugLogger.Println(msg)
	}
}

// To log error
func (l logging) Error(err error) {
	ErrorLogger = log.New(l.file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	debug := l.debug
	var logs string
	pc, fn, line, _ := runtime.Caller(1)
	if debug {
		logs = fmt.Sprintf("%s [%s:%s:%d]", err, runtime.FuncForPC(pc).Name(), fn, line)
	} else {
		logs = fmt.Sprintf("%s [%s:%d]", err, fn, line)
	}
	ErrorLogger.Println(logs)
}
