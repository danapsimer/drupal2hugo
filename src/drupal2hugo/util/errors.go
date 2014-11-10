package util

import (
	"log"
	"fmt"
	"os"
)

func Stderr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func Fatal(format string, args ...interface{}) {
	Stderr(format, args...)
	os.Exit(1)
}

func logMsgs(msg ...interface{}) {
	if len(msg) > 0 {
		fmt.Fprintln(os.Stderr, msg...)
	}
}

func CheckErrFatal(err error, msg ...interface{}) {
	if err != nil {
		logMsgs(msg...)
		Fatal(err.Error())
	}
}

func CheckErrPanic(err error, msg ...interface{}) {
	if err != nil {
		logMsgs(msg...)
		log.Panicln(err)
	}
}

func LogError(err error, msg ...interface{}) bool {
	if err != nil {
		logMsgs(msg...)
		Stderr(err.Error())
		return true
	}
	return false
}
