package container

import "testing"

func TestMapInit(t *testing.T) {
	var a map[string]string = map[string]string{}
	b := make(map[string]string)
	c := map[string]string{"c": "c"}
	a["a"] = "a"
	b["b"] = "b"
	t.Log(a)
	t.Log(b)
	t.Log(c)

}

func TestMapTranverse(t *testing.T) {
	a := map[string]string{"1": "a", "2": "b", "3": "c"}
	for key, value := range a {
		t.Logf("%s,%s", key, value)
	}
	if item, exist := a["abcd"]; exist == false {
		t.Log("no key found")
	} else {
		t.Log(item, exist)
	}
	if item, exist := a["1"]; exist == false {
		t.Log("no key found")
	} else {
		t.Log(item, exist)
	}
}
