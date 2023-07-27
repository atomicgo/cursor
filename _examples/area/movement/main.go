package main

import (
	"fmt"
	"time"

	"atomicgo.dev/cursor"
)

func main() {
	fmt.Println("Multiline cursor area demo")
	fmt.Println("--------------------------")

	area := cursor.NewArea()
	content := `Start content with multiple rows
	1. Row1
	2. Row2
	3. Row3
	---
	`
	area.Update(content)

	time.Sleep(1 * time.Second)
	area.Up(3)
	area.Move(3, 0)
	fmt.Print("Replaced row 2")

	time.Sleep(1 * time.Second)
	area.StartOfLine()
	area.Move(8, -2)
	fmt.Print("Appended stuff")

	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content)")

	time.Sleep(1 * time.Second)
	area.Up(6)
	area.Update(content + "(restored content after cursor up out of bounds)")

	time.Sleep(1 * time.Second)
	area.Down(6)
	area.Update(content + "(restored content after cursor down out of bounds)")

	time.Sleep(1 * time.Second)
	area.Update("Show only single line now")

	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content after single line)")

	time.Sleep(1 * time.Second)
	area.Update("")

	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content after empty line)")

	time.Sleep(1 * time.Second)
	fmt.Println("\n--- DONE")
}
