package cursor

import (
	"os"
	"strings"
	"sync"
	"testing"
)

func TestAreaUpdate(t *testing.T) {
	t.Parallel()

	var buf syncBuffer

	area := NewArea().WithWriter(&buf)

	area.Update("a\nb\n")

	if got, want := buf.take(), "\x1b[1G\x1b[2Ka\nb\n"; got != normalizeNewlines(want) {
		t.Errorf("first update wrote %q, want %q", got, want)
	}

	// The second update clears the two lines of the first one.
	area.Update("c")

	want := "\x1b[1G\x1b[2K\x1b[1A\x1b[2K\x1b[1A\x1b[2K\x1b[1Gc"
	if got := buf.take(); got != want {
		t.Errorf("second update wrote %q, want %q", got, want)
	}
}

func TestAreaMovementStaysInside(t *testing.T) {
	t.Parallel()

	var buf syncBuffer

	area := NewArea().WithWriter(&buf)
	area.Update("1\n2\n3\n")
	buf.take()

	area.Up(10)
	area.Up(1)
	area.Down(1)
	area.Down(10)
	area.Top()
	area.Bottom()
	area.Down(1)

	want := "\x1b[3A\x1b[1B\x1b[2B\x1b[3A\x1b[3B"
	if got := buf.take(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAreaWithWriterDoesNotChangeGlobalCursor(t *testing.T) {
	var global syncBuffer

	SetTarget(&global)
	t.Cleanup(func() { SetTarget(os.Stdout) })

	var buf syncBuffer

	area := NewArea().WithWriter(&buf)
	area.Update("x\n")
	area.Up(1)

	if got := global.take(); got != "" {
		t.Errorf("area wrote %q to the global cursor", got)
	}

	if !strings.Contains(buf.take(), "\x1b[1A") {
		t.Error("area did not move its own cursor")
	}
}

func TestZeroValueArea(t *testing.T) {
	var buf syncBuffer

	area := Area{writer: &buf}
	area.Update("x\n")
	area.Up(1)

	if got := buf.take(); !strings.HasSuffix(got, "\x1b[1A") {
		t.Errorf("got %q", got)
	}
}

func TestAreaConcurrentUse(t *testing.T) {
	t.Parallel()

	var buf syncBuffer

	area := NewArea().WithWriter(&buf)

	var wg sync.WaitGroup

	for i := range 8 {
		wg.Go(func() {
			for j := range 50 {
				area.Update(strings.Repeat("line\n", (i+j)%4))
				area.Up(1)
				area.Move(1, -1)
				area.StartOfLineUp(1)
				area.Top()
				area.Bottom()
				area.ClearLinesDown(1)
				area.Clear()
			}
		})
	}

	wg.Wait()

	if buf.take() == "" {
		t.Error("expected output")
	}
}
