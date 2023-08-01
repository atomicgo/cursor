package cursor

import (
	"fmt"
	"testing"
)

func TestHeightChanges(t *testing.T) {
	for i := 0; i < 4; i++ {
		fmt.Println()
	}
	Up(3)

	if autoheight != 3 {
		t.Errorf("height should be 3 but is %d", autoheight)
	}

	Down(3)

	if autoheight != 0 {
		t.Errorf("height should be 0 but is %d", autoheight)
	}
}

func TestHeightCannotBeNegative(t *testing.T) {
	Down(10)

	if autoheight < 0 {
		t.Errorf("height is negative: %d", autoheight)
	}
}
