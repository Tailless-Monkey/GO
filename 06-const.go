package main

import (
	"fmt"
)

func main() {
	//变量是程序运行期间可以被随时改变的量，常量在程序运行时是不能变化的
	//常量的声明使用const

	const num int = 10
	//常量不允许再被赋值修改，执行会报错
	//num = 20

	//常量的自动推导，不需要使用:=
	const num1 = 20.0

	fmt.Printf("num的值是%d\n", num)
	fmt.Println("num1 = ", num1)
	fmt.Printf("num1的类型是 %T\n", num1)
}
