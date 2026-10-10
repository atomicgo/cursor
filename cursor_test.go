package cursor

import (
	"bytes"
	"os"
	"sync"
	"testing"
)

// syncBuffer is a bytes.Buffer that is safe for concurrent use.
type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.buf.Write(p)
}

func (b *syncBuffer) take() string {
	b.mu.Lock()
	defer b.mu.Unlock()

	s := b.buf.String()
	b.buf.Reset()

	return s
}

func TestCursorWritesANSISequences(t *testing.T) {
	t.Parallel()

	var buf syncBuffer

	c := NewCursor().WithWriter(&buf)

	tests := []struct {
		name string
		do   func()
		want string
	}{
		{"Up", func() { c.Up(2) }, "\x1b[2A"},
		{"Down", func() { c.Down(2) }, "\x1b[2B"},
		{"Right", func() { c.Right(2) }, "\x1b[2C"},
		{"Left", func() { c.Left(2) }, "\x1b[2D"},
		{"Up(0)", func() { c.Up(0) }, ""},
		{"Down(-1)", func() { c.Down(-1) }, ""},
		{"Hide", c.Hide, "\x1b[?25l"},
		{"Show", c.Show, "\x1b[?25h"},
		{"ClearLine", c.ClearLine, "\x1b[2K"},
		{"Clear", c.Clear, "\x1b[K"},
		{"HorizontalAbsolute", func() { c.HorizontalAbsolute(3) }, "\x1b[4G"},
	}

	for _, tt := range tests {
		tt.do()

		if got := buf.take(); got != tt.want {
			t.Errorf("%s wrote %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestCursorWithNilWriterKeepsWriter(t *testing.T) {
	t.Parallel()

	var buf syncBuffer

	c := NewCursor().WithWriter(&buf).WithWriter(nil)
	c.Up(1)

	if got := buf.take(); got != "\x1b[1A" {
		t.Errorf("got %q, want %q", got, "\x1b[1A")
	}
}

// The global helpers share state, so these tests do not run in parallel.

func useGlobalTarget(t *testing.T) *syncBuffer {
	t.Helper()

	var buf syncBuffer

	SetTarget(&buf)
	Bottom()
	buf.take()

	t.Cleanup(func() {
		Bottom()
		SetTarget(os.Stdout)
	})

	return &buf
}

func TestGlobalHelpers(t *testing.T) {
	buf := useGlobalTarget(t)

	Up(3)
	Down(1)
	Move(2, -1)
	Move(-2, 1)
	HorizontalAbsolute(0)
	ClearLine()
	Hide()
	Show()

	want := "\x1b[3A\x1b[1B\x1b[2C\x1b[1B\x1b[2D\x1b[1A\x1b[1G\x1b[2K\x1b[?25l\x1b[?25h"
	if got := buf.take(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGlobalHeightTracking(t *testing.T) {
	buf := useGlobalTarget(t)

	Up(3)
	Down(1)
	Bottom()

	if got, want := buf.take(), "\x1b[3A\x1b[1B\x1b[2B\x1b[1G"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	// Moving further down than up must not make the height negative.
	Down(10)
	Bottom()
	Up(1)
	Bottom()

	if got, want := buf.take(), "\x1b[10B\x1b[1A\x1b[1B\x1b[1G"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGlobalHelpersConcurrently(t *testing.T) {
	buf := useGlobalTarget(t)

	var wg sync.WaitGroup

	for range 8 {
		wg.Go(func() {
			for range 50 {
				Up(1)
				Down(1)
				Bottom()
				ClearLinesUp(1)
				ClearLinesDown(1)
			}
		})
	}

	wg.Wait()

	if buf.take() == "" {
		t.Error("expected output")
	}
}
