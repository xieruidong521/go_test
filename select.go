package main

import (
	"fmt"
	"time"
)

func main(){
	ch:=make(chan int)
	timeout:=make(chan bool,1)

	go func() {
		time.Sleep(1)
		timeout<-true
	}()

	select {
	case <-ch:
		fmt.Println("read ch success")
	case <-timeout:
		fmt.Println("oh timeout!")
	}
	fmt.Println("finished")
}
