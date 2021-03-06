package vhdl

import (
	"bufio"
	"bytes"
)

type scanContext struct {
	scanner *bufio.Scanner
	lineNum uint
	line    []byte
}

// scan returns false on EOF.
func (sc *scanContext) scan() bool {
	sc.lineNum += 1

	if !sc.scanner.Scan() {
		return false
	}

	sc.line = sc.scanner.Bytes()

	return true
}

// decomment removes the comment at the end of the line.
func (sc *scanContext) decomment() {
	sc.line = bytes.Split(sc.line, []byte("--"))[0]
}
