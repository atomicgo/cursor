package cursor

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"
)

func TestHeightChanges(t *testing.T) {
	for i := 0; i < 4; i++ {
		fmt.Println()
	}
	Up(3)
	if height != 3 {
		t.Errorf("height should be 3 but is %d", height)
	}
	Down(3)
	if height != 0 {
		t.Errorf("height should be 0 but is %d", height)
	}
}

func TestHeightCannotBeNegative(t *testing.T) {
	Down(10)
	if height < 0 {
		t.Errorf("height is negative: %d", height)
	}
}

func TestCustomIOWriter(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping these tests on windows")
	}

	var w bytes.Buffer
	SetTarget(&w)

	Up(2)
	expected := "\x1b[2A"
	actual := w.String()
	if expected != actual {
		t.Errorf("wanted: %v, got %v", expected, actual)
	}

	w.Reset()
	Down(2)
	expected = "\x1b[2B"
	actual = w.String()
	if expected != actual {
		t.Errorf("wanted: %v, got %v", expected, actual)
	}

	w.Reset()
	Right(2)
	expected = "\x1b[2C"
	actual = w.String()
	if expected != actual {
		t.Errorf("wanted: %v, got %v", expected, actual)
	}

	w.Reset()
	Left(2)
	expected = "\x1b[2D"
	actual = w.String()
	if expected != actual {
		t.Errorf("wanted: %v, got %v", expected, actual)
	}

	w.Reset()
	Hide()
	expected = "\x1b[?25l"
	actual = w.String()
	if expected != actual {
		t.Errorf("wanted: %v, got %v", expected, actual)
	}

	w.Reset()
	Show()
	expected = "\x1b[?25h"
	actual = w.String()
	if expected != actual {
		t.Errorf("wanted: %v, got %v", expected, actual)
	}

	w.Reset()
	ClearLine()
	expected = "\x1b[2K"
	actual = w.String()
	if expected != actual {
		t.Errorf("wanted: %v, got %v", expected, actual)
	}

	w.Reset()
	HorizontalAbsolute(3)
	expected = "\x1b[4G"
	actual = w.String()
	if expected != actual {
		t.Errorf("wanted: %v, got %v", expected, actual)
	}
}
