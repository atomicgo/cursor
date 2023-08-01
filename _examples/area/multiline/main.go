package main

import (
	"fmt"
	"math/rand"
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
	for i := 1; i < 6; i++ {
		if i%2 == 0 {
			content += fmt.Sprintf(" + %d\n", i)
		} else {
			content += fmt.Sprintf(" - line: %d", i)
		}
		time.Sleep(1 * time.Second)
		area.Update(content)
	}

	time.Sleep(1 * time.Second)
	area.Update("Test varying area sizes now")
	time.Sleep(500 * time.Millisecond)
	area.Update(buildContent(1, 2))
	time.Sleep(500 * time.Millisecond)
	area.Update(buildContent(2, 9))
	time.Sleep(500 * time.Millisecond)
	area.Update(buildContent(3, 5))
	time.Sleep(500 * time.Millisecond)
	area.Update(buildContent(4, 0))
	time.Sleep(500 * time.Millisecond)
	area.Update(buildContent(5, 6))
	time.Sleep(500 * time.Millisecond)
	area.Update(buildContent(6, 1))
	time.Sleep(500 * time.Millisecond)
	area.Update(buildContent(7, 3))

	time.Sleep(1 * time.Second)
	fmt.Println("\n--- DONE")
}

func buildContent(idx int, n int) string {
	content := fmt.Sprintf(">>> START OF CONTENT %d/%d <<<\n", idx, n)
	for i := 0; i < n; i++ {
		for i := 0; i < 5; i++ {
			content += words[rand.Intn(len(words))] + " "
		}
		content += "\n"
	}

	return content
}

var words = []string{
	"ball", "summer", "hint", "mountain", "island", "onion", "world",
	"run", "hit", "fly", "swim", "crawl", "build", "dive", "jump",
	"crazy", "funny", "strange", "yellow", "red", "blue", "green", "white",
}
