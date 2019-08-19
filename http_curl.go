package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var url string="https://movie.douban.com/subject/24751763/"
func main(){
	data,_:=http.Get(url)
	defer data.Body.Close()
	html,_:=ioutil.ReadAll(data.Body)
	htmlLast:=string(html)

	regName:=regexp.MustCompile("<span\\s+property=\"v:itemreviewed\">(.*)</span>")
	regScore:=regexp.MustCompile("<strong\\s+class=\"ll\\s+rating_num\"\\s+property=\"v:average\">(.*)</strong>")
	nameRes:=regName.FindAllStringSubmatch(htmlLast,-1)
	scoreRes:=regScore.FindAllStringSubmatch(htmlLast,-1)
	fmt.Println(nameRes[0][1])
	fmt.Println(scoreRes[0][1])
}