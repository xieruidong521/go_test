package main

import (
	"fmt"
	"regexp"
)

func main() {
	//exp,_:=regexp.MatchString("^1[2|3]{1}\\d{9}$","13500000000")
	r,_:=regexp.Compile("1[2|3]{1}\\d{9}")
	fmt.Println(r.FindAllString("12345678901 12345678902",-1))
}