package oop

import (
	"fmt"
	"testing"
)

type Duck struct {
	color int
	age   int
}

type Gaga interface {
	gaga()
}

func (d *Duck) tellAge() {
	fmt.Printf("im %d years old\n", d.age)
}

func (d *Duck) gaga() {
	fmt.Printf("gaga\n")
}
func TestDuckAge(t *testing.T) {
	duck := new(Duck)
	duck.age = 2
	duck.tellAge()
	var gaga Gaga
	gaga = duck
	gaga.gaga()
}
