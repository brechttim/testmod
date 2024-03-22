package modules

import (
	"io"
	"os"
	"strconv"
	"testing"
)

func TestNewContent(t *testing.T) {
	c := NewContent()

	if c == nil {
		t.Errorf("NewContent() returned nil, expected a non-nil value")
		return
	}

	if c.Num != 0 {
		t.Errorf("NewContent() returned a Content object with Num=%d, expected Num=0", c.Num)
	}

	if c.Str != "" {
		t.Errorf("NewContent() returned a Content object with Str=%s, expected Str=\"\"", c.Str)
	}
}

func TestContent_SetNum(t *testing.T) {
	c := NewContent()
	num := 42

	c.SetNum(num)

	if c.Num != num {
		t.Errorf("SetNum() did not set the correct value. Expected %d, got %d", num, c.Num)
	}
}

func TestContent_SetStr(t *testing.T) {
	c := NewContent()
	str := "Hello, World!"

	c.SetStr(str)

	if c.Str != str {
		t.Errorf("SetStr() did not set the correct value. Expected %s, got %s", str, c.Str)
	}
}

func TestContent_PrintStr(t *testing.T) {
	c := NewContent()
	str := "Hello, World!"
	c.SetStr(str)

	// Redirect stdout to capture the printed output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	c.PrintStr()

	// Read the printed output from the pipe
	w.Close()
	out, _ := io.ReadAll(io.Reader(r))
	os.Stdout = old

	if string(out) != str+"\n" {
		t.Errorf("PrintStr() did not print the correct value. Expected %s, got %s", str, string(out))
	}
}

func TestContent_PrintNum(t *testing.T) {
	c := NewContent()
	num := 42
	c.SetNum(num)

	// Redirect stdout to capture the printed output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	c.PrintNum()

	// Read the printed output from the pipe
	w.Close()
	out, _ := io.ReadAll(io.Reader(r))
	os.Stdout = old

	if string(out) != strconv.Itoa(num)+"\n" {
		t.Errorf("PrintNum() did not print the correct value. Expected %d, got %s", num, string(out))
	}
}
