package log

import "fmt"

type level string

const (
	info  level = "info"
	debug level = "debug"
	err   level = "error"
)

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
	fmt.Printf("radio: ")
	fmt.Printf(msg, args...)
	fmt.Printf("\n")
}
