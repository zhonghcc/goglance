package array

import "testing"

func TestInit(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [3]int{}
	t.Log(a)
	t.Log(b)
	b[1] = 2
	t.Log(b)
}

func TestCompre(t *testing.T) {
	a := [...]int{1, 2, 3}
	// b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 4}
	d := [...]int{1, 2, 3}
	// t.Log(a == b)
	t.Log(a == c)
	t.Log(a == d)

}

func TestSub(t *testing.T) {
	a := [...]int{1, 2, 3}
	t.Log(a[0:1])
	t.Log(a[0:len(a)])
	// t.Log(a[0:-1])

}
