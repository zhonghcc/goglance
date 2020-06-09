package string

import (
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
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

func TestCodec(t *testing.T) {
	a := "hello"
	b := "你好"
	t.Logf("%s,len=%d,rune=%d", a, len(a), len([]rune(a)))
	t.Logf("%s,len=%d,rune=%d", b, len(b), len([]rune(b)))
	t.Log(utf8.RuneCountInString(b))

	c := []byte(b)[3:6]
	t.Log(string(c))

	d := []rune(b)[1]
	t.Log(string(d))
}
