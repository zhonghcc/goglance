package main

import (
	"fmt"
	"testing"
)

var (
	i int = 1
)

func TestVariable(t *testing.T) {
	fmt.Println(i)
	t.Log(i)
}
