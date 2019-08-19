package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	page,err:=http.Post("http://mine.a.com/api/test","application/x-www-form-urlencoded",strings.NewReader("id=1&make=2"))
	if err!=nil {
		panic(err)
	}
	defer page.Body.Close()
	html,err:=ioutil.ReadAll(page.Body)
	realHtml:=string(html)
	fmt.Println(realHtml)
	var obj interface{}
	err1:=json.Unmarshal([]byte(realHtml),&obj)
	if err1!=nil{
		panic(err1)
	}
	fmt.Println(obj)
}
