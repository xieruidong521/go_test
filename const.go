package main

import "fmt"
import "runtime"

const fl = 0.36564168546516 / 0.2564646431348545641235

func main() {
	var a *int
	var b string = "你好"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(runtime.GOOS)
}
