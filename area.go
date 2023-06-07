package cursor

import (
	"fmt"
	"strings"
)

// Area displays content which can be updated on the fly.
// You can use this to create live output, charts, dropdowns, etc.
type Area struct {
	height int
}

// NewArea returns a new Area.
func NewArea() Area {
	return Area{}
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
	fmt.Println(content)
	height = 0

	area.height = len(strings.Split(content, "\n"))
}
