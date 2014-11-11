//-----------------------------------------------------------------------------
// The MIT License
//
// Copyright (c) 2012 Rick Beton <rick@bigbeeconsultants.co.uk>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//-----------------------------------------------------------------------------

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
