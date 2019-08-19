package main

var b = "G"

func main() {
	nn()
	mm()
	nn()
}

func nn() {
	print(b)
}
func mm()  {
	b="O"
	print(b)
}