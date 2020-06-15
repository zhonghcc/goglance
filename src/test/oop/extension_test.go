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
func (p *Pet) getColor() int {
	return p.color
}

type Dog struct {
}

func (*Dog) speak() {
	fmt.Println("wang wang")
}

func TestExtension(t *testing.T) {
	var dog *Dog = new(Dog)
	dog.speak()
	var pet *Pet = new(Pet)
	pet.speak()
	// var pet2 *Pet = dog
	// pet2.speak()

}

type DogWithParent struct {
	*Pet
}

func (this *DogWithParent) speak() {
	this.Pet.speak()
}

func TestExtensionParent(t *testing.T) {
	var dog *DogWithParent = new(DogWithParent)
	dog.speak()
	// t.Log(dog.getColor()) 因为没有分配内存，所以这个用法是不行的
}

type DogWithStruct struct {
	Pet
}

//重写了speak，如果不重写，实际会调用Pet对象的speak
func (this *DogWithStruct) speak() {
	fmt.Println("wangwangwang")
}
func TestExtensionStruct(t *testing.T) {
	var dog *DogWithStruct = new(DogWithStruct)
	dog.speak()
	dog.color = 4
	t.Log(dog.getColor())
	t.Log(dog.Pet.getColor())
}
