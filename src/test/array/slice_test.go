package array

import "testing"

func TestSliceInit(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{}
	t.Log(a)
	t.Logf("len=%d,cap=%d", len(a), cap(a))
	t.Logf("len=%d,cap=%d", len(b), cap(b))
}

func TestAppend(t *testing.T) {
	a := []int{1, 2, 3}
	t.Log(a)
	t.Logf("len=%d,cap=%d", len(a), cap(a))
	//这里，使用一个a来承接append之后的切片，因为可能地址改变
	b := append(a, 4)
	t.Log(a)
	t.Logf("len=%d,cap=%d", len(a), cap(a))
	t.Log(b)
	t.Logf("len=%d,cap=%d", len(b), cap(b))

	//slice can only be compared to nil
	//t.Log(a == b) X
	c := append(a, 5)
	t.Log(a) //[1 2 3]
	t.Logf("len=%d,cap=%d", len(a), cap(a))
	t.Log(b) //[1 2 3 4]
	t.Logf("len=%d,cap=%d", len(b), cap(b))
	t.Log(c) //[1 2 3 5]
	t.Logf("len=%d,cap=%d", len(c), cap(c))

	d := append(b, 5)
	t.Log(d) //[1 2 3 4 5]
	t.Logf("len=%d,cap=%d", len(d), cap(d))
	e := append(b, 6)
	t.Log(d) //[1 2 3 4 6] 我的5呢？
	t.Logf("len=%d,cap=%d", len(d), cap(d))
	t.Log(e) //[1 2 3 4 6]
	t.Logf("len=%d,cap=%d", len(e), cap(e))

}

func TestSliceTranverse(t *testing.T) {
	a := []int{1, 2, 3, 4}
	d := append(a, 5)
	d = append(d, 6)
	for index, item := range a {
		t.Logf("index=%d,item=%d", index, item)
	}
	t.Log()
	for index, item := range d {
		t.Logf("index=%d,item=%d", index, item)
	}
	t.Log()
	e := d[3:6]
	for index, item := range e {
		t.Logf("index=%d,item=%d", index, item)
	}
}
