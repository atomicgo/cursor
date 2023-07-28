package main

import (
	"fmt"
	"time"

	"atomicgo.dev/cursor"
)

func main() {
	fmt.Println("Single line cursor area demo")
	fmt.Println("----------------------------")

	area := cursor.NewArea()

	header := "This is a singleline without newline"
	area.Update(header)
	for i := 1; i < 6; i++ {
		time.Sleep(1 * time.Second)
		area.Update(fmt.Sprintf("%s: %d", header, i))
	}

	header = "This is a singleline with newline"
	area.Update(header + "\n")
	for i := 1; i < 6; i++ {
		time.Sleep(1 * time.Second)
		area.Update(fmt.Sprintf("%s: %d\n", header, i))
	}

	time.Sleep(1 * time.Second)
	fmt.Println("\n--- DONE")
}
