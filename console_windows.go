//go:build windows

package cursor

import (
	"io"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

// golang.org/x/sys/windows does not wrap these console functions.
var (
	kernel32                       = windows.NewLazySystemDLL("kernel32.dll")
	procGetConsoleCursorInfo       = kernel32.NewProc("GetConsoleCursorInfo")
	procSetConsoleCursorInfo       = kernel32.NewProc("SetConsoleCursorInfo")
	procFillConsoleOutputCharacter = kernel32.NewProc("FillConsoleOutputCharacterW")
)

// consoleCursorInfo is CONSOLE_CURSOR_INFO.
type consoleCursorInfo struct {
	size    uint32
	visible int32
}

// console moves the cursor of a legacy Windows console through the console API.
type console struct {
	handle windows.Handle
}

// legacyConsole reports whether w is a Windows console that does not understand
// ANSI escape sequences. It first tries to turn on virtual terminal processing,
// which every Windows 10 and later console supports.
func legacyConsole(w io.Writer) (console, bool) {
	f, ok := w.(interface{ Fd() uintptr })
	if !ok {
		return console{}, false
	}

	handle := windows.Handle(f.Fd())

	var mode uint32
	if windows.GetConsoleMode(handle, &mode) != nil {
		// Not a console, for example a file or a pipe.
		return console{}, false
	}

	if mode&windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING != 0 {
		return console{}, false
	}

	if windows.SetConsoleMode(handle, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING) == nil {
		return console{}, false
	}

	return console{handle: handle}, true
}

func (c console) info() (windows.ConsoleScreenBufferInfo, bool) {
	var info windows.ConsoleScreenBufferInfo

	return info, windows.GetConsoleScreenBufferInfo(c.handle, &info) == nil
}

func (c console) move(x, y int) {
	info, ok := c.info()
	if !ok {
		return
	}

	pos := info.CursorPosition
	pos.X = clamp(int(pos.X)+x, 0, int(info.Size.X)-1)
	pos.Y = clamp(int(pos.Y)+y, 0, int(info.Size.Y)-1)

	_ = windows.SetConsoleCursorPosition(c.handle, pos)
}

func (c console) horizontalAbsolute(n int) {
	info, ok := c.info()
	if !ok {
		return
	}

	pos := info.CursorPosition
	pos.X = clamp(n, 0, int(info.Size.X)-1)

	_ = windows.SetConsoleCursorPosition(c.handle, pos)
}

func (c console) setVisible(visible bool) {
	var cursorInfo consoleCursorInfo

	ok, _, _ := procGetConsoleCursorInfo.Call(uintptr(c.handle), uintptr(unsafe.Pointer(&cursorInfo))) //nolint:gosec // Required by the Windows API call.
	if ok == 0 {
		return
	}

	cursorInfo.visible = 0
	if visible {
		cursorInfo.visible = 1
	}

	_, _, _ = procSetConsoleCursorInfo.Call(uintptr(c.handle), uintptr(unsafe.Pointer(&cursorInfo))) //nolint:gosec // Required by the Windows API call.
}

func (c console) clearLine() {
	info, ok := c.info()
	if !ok {
		return
	}

	c.fill(windows.Coord{X: 0, Y: info.CursorPosition.Y}, int(info.Size.X))
}

func (c console) clear() {
	info, ok := c.info()
	if !ok {
		return
	}

	c.fill(info.CursorPosition, int(info.Size.X)-int(info.CursorPosition.X))
}

// fill writes length spaces starting at start without moving the cursor.
func (c console) fill(start windows.Coord, length int) {
	if length <= 0 {
		return
	}

	var written uint32

	// COORD is passed by value, packed into a single register.
	coord := uintptr(uint16(start.X)) | uintptr(uint16(start.Y))<<16 //nolint:gosec // Bit packing, not a conversion of the value.

	_, _, _ = procFillConsoleOutputCharacter.Call(uintptr(c.handle), uintptr(' '), uintptr(length), coord, uintptr(unsafe.Pointer(&written))) //nolint:gosec // Required by the Windows API call.
}

func clamp(v, lower, upper int) int16 {
	return int16(max(lower, min(v, upper))) //nolint:gosec // Console coordinates are bounded by the int16 buffer size.
}

// normalizeNewlines replaces every "\n" that is not part of "\r\n" with "\r\n",
// so multi-line content starts each line in the first column on Windows.
func normalizeNewlines(content string) string {
	if !strings.Contains(content, "\n") {
		return content
	}

	var b strings.Builder

	b.Grow(len(content) + strings.Count(content, "\n"))

	var last rune

	for _, r := range content {
		if r == '\n' && last != '\r' {
			b.WriteByte('\r')
		}

		b.WriteRune(r)

		last = r
	}

	return b.String()
}
