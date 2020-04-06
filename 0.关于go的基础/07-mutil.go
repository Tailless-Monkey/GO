package main

import (
	"fmt"
)

func main() {
	//不同类型变量的声明，可以用多个来定义
	//var a int
	//var b float64

	//多个变量的定义方法
	var (
		a int
		b float64
	)
	a, b = 10, 3.1415926
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//多个常量的定义方法,支持和上面变量一样的简化写法，支持自动推导类型
	const i int = 10
	const j float64 = 3.1415926
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

}
