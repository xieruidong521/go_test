package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

const (
	ListenAddr = "0.0.0.0:8080"
)

//存储用户的connection
var allConnections = make(map[string]*websocket.Conn)
var msgChannel     = make(chan string)

func main(){
	http.HandleFunc("/",ws)
	_ = http.ListenAndServe(ListenAddr, nil)
}
func ws(w http.ResponseWriter,r *http.Request) {
	conn, _ := websocket.Upgrade(w, r, nil, 0, 0)
	handleConn(conn)
	go func(connection *websocket.Conn) {
		defer func() {
			delete(allConnections,connection.RemoteAddr().String())
			connection.Close()
		}()
		for {
			msgType, message, err := connection.ReadMessage()
			fmt.Println("type",msgType)
			if err!=nil {
				fmt.Println(connection.RemoteAddr().String(),"已关闭",err.Error())
				break
			}
			fmt.Println("收到",connection.RemoteAddr().String(),"消息",string(message))
			handleMsg(message)
		}
	}(conn)
}
func handleConn(conn *websocket.Conn) {
	addr:=conn.RemoteAddr().String()
	fmt.Println(addr,"已连接")
	if allConnections[addr]==nil{
		allConnections[addr]=conn
	}
}
func handleMsg(message []byte) {
	msg:=string(message)+"！"
	for _,conn:=range allConnections{
		//fmt.Println(conn.WriteMessage(1,[]byte("1")))
		err := conn.WriteMessage(1, []byte(msg))
		if err != nil {
			fmt.Println("发送消息出错",err.Error())
			continue
		}
	}
}



