package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("the number is ",i)
		}(i)
	}
	time.Sleep(time.Second*10)
}
