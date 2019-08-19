package main

import "fmt"
import (
	"container/list"
)

func main() {
	make2 :=list.New()
	for i := 0; i<10; i++ {
		make2.PushBack(i)
	}
	make2.MoveBefore(make2.Back(),make2.Front())
	for d:=make2.Front();d!=nil ; d=d.Next() {
		fmt.Println(d.Value)
	}
	fmt.Println(make2)
	fmt.Println(make2.Len())
}
