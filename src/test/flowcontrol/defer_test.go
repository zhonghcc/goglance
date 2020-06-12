package flowcontrol

import (
	"os"
	"testing"
)

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("defer")
	}()
	defer func() {
		t.Log("defer2")
	}()
	defer func() {
		t.Log("defer3")
	}()
	t.Log("do Something")
	defer func() {
		t.Log("defer4")
	}()
}
func TestDeferWithPanic(t *testing.T) {
	defer func() {
		t.Log("defer")
	}()
	t.Log("before wrong")
	panic("wrong wrong wrong")
	t.Log("after wrong")
}

func TestDeferWithExit(t *testing.T) {
	defer func() {
		t.Log("defer")
	}()
	os.Exit(0)
}
