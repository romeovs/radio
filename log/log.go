package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// LogFile is the file to which the log messages will be written.
var LogFile = "./radio.log"

type level string

const (
	info  level = "info"
	debug level = "debug"
	err   level = "error"
)

// Entry is a log entry
type Entry struct {
	Level     level         `json:"lvl"`
	Message   string        `json:"msg"`
	Timestamp time.Time     `json:"ts"`
	Args      []interface{} `json:"args"`
}

// Info prints an info log message.
func Info(msg string, args ...interface{}) {
	log(info, msg, args...)
}

// Debug prints a debug log message.
func Debug(msg string, args ...interface{}) {
	log(debug, msg, args...)
}

// Error prints a log error.
func Error(msg string, args ...interface{}) {
	log(err, msg, args...)
}

func log(level level, msg string, args ...interface{}) {
	entry := &Entry{
		Level:     level,
		Message:   fmt.Sprintf(msg, args...),
		Timestamp: time.Now(),
		Args:      args,
	}

	console(entry)
	logfile(entry)
}

// console prints the log message to the console.
func console(entry *Entry) {
	Fprint(os.Stdout, entry)
}

// Fprint pretty prints the log entry to the writer.
func Fprint(w io.Writer, entry *Entry) {
	fmt.Fprintf(w, "radio [%s] ", entry.Timestamp.Format("2006-01-02 15:04:05.000"))
	fmt.Fprintf(w, entry.Message)
	fmt.Fprintf(w, "\n")
}

// logfile prints the log message to the log file.
func logfile(entry *Entry) {
	f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(entry)
	if err != nil {
		panic(err)
	}
}

func init() {
	f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Truncate(0)
	f.Seek(0, 0)
}
