package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Hello,World")
	// 变量的声明
	var x int
	var s string
	var b bool
	// 随机生成数据
	var num = rand.Intn(12)
	// 变量打印使用 基本数据类型当变量只声明未赋值，默认为零值
	fmt.Println(x)
	fmt.Println(s)
	fmt.Println(b)
	if _, err := fmt.Println(num); err != nil {
		return
	}

}
