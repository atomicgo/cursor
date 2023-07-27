package main

import (
	"fmt"
	"strings"
	"time"

	"atomicgo.dev/cursor"
)

func main() {

	fmt.Println("=======================")
	fmt.Println("SIMPLE CURSOR AREA DEMO")
	fmt.Println("=======================")

	area := cursor.NewArea()
	lines := []string{
		"Example content for an area",
		"with multiple lines of content",
		"and nothing else",
	}
	area.Update(strings.Join(lines, "\n"))

	time.Sleep(2 * time.Second)
	lines = append(lines, "and even more lines <--")

	area.Update(strings.Join(lines, "\n"))
	time.Sleep(2 * time.Second)
	lines[2] = "another line <--"
	area.Update(strings.Join(lines, "\n"))

	fmt.Println("\n<<< DONE")
}
