package flowcontrol

import "testing"

func TestIf(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	if a == b {
		t.Log("a equals b")
	}
}

func TestSwitch1(t *testing.T) {
	a := "Monday"
	switch a {
	case "Monday":
		t.Log(a)
	case "TuesDay":
		t.Log(a)
	default:
		t.Error("not known")
	}
	a = "TuesDay"
	switch a {
	case "Monday", "TuesDay":
		t.Log(a)
	default:
		t.Error("not known")

	}
}

func TestSwitch2(t *testing.T) {
	a := -100
	switch {
	// case a < 0 || a > 100:
	// 	t.Log("wrong score")
	case a < 60 && a >= 0:
		t.Log("fail")
	case a >= 60 && a < 80:
		t.Log("so so")
	case a >= 100:
		t.Log("wonderful")
	default:
		t.Log("anothor wrong score")
	}
}
