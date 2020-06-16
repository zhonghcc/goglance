package oop

import (
	"fmt"
	"testing"
)

type Obj interface{}

func doSometing(obj Obj) {
	if value, tag := obj.(string); tag == true {
		fmt.Println(value)
	}
}
func doSometingSwitch(obj Obj) {
	switch value := obj.(type) {
	case string:
		fmt.Println(value)
	case int:
		fmt.Println(value)
	default:
		fmt.Println("unknown type")
	}
}

func TestEmptyInterface(t *testing.T) {
	doSometing("halo")
	doSometingSwitch("halo")
	doSometingSwitch(1)
	var a int32 = 1
	doSometingSwitch(a)

}
