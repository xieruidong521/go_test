package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){
	ch:=make(chan int)
	go func() {
		for a:=1;a<=10;a++{
			if a == 5 {
				<-ch
			}
			fmt.Println("routine 1,number:"+strconv.Itoa(a))
		}
	}()
	go func() {
		for i:=11;i<=20;i++{
			fmt.Println("routine 2,number:"+strconv.Itoa(i))
		}
		ch<-1
	}()
	time.Sleep(time.Second)
	fmt.Println("finished")
}