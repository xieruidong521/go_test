package main

import (
	"fmt"
	"math"
)

type (
	Int int
	Fl float64
)

func main() {
	var a Int
	var b bool
	var arr [10]rune
	fmt.Println(a,b,arr)
	fmt.Println(math.MaxUint8)
}