package pkg1

import (
	"fmt"
	v2 "goglance/src/test/variables/pkg2"
)

func Add() {

	fmt.Println(v2.Pkg2Int)
	v2.Pkg2Int = v2.Pkg2Int + 1
}
