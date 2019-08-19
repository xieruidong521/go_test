package main

import (
	"fmt"
	"strings"
)

func main() {
	str:="aa                                     "
	str=strings.ReplaceAll(str," ","")
	fmt.Println(str)
}

type f func()