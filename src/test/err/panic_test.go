package err

import (
	"errors"
	"testing"
)

func doSomethingWrong() {
	panic(errors.New("i feel panic"))
}

func TestRecover(t *testing.T) {
	defer func() {
		t.Log("try recover")
		if err := recover(); err != nil {
			t.Log("你是个有error的panic", err)
		}
	}()
	doSomethingWrong()
}

func TestPanic(t *testing.T) {
	i := 1
	defer func() {
		t.Log("defer", i)
	}()
	i++
	doSomethingWrong()
	i++
	t.Log("unreachable", i)
}
