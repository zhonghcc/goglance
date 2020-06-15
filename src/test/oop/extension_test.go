package oop

import (
	"fmt"
	"testing"
)

type Pet struct {
	color int
}

func (*Pet) speak() {
	fmt.Println("...")
}

type Dog struct {
}

func (*Dog) speak() {
	fmt.Println("wang wang")
}

func TestExtension(t *testing.T) {
	var dog *Dog = new(Dog)
	dog.speak()
}
