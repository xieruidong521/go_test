package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		code,_:=writer.Write([]byte("你好"))
		fmt.Println("success",code)
	})
	http.ListenAndServe("127.0.0.1:8888",nil)
}
