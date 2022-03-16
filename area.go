package cursor

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

// Area displays content which can be updated on the fly.
// You can use this to create live output, charts, dropdowns, etc.
type Area struct {
	height int
	writer io.Writer
}

// NewArea returns a new Area.
func NewArea() Area {
	return Area{}
}

// WithCustomWriter sets the custom writer
func (area *Area) WithCustomWriter(writer io.Writer) *Area {
	area.writer = writer

	return area
}

// Clear clears the content of the Area.
func (area *Area) Clear() {
	Bottom()
	if area.height > 0 {
		ClearLinesUp(area.height)
	}
}

// Update overwrites the content of the Area.
func (area *Area) Update(content string) {
	area.Clear()
	lines := strings.Split(content, "\n")
	if runtime.GOOS == "windows" {
		for _, line := range lines {
			fmt.Print(line)
			StartOfLineDown(1)
		}
	} else {
		for _, line := range lines {
			fmt.Fprintln(area.writer, line)
		}
	}
	height = 0

	area.height = len(lines)
}
