package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main(){
	var redis goredis.Client
	redis.Addr="127.0.0.1:6379"
	setErr:=redis.Set("user",[]byte("Lucy"))
	if setErr!=nil{
		panic(setErr)
	}
	fmt.Println("写入值：Lucy")
	get,getErr:=redis.Get("user")
	if getErr!=nil{
		panic(getErr)
	}
	fmt.Println("获取到值"+string(get))

	user:=make(map[string]interface{})
	user["name"]="lucy"
	user["age"] =20
	user["sex"] ="male"
	user["addr"]="临沂市兰山区"
	hmErr:=redis.Hmset("user_hash",user)
	if hmErr!=nil{
		panic(hmErr)
	}
	_, _ = redis.Hset("name", "_name", []byte("xiaoli"))
}