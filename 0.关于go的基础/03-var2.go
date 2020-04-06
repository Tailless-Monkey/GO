package main

import (
	"fmt"
)

func main() {
	//变量的初始化，在声明变量的同时进行赋值,和直接赋值没有区别
	var num int = 10
	fmt.Println("num is:", num)

	//赋值前必须声明变量类型，否则会报错。不过赋值接受多个变量名值叠加
	var num1 int
	num1 = 20
	num1 = 24
	fmt.Println("num1 type is %T", num1)

	//自动推导类型，必须先初始化，通过初始化的值确定类型，然后再赋值给变量。
	num2 := 30
	//使用:=推导的变量只能出现一次，不能叠加使用,这个是自动推导类型和赋值的最大区别
	//num2 := 39
	fmt.Printf("num2 type is %T\n", num2)
}
