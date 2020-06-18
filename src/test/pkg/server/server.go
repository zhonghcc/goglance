package server

import "fmt"

func init() {
	fmt.Println("im init 1")
}

func init() {
	fmt.Println("im init 2")
}

func doSomethingPrivate() {
	fmt.Println("you cant see me")
}

func GetIntSqure(i int) (int, error) {
	return i * i, nil
}
