package util

import (
	"io"
	"os"
)

type Anything interface{}

type Nothing struct{}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

func AppendIfNeeded(s string, suffix byte) string {
	if s[len(s)-1] == suffix {
		return s
	}
	return string(append([]byte(s), suffix))
}

func Chdir(s string) {
	err := os.Chdir(s)
	CheckErrFatal(err)
}

func ConstructSomeLogWriter(logName string, dfltWriter io.Writer) io.Writer {
	writer := dfltWriter
	if logName != "" && logName != "-" {
		var err error
		var flag = 0
		if FileExists(logName) {
			flag = os.O_WRONLY | os.O_APPEND
		} else {
			flag = os.O_WRONLY | os.O_CREATE
		}
		writer, err = os.OpenFile(logName, flag, os.FileMode(0644))
		CheckErrFatal(err)
	}
	return writer
}

//func SliceEqual(a, b []string) bool {
//	if len(a) != len(b) {
//		return false
//	}
//	for i, v := range a {
//		if v != b[i] {
//			return false
//		}
//	}
//	return true
//}
