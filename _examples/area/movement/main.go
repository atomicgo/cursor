package main

import (
	"fmt"
	"time"

	"atomicgo.dev/cursor"
)

func main() {
	fmt.Println("Cursor area movement demo")
	fmt.Println("--------------------------")

	area := cursor.NewArea()
	content := `Start content with some rows
	1. Row1
	2. Row2
	---
	`
	area.Update(content)

	time.Sleep(1 * time.Second)
	area.Up(2)
	area.Move(3, 0)
	fmt.Print("Replaced row 2")

	time.Sleep(1 * time.Second)
	area.StartOfLine()
	area.Move(8, -1)
	fmt.Print("3. Appended row")

	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content after move)")

	time.Sleep(1 * time.Second)
	area.Up(6)
	fmt.Print("<<< AFTER Up(6)")
	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content after cursor up out of bounds)")

	time.Sleep(1 * time.Second)
	area.Down(6)
	fmt.Print("<<< AFTER Down(6)")
	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content after cursor down out of bounds)")

	time.Sleep(1 * time.Second)
	area.Top()
	fmt.Print("<<< AFTER Top()")
	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content after cursor top)")

	time.Sleep(1 * time.Second)
	area.Bottom()
	fmt.Print("<<< AFTER Bottom()")
	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content after cursor bottom)")

	time.Sleep(1 * time.Second)
	area.Update("")
	time.Sleep(1 * time.Second)
	area.Update(content + "(restored content after empty line)")

	time.Sleep(1 * time.Second)
	fmt.Println("\n--- DONE")
}
