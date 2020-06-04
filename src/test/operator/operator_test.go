package operator

import (
	"testing"
)

func TestString(t *testing.T) {
	var a string = "a"
	var b string = "b"
	t.Log(a + b)
}

func TestBitZero(t *testing.T) {
	a := 7
	check := 0x2
	t.Log(a &^ check)
	check = 6
	t.Log(a &^ check)
}
