package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
	age int
}

func main(){
	arr:=[5]int{1,2,3,4,5}
	mp:=make(map[string]string)
	mp["name"]="晓丽"
	arrJson,_:=json.Marshal(arr)
	mpJson,_:=json.Marshal(mp)
	fmt.Println(string(arrJson))
	fmt.Println(string(mpJson))
	var decode interface{}
	err:=json.Unmarshal([]byte(mpJson),&decode)
	err=json.Unmarshal([]byte(arrJson),&decode)
	fmt.Println(err,decode)

	md5Str:=[]byte("你好")
	md5Last:=md5.Sum(md5Str)
	fmt.Printf("%x",md5Last)
}