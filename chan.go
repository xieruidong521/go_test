package main

import (
	"fmt"
	"time"
)

var ch chan int

func routine(){
	ch<-1
	fmt.Println("routine 1")
	ch<-2
	fmt.Println("routine 2")
	ch<-3
	fmt.Println("i'm routine 3")
}
func main(){
	ch=make(chan int,2)
	go routine()
	time.Sleep(time.Second)
	fmt.Println("finished")
	<-ch
	time.Sleep(10)
}