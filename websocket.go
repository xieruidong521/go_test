package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	LogDir    ="D:/"
	LogServer ="server.log"
	Listen    ="0.0.0.0:8080"
)

//存储websocket连接
var connections=make(map[string]net.Conn)
var msgChan=make(chan string)
var logFile *os.File
var logger *log.Logger

func main() {
	//处理日志
	logFile,err:=os.OpenFile(LogDir+LogServer,os.O_WRONLY|os.O_CREATE,0)
	if err!=nil{
		fmt.Println("打开日志文件失败："+err.Error())
		os.Exit(999)
	}
	logger=log.New(logFile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("创建log文件完成...")

	ws,err:=net.Listen("tcp",Listen)
	if err != nil {
		panic(err)
	}
	//close
	defer ws.Close()

	//处理消息
	go forwardMessage()

	fmt.Println("正在监听...")
	logger.Println("正在监听...")

	for  {
		conn,err:=ws.Accept()
		if err != nil {
			fmt.Println("连接出错："+err.Error())
		}

		addr:=conn.RemoteAddr().String()
		if _, ok := connections[addr];ok {
			continue
		}
		connections[addr]=conn

		go receive(conn)
	}
}

func receive(conn net.Conn) {
	bytes:=make([]byte,1024)
	addr:=conn.RemoteAddr()

	defer func(conn net.Conn,connections map[string]net.Conn) {
		addr:=conn.RemoteAddr().String()
		delete(connections,addr)
		err:=conn.Close()
		if err!=nil{
			fmt.Println("关闭"+addr+"出错："+err.Error())
		}
		fmt.Println("用户"+addr+"已关闭连接")
		logger.Println("用户"+addr+"已关闭连接")
		fmt.Println(connections)
	}(conn,connections)

	//连接提示
	fmt.Println("用户"+addr.Network()+":"+addr.String()+"已连接")
	logger.Println("用户"+addr.Network()+":"+addr.String()+"已连接")
	fmt.Println(connections)

	for {
		byteNum,err:=conn.Read(bytes)
		if err!=nil{
			fmt.Println("连接断开："+err.Error())
			break
		}
		msg:=string(bytes[:byteNum])
		if strings.ToLower(msg)=="list"{
			var ips string = "===============\n"
			for addr,_:=range connections{
				ips+=addr+"\n"
			}
			ips+="===============\n"
			_, _ = conn.Write([]byte(ips))
			continue
		}
		msgChan<-msg
	}
}

func forwardMessage(){
	for{
		select {
		case msg := <-msgChan:
			msgInfo:=strings.Split(msg,"#")
			if len(msgInfo)<2{
				fmt.Println("未找到转发地址，忽略此消息...")
				continue
			}
			addr:=msgInfo[0]
			msgSend:=strings.Join(msgInfo[1:],"#")
			if conn,ok:=connections[addr];ok{
				_, _ = conn.Write([]byte(msgSend))
			}
		}
	}
}