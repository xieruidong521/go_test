package main

import (
	"fmt"
	"strconv"
)

func main() {
	a:=100
	b:=string(a)
	c:="101"
	fmt.Println(b)
	fmt.Println(strconv.Atoi(c))
	fmt.Println(fmt.Sprintf("%d",a))
}
