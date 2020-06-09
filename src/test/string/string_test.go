package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringInit(t *testing.T) {
	a := "a"
	t.Log(a)
	b := "你"
	t.Log(b)
}

func TestStringFunc(t *testing.T) {
	a := "1,2,3"
	b := strings.Split(a, ",")
	for item := range b {
		t.Log(item)
	}

	c := strings.Join(b, "-")
	t.Log(c)
	if ret, err := strconv.Atoi("+1"); err == nil {
		t.Log(ret)
	} else {
		t.Error(err)
	}
}
