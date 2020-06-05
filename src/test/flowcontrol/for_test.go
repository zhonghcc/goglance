package flowcontrol

import "testing"

func TestFor(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 8}
	for _, item := range a {
		t.Log(item)
	}

	i := 0
	for {
		if i > 100 {
			break
		}
		t.Log(i)
		i++
	}
}
