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
	header := "This is a multiline demo\nwith 2 lines:\n"
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
	fmt.Println("X\n--- DONE")
}
