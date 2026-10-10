package cursor

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Area displays content which can be updated on the fly.
// You can use this to create live output, charts, dropdowns, etc.
//
// Create an Area with NewArea. All methods are safe for concurrent use, except
// WithWriter, which is meant to configure the Area before it is used.
type Area struct {
	mu         *sync.Mutex
	height     int
	writer     io.Writer
	cursor     *Cursor
	cursorPosY int
}

// zeroAreaMu guards Areas that were not created with NewArea.
var zeroAreaMu sync.Mutex //nolint:gochecknoglobals // Fallback lock for zero value Areas.

// NewArea returns a new Area that writes to os.Stdout.
func NewArea() Area {
	return Area{
		mu:     &sync.Mutex{},
		writer: os.Stdout,
		cursor: NewCursor(),
	}
}

// WithWriter returns a copy of the Area that writes its content and its cursor
// movements to writer. A nil writer is ignored.
func (area Area) WithWriter(writer io.Writer) Area {
	area.mu = &sync.Mutex{}

	if writer != nil {
		area.writer = writer
		area.cursor = NewCursor().WithWriter(writer)
	}

	return area
}

// lock locks the Area and makes sure a zero value Area is usable.
func (area *Area) lock() func() {
	mu := area.mu
	if mu == nil {
		mu = &zeroAreaMu
	}

	mu.Lock()

	if area.writer == nil {
		area.writer = os.Stdout
	}

	if area.cursor == nil {
		area.cursor = NewCursor().WithWriter(area.writer)
	}

	return mu.Unlock
}

// Clear clears the content of the Area.
func (area *Area) Clear() {
	defer area.lock()()

	area.clear()
}

// Update overwrites the content of the Area and adjusts its height based on content.
func (area *Area) Update(content string) {
	defer area.lock()()

	area.clear()
	_, _ = fmt.Fprint(area.writer, normalizeNewlines(content))
	area.cursorPosY = 0
	area.height = strings.Count(content, "\n")
}

// Up moves the cursor of the area up n lines, but not above the top of the Area.
func (area *Area) Up(n int) {
	defer area.lock()()

	area.up(n)
}

// Down moves the cursor of the area down n lines, but not below the bottom of the Area.
func (area *Area) Down(n int) {
	defer area.lock()()

	area.down(n)
}

// Bottom moves the cursor to the bottom of the area.
// This is done by calculating how many lines were moved by Up and Down.
func (area *Area) Bottom() {
	defer area.lock()()

	area.bottom()
}

// Top moves the cursor to the top of the area.
// This is done by calculating how many lines were moved by Up and Down.
func (area *Area) Top() {
	defer area.lock()()

	area.up(area.height - area.cursorPosY)
}

// StartOfLine moves the cursor to the start of the current line.
func (area *Area) StartOfLine() {
	defer area.lock()()

	area.cursor.HorizontalAbsolute(0)
}

// StartOfLineDown moves the cursor down by n lines, then moves to cursor to the start of the line.
func (area *Area) StartOfLineDown(n int) {
	defer area.lock()()

	area.down(n)
	area.cursor.HorizontalAbsolute(0)
}

// StartOfLineUp moves the cursor up by n lines, then moves to cursor to the start of the line.
func (area *Area) StartOfLineUp(n int) {
	defer area.lock()()

	area.up(n)
	area.cursor.HorizontalAbsolute(0)
}

// UpAndClear moves the cursor up by n lines, then clears the line.
func (area *Area) UpAndClear(n int) {
	defer area.lock()()

	area.up(n)
	area.cursor.ClearLine()
}

// DownAndClear moves the cursor down by n lines, then clears the line.
func (area *Area) DownAndClear(n int) {
	defer area.lock()()

	area.down(n)
	area.cursor.ClearLine()
}

// Move moves the cursor relative by x and y.
func (area *Area) Move(x, y int) {
	defer area.lock()()

	if x > 0 {
		area.cursor.Right(x)
	} else if x < 0 {
		area.cursor.Left(-x)
	}

	if y > 0 {
		area.up(y)
	} else if y < 0 {
		area.down(-y)
	}
}

// ClearLinesUp clears n lines upwards from the current position and moves the cursor.
func (area *Area) ClearLinesUp(n int) {
	defer area.lock()()

	area.clearLinesUp(n)
}

// ClearLinesDown clears n lines downwards from the current position and moves the cursor.
func (area *Area) ClearLinesDown(n int) {
	defer area.lock()()

	area.cursor.HorizontalAbsolute(0)
	area.cursor.ClearLine()

	for range n {
		area.down(1)
		area.cursor.ClearLine()
	}
}

func (area *Area) clear() {
	if area.height > 0 {
		area.bottom()
		area.clearLinesUp(area.height)
		area.cursor.HorizontalAbsolute(0)
	} else {
		area.cursor.HorizontalAbsolute(0)
		area.cursor.ClearLine()
	}
}

func (area *Area) clearLinesUp(n int) {
	area.cursor.HorizontalAbsolute(0)
	area.cursor.ClearLine()

	for range n {
		area.up(1)
		area.cursor.ClearLine()
	}
}

// up moves the cursor up. cursorPosY counts the lines above the bottom of the Area.
func (area *Area) up(n int) {
	n = min(n, area.height-area.cursorPosY)
	if n > 0 {
		area.cursor.Up(n)
		area.cursorPosY += n
	}
}

func (area *Area) down(n int) {
	n = min(n, area.cursorPosY)
	if n > 0 {
		area.cursor.Down(n)
		area.cursorPosY -= n
	}
}

func (area *Area) bottom() {
	area.down(area.cursorPosY)
}
