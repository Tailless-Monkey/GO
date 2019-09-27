package main

import (
	"fmt"
)

func main() {
	//变量的声明格式为 var  变量名 类型
	//变量被声明，必须要使用
	//在同一个{}里，声明的变量名必须唯一
	//可以同时声明多个变量
	var num1, num2 int
	fmt.Println("num1 is:", num1)
	fmt.Println("num2 is:", num2)

	//为变量赋值
	num2 = 100
	fmt.Println("num2 is:", num2)
}
