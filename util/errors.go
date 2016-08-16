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
		Fatal(err.Error() + "\n")
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
		Stderr(err.Error() + "\n")
		return true
	}
	return false
}
