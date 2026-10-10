//go:build !windows

package cursor

import "io"

// console is only implemented on Windows.
type console struct{}

func (console) move(_, _ int)            {}
func (console) horizontalAbsolute(_ int) {}
func (console) setVisible(_ bool)        {}
func (console) clearLine()               {}
func (console) clear()                   {}

// legacyConsole reports whether w needs the Windows console API. Every terminal
// outside of Windows understands ANSI escape sequences.
func legacyConsole(_ io.Writer) (console, bool) {
	return console{}, false
}

// normalizeNewlines returns content unchanged outside of Windows.
func normalizeNewlines(content string) string {
	return content
}
