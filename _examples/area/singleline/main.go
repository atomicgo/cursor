package main

import (
	"fmt"
	"time"

	"atomicgo.dev/cursor"
)

func main() {
	fmt.Println("Single line cursor area demo")
	fmt.Println("----------------------------")

	header := "This is a singleline without newline"
	area := cursor.NewArea()
	area.Update(header)
	for i := 1; i < 10; i++ {
		time.Sleep(1 * time.Second)
		area.Update(fmt.Sprintf("%s: %d", header, i))
	}

	time.Sleep(1 * time.Second)
	fmt.Println("\n--- DONE")
}
