package array

func TestInit(t *testing.T){
	a := [...]int{1,2,3}
	b := [3]int
	t.Log(a)
	t.Log(b)
}