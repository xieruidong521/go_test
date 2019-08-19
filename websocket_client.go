package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
func main() {
	dial,err:=net.Dial("tcp","127.0.0.1:8080")

	if err!=nil{
		fmt.Println("连接出错："+err.Error())
		os.Exit(1)
	}

	defer dial.Close()

	go sendMsg(dial)

	for {
		bytes:=make([]byte,1024)
		msgNum,err:=dial.Read(bytes)
		if err!=nil{
			fmt.Println("接收消息出错："+err.Error())
			os.Exit(0)
		}
		fmt.Println("接受到消息："+string(bytes[:msgNum]))
	}
}
func sendMsg(conn net.Conn){
	//bytes:=make(net.Buffers,1024)

	for {
		reader:=bufio.NewReader(os.Stdin)
		input,_,_:=reader.ReadLine()
		if strings.ToUpper(string(input)) == "EXIT" {
			conn.Close()
			fmt.Println("即将退出")
			break
		}
		_,err:=conn.Write([]byte(input))
		if err!=nil{
			fmt.Println("发送消息出错："+err.Error())
			continue
		}
	}
}