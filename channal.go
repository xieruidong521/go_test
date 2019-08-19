package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func yes() {
	fmt.Println("im a goroutine")
}
func numbers(){
	for i:=1;i<=5;i++{
		time.Sleep(time.Second*2)
		fmt.Println(i)
	}
}
func abcde(){
	for i:='a';i<='e';i++{
		time.Sleep(time.Second*3)
		fmt.Println(string(i))
	}
}

func main() {
	//go numbers()
	//go abcde()
	//time.Sleep(time.Second*20)
	//fmt.Println("im a main goroutine")
	md:=md5.New()
	md.Write([]byte("123456"))
	mdRes:=md.Sum([]byte(""))
	fmt.Println(fmt.Sprintf("%x",mdRes))
}