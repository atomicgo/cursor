//go:build windows
// +build windows

package cursor

import (
	"fmt"
	"strings"
)

// Update overwrites the content of the Area and adjusts its height based on content.
// For Windows newlines '\n' in the content  are replaced by '\r\n'
func (area *Area) Update(content string) {
	area.Clear()
	last := ' '
	for _, r := range content {
		if r == '\n' && last != '\r' {
			fmt.Fprint(area.writer, "\r\n")
			continue
		}
		fmt.Fprint(area.writer, string(r))
	}
	// Detect height of cursor area
	area.height = strings.Count(content, "\n")
}
