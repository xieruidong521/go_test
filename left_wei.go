package main

import "fmt"

type Byte float64

const (
	_=iota
	KB Byte =1<<(10*iota)
	MB
	GB
	TB
)

func main() {
	fmt.Println(1|1)
}