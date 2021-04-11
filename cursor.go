package cursor

import (
	"fmt"
	"strings"
)

const (
	ansiPrefix = "\x1b["

	upTemplate    = "%dA"
	downTemplate  = "%dB"
	leftTemplate  = "%dD"
	rightTemplate = "%dC"

	nextLineTemplate     = "%dE"
	previousLineTemplate = "%dF"

	positionTemplate        = "%d;%dH"
	savePositionTemplate    = "s"
	restorePositionTemplate = "u"

	openAlternativeScreenTemplate  = "?1049h"
	closeAlternativeScreenTemplate = "?1049l"

	clearDisplayTemplate = "%dJ"
	clearLineTemplate    = "%dK"

	showTemplate = "?25h"
	hideTemplate = "?25l"
)

// Move cursor.

// Up moves the cursor n cells up. If the cursor is already at the edge of the screen, this has no effect.
func Up(n int) {
	fmt.Printf(ansiPrefix+upTemplate, n)
}

// Down moves the cursor n cells down. If the cursor is already at the edge of the screen, this has no effect.
func Down(n int) {
	fmt.Printf(ansiPrefix+downTemplate, n)
}

// Left moves the cursor n cells left. If the cursor is already at the edge of the screen, this has no effect.
func Left(n int) {
	fmt.Printf(ansiPrefix+leftTemplate, n)
}

// Right moves the cursor n cells right. If the cursor is already at the edge of the screen, this has no effect.
func Right(n int) {
	fmt.Printf(ansiPrefix+rightTemplate, n)
}

// Move moves the cursor to a specific row and column.
func Move(row int, column int) {
	fmt.Printf(ansiPrefix+positionTemplate, row, column)
}

func NextLine(n int) {
	fmt.Printf(ansiPrefix+nextLineTemplate, n)
}

func PrevLine(n int) {
	fmt.Printf(ansiPrefix+previousLineTemplate, n)
}

func SavePosition() {
	fmt.Print(ansiPrefix + savePositionTemplate)
}

func RestorePosition() {
	fmt.Print(ansiPrefix + restorePositionTemplate)
}

// Alternative screen.

func OpenAlternativeScreen() {
	fmt.Print(ansiPrefix + openAlternativeScreenTemplate)
}

func CloseAlternativeScreen() {
	fmt.Print(ansiPrefix + closeAlternativeScreenTemplate)
}

// Erase

func ClearScreen() {
	fmt.Printf(ansiPrefix+clearDisplayTemplate, 2)
	Move(1, 1)
}

func ClearLine() {
	fmt.Printf(ansiPrefix+clearLineTemplate, 2)
}

func ClearLines(n int) {
	clear := fmt.Sprintf(ansiPrefix+clearLineTemplate, 2)
	fmt.Print(clear + strings.Repeat(fmt.Sprintf(ansiPrefix+upTemplate, 1)+clear, n))
}

// Show and hide cursor

func Show() {
	fmt.Printf(ansiPrefix + showTemplate)
}

func Hide() {
	fmt.Printf(ansiPrefix + hideTemplate)
}
