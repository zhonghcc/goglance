package function

import (
	"fmt"
	"testing"
)

func doSomething(a int, b int) int {
	fmt.Printf("input a=%d,b=%d\n", a, b)
	ret := a + b
	fmt.Printf("output ret=%d\n", ret)
	return ret
}

func aop(f func(int, int) int) func(int, int) int {
	generate := func(a int, b int) int {
		fmt.Println("do something before")
		ret := f(a, b)
		fmt.Println("do something after")
		return ret
	}
	return generate
}
func TestFirstClass(t *testing.T) {
	arr := [...]func(int, int) int{doSomething, doSomething}
	t.Log(arr)
	t.Log(arr[0](1, 2))
}

func TestAop(t *testing.T) {
	f := aop(doSomething)
	t.Log(f(1, 2))
}

func TestLambda(t *testing.T) {
	a := func(a int, b int) int {
		return a * b
	}
	t.Log(a(2, 4))

}

func TestClosure(t *testing.T) {
	c := 4
	f := func(a int, b int) int {
		c = 2
		return a * b * c
	}
	t.Log(f(2, 4))
	t.Log(c)
}
