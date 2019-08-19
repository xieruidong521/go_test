package main

import (
	"fmt"
	"time"
)

func main(){
	ch:=make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch<-1
		fmt.Println("run routine")
	}()

	//time.Sleep(time.Second)

	boo:=true
	for boo{
		select {
		case <-ch:
			fmt.Println("read ch success")
			boo=false
		case <-time.After(time.Second):
			fmt.Println("oh timeout!")
		}
	}
	fmt.Println("finished")
}
