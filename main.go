package main

import "fmt"

func Add(x, y int,quit chan int) {
	a:=x+y
	fmt.Println(a)
	quit<-1
}
func main() {
	chs:=make([]chan int,10)
	for i := 0; i < 10; i++ {
		chs[i]=make(chan int)
		go Add(i,i,chs[i])
	}
	for _,v:=range chs{
		<-v
	}
	fmt.Println("finished")
}
