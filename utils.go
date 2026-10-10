package cursor

import (
	"io"
	"sync"
)

//
// Helpers for global cursor handling on os.Stdout
//

var (
	cursor = NewCursor() //nolint:gochecknoglobals // The package level helpers share one cursor.

	// autoheightMu guards autoheight, the number of lines the global cursor was moved up.
	autoheightMu sync.Mutex //nolint:gochecknoglobals // Guards autoheight.
	autoheight   int        //nolint:gochecknoglobals // Used by Bottom.
)

// Writer is an io.Writer with a file descriptor, such as os.Stdout.
//
// Cursor and Area accept any io.Writer. A file descriptor is only used on
// Windows, to move the cursor of legacy consoles that do not support ANSI
// escape sequences.
type Writer interface {
	io.Writer
	Fd() uintptr
}

// SetTarget sets the output target of the default cursor to the provided writer.
func SetTarget(w io.Writer) {
	cursor.WithWriter(w)
}

// Up moves the cursor n lines up relative to the current position.
func Up(n int) {
	if n <= 0 {
		return
	}

	autoheightMu.Lock()
	defer autoheightMu.Unlock()

	cursor.Up(n)
	autoheight += n
}

// Down moves the cursor n lines down relative to the current position.
func Down(n int) {
	if n <= 0 {
		return
	}

	autoheightMu.Lock()
	defer autoheightMu.Unlock()

	cursor.Down(n)
	autoheight = max(autoheight-n, 0)
}

// Right moves the cursor n characters to the right relative to the current position.
func Right(n int) {
	cursor.Right(n)
}

// Left moves the cursor n characters to the left relative to the current position.
func Left(n int) {
	cursor.Left(n)
}

// HorizontalAbsolute moves the cursor to n horizontally.
// The position n is absolute to the start of the line.
func HorizontalAbsolute(n int) {
	cursor.HorizontalAbsolute(n)
}

// Show the cursor if it was hidden previously.
// Don't forget to show the cursor at least at the end of your application.
// Otherwise the user might have a terminal with a permanently hidden cursor, until they reopen the terminal.
func Show() {
	cursor.Show()
}

// Hide the cursor.
// Don't forget to show the cursor at least at the end of your application with Show.
// Otherwise the user might have a terminal with a permanently hidden cursor, until they reopen the terminal.
func Hide() {
	cursor.Hide()
}

// ClearLine clears the current line.
func ClearLine() {
	cursor.ClearLine()
}

// Clear clears the line from the cursor position to its end.
func Clear() {
	cursor.Clear()
}

// Bottom moves the cursor to the bottom of the terminal.
// This is done by calculating how many lines were moved by Up and Down.
func Bottom() {
	autoheightMu.Lock()
	n := autoheight
	autoheightMu.Unlock()

	if n > 0 {
		Down(n)
		StartOfLine()
	}
}

// StartOfLine moves the cursor to the start of the current line.
func StartOfLine() {
	HorizontalAbsolute(0)
}

// StartOfLineDown moves the cursor down by n lines, then moves to cursor to the start of the line.
func StartOfLineDown(n int) {
	Down(n)
	StartOfLine()
}

// StartOfLineUp moves the cursor up by n lines, then moves to cursor to the start of the line.
func StartOfLineUp(n int) {
	Up(n)
	StartOfLine()
}

// UpAndClear moves the cursor up by n lines, then clears the line.
func UpAndClear(n int) {
	Up(n)
	ClearLine()
}

// DownAndClear moves the cursor down by n lines, then clears the line.
func DownAndClear(n int) {
	Down(n)
	ClearLine()
}

// Move moves the cursor relative by x and y.
func Move(x, y int) {
	if x > 0 {
		Right(x)
	} else if x < 0 {
		Left(-x)
	}

	if y > 0 {
		Up(y)
	} else if y < 0 {
		Down(-y)
	}
}

// ClearLinesUp clears n lines upwards from the current position and moves the cursor.
func ClearLinesUp(n int) {
	for range n {
		UpAndClear(1)
	}
}

// ClearLinesDown clears n lines downwards from the current position and moves the cursor.
func ClearLinesDown(n int) {
	for range n {
		DownAndClear(1)
	}
}
