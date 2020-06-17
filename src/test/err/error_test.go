package err

import (
	"errors"
	"testing"
)

func doublePositive(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("negitive")
	}
	return i * 2, nil
}

func TestError(t *testing.T) {
	if v, err := doublePositive(10); err == nil {
		t.Log(v)
	} else {
		t.Error(err)
	}
	if v, err := doublePositive(-1); err == nil {
		t.Log(v)
	} else {
		t.Error(err)
	}
	v, err := doublePositive(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}
