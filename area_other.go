//go:build !windows
// +build !windows

package cursor

import (
	"fmt"
	"strings"
)

// Update overwrites the content of the Area and adjusts its height based on content.
func (area *Area) Update(content string) {
	area.Clear()
	fmt.Fprint(area.writer, content)
	// Detect height of cursor area
	area.height = strings.Count(content, "\n")
}
