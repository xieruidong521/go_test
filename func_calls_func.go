package main

var str string

func main() {
	str="G"
	print(str)
	fcf1()
}

func fcf1() {
	str:="O"
	print(str)
	fcf2()
}
func fcf2() {
	print(str)
}