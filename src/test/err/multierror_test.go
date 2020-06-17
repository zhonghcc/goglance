package err

import (
	"errors"
	"testing"
)

var nagitiveError error = errors.New("less than 0")
var tooBigError error = errors.New("too big")

func doSomeCheck(i int) (int, error) {
	if i < 0 {
		return 0, nagitiveError
	}
	if i > 100 {
		return 0, tooBigError
	}
	return i * 2, nil
}

func TestMultiError(t *testing.T) {
	a := -1
	if v, err := doSomeCheck(-1); err == nil {
		t.Log(v)
	} else {
		switch err {
		case nagitiveError:
			t.Error("nagitive", a)
		case tooBigError:
			t.Error("too big", a)
		default:
			t.Error("unknown error", err)
		}
	}
}
