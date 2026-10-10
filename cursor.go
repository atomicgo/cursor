package cursor

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Cursor moves the terminal cursor of a writer.
// You can use this to create live output, charts, dropdowns, etc.
//
// A Cursor writes ANSI escape sequences. On legacy Windows consoles that do not
// support them, it uses the Windows console API instead.
//
// All methods are safe for concurrent use.
type Cursor struct {
	mu     sync.Mutex
	writer io.Writer
}

// NewCursor creates a new Cursor instance writing to os.Stdout.
func NewCursor() *Cursor {
	return &Cursor{writer: os.Stdout}
}

// WithWriter sets the writer the cursor writes to and returns the cursor.
// A nil writer is ignored.
func (c *Cursor) WithWriter(w io.Writer) *Cursor {
	if w != nil {
		c.mu.Lock()
		c.writer = w
		c.mu.Unlock()
	}

	return c
}

// Up moves the cursor n lines up relative to the current position.
func (c *Cursor) Up(n int) {
	if n > 0 {
		c.do(func(con console) { con.move(0, -n) }, "\x1b[%dA", n)
	}
}

// Down moves the cursor n lines down relative to the current position.
func (c *Cursor) Down(n int) {
	if n > 0 {
		c.do(func(con console) { con.move(0, n) }, "\x1b[%dB", n)
	}
}

// Right moves the cursor n characters to the right relative to the current position.
func (c *Cursor) Right(n int) {
	if n > 0 {
		c.do(func(con console) { con.move(n, 0) }, "\x1b[%dC", n)
	}
}

// Left moves the cursor n characters to the left relative to the current position.
func (c *Cursor) Left(n int) {
	if n > 0 {
		c.do(func(con console) { con.move(-n, 0) }, "\x1b[%dD", n)
	}
}

// HorizontalAbsolute moves the cursor to n horizontally.
// The position n is absolute to the start of the line.
func (c *Cursor) HorizontalAbsolute(n int) {
	n = max(n, 0)

	// ANSI columns start at 1.
	c.do(func(con console) { con.horizontalAbsolute(n) }, "\x1b[%dG", n+1)
}

// Show the cursor if it was hidden previously.
// Don't forget to show the cursor at least at the end of your application.
// Otherwise the user might have a terminal with a permanently hidden cursor, until they reopen the terminal.
func (c *Cursor) Show() {
	c.do(func(con console) { con.setVisible(true) }, "\x1b[?25h")
}

// Hide the cursor.
// Don't forget to show the cursor at least at the end of your application with Show.
// Otherwise the user might have a terminal with a permanently hidden cursor, until they reopen the terminal.
func (c *Cursor) Hide() {
	c.do(func(con console) { con.setVisible(false) }, "\x1b[?25l")
}

// ClearLine clears the current line.
func (c *Cursor) ClearLine() {
	c.do(func(con console) { con.clearLine() }, "\x1b[2K")
}

// Clear clears the line from the cursor position to its end.
func (c *Cursor) Clear() {
	c.do(func(con console) { con.clear() }, "\x1b[K")
}

// do runs legacy on legacy Windows consoles and writes the ANSI sequence otherwise.
func (c *Cursor) do(legacy func(console), format string, args ...any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	w := c.writer
	if w == nil {
		w = os.Stdout
	}

	if con, ok := legacyConsole(w); ok {
		legacy(con)

		return
	}

	_, _ = fmt.Fprintf(w, format, args...)
}
