package cursor

import (
	"io"
	"testing"

	"github.com/atomicgo/testutil"
)

func equal(t *testing.T, output string, input func(w io.Writer) error) {
	t.Helper()
	in, _ := testutil.CaptureStdout(input)
	if in != output {
		t.Errorf("Input: '%s' is not equal to '%s'", in, output)
	}
}

func TestClearLine(t *testing.T) {
	equal(t, "\x1b[2K", func(w io.Writer) error {
		ClearLine()

		return nil
	})
}

func TestClearLines(t *testing.T) {
	equal(t, "\x1b[2K\x1b[1A\x1b[2K\x1b[1A\x1b[2K\x1b[1A\x1b[2K", func(w io.Writer) error {
		ClearLines(3)

		return nil
	})
}

func TestClearScreen(t *testing.T) {
	equal(t, "\x1b[2J\x1b[1;1H", func(w io.Writer) error {
		ClearScreen()

		return nil
	})
}

func TestCloseAlternativeScreen(t *testing.T) {
	equal(t, "\x1b[?1049l", func(w io.Writer) error {
		CloseAlternativeScreen()

		return nil
	})
}

func TestDown(t *testing.T) {
	equal(t, "\x1b[6B", func(w io.Writer) error {
		Down(6)

		return nil
	})
}

func TestHide(t *testing.T) {
	equal(t, "\x1b[?25l", func(w io.Writer) error {
		Hide()

		return nil
	})
}

func TestLeft(t *testing.T) {
	equal(t, "\x1b[6D", func(w io.Writer) error {
		Left(6)

		return nil
	})
}

func TestMove(t *testing.T) {
	equal(t, "\x1b[6;7H", func(w io.Writer) error {
		Move(6, 7)

		return nil
	})
}

func TestNextLine(t *testing.T) {
	equal(t, "\x1b[6E", func(w io.Writer) error {
		NextLine(6)

		return nil
	})
}

func TestOpenAlternativeScreen(t *testing.T) {
	equal(t, "\x1b[?1049h", func(w io.Writer) error {
		OpenAlternativeScreen()

		return nil
	})
}

func TestPrevLine(t *testing.T) {
	equal(t, "\x1b[6F", func(w io.Writer) error {
		PrevLine(6)

		return nil
	})
}

func TestRestorePosition(t *testing.T) {
	equal(t, "\x1b[u", func(w io.Writer) error {
		RestorePosition()

		return nil
	})
}

func TestRight(t *testing.T) {
	equal(t, "\x1b[6C", func(w io.Writer) error {
		Right(6)

		return nil
	})
}

func TestSavePosition(t *testing.T) {
	equal(t, "\x1b[s", func(w io.Writer) error {
		SavePosition()

		return nil
	})
}

func TestShow(t *testing.T) {
	equal(t, "\x1b[?25h", func(w io.Writer) error {
		Show()

		return nil
	})
}

func TestUp(t *testing.T) {
	equal(t, "\x1b[6A", func(w io.Writer) error {
		Up(6)

		return nil
	})
}
