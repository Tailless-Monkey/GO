package main

import (
	"fmt"
)

func main() {
	//声明变量，类型bool,如果没有初始化，初始值为false
	var a bool
	a = true
	fmt.Println("a = ", a)

	//自动推导类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)
}
