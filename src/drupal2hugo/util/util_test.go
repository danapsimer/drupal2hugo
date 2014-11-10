package util

import (
	. "github.com/robertkrimen/terst"
	"testing"
)

func TestFileExistsPositive(t *testing.T) {
	Terst(t)
	ex := FileExists("/")
	Is(ex, true)
}

func TestFileExistsNegative(t *testing.T) {
	Terst(t)
	ex := FileExists("/no-such-file")
	Is(ex, false)
}

func TestAppendIfNeededWithout(t *testing.T) {
	Terst(t)
	str := AppendIfNeeded("string", '/')
	Is(str, "string/")
}

func TestAppendIfNeededWithAlready(t *testing.T) {
	Terst(t)
	str := AppendIfNeeded("string/", '/')
	Is(str, "string/")
}
