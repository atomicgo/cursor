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

	time.Sleep(2 * time.Second)
	area.Up(3)
	area.Move(3, 0)
	fmt.Print("Replaced row 2")

	time.Sleep(2 * time.Second)
	area.StartOfLine()
	area.Move(8, -2)
	fmt.Print("Appended stuff")

	time.Sleep(2 * time.Second)
	area.Update(content + "(restored content)")

	time.Sleep(2 * time.Second)
	fmt.Println("\n--- DONE")
}
