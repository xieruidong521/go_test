package pkg

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	fmt.Println("正在初始化变量Pi...")
	Pi=math.Pi
}