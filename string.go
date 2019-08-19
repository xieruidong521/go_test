package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello")
	fmt.Println(strings.Replace("abcdea", "a", "-", -1))
}
