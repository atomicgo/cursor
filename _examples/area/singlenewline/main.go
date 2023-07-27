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
	header := "This is a singleline with newline"
	area.Update(header)
	content := header
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			content += fmt.Sprintf(" + %d\n", i)
		} else {
			content += fmt.Sprintf(" - line: %d", i)
		}
		time.Sleep(1 * time.Second)
		area.Update(content)
	}
	fmt.Println("x\n--- DONE")
}
