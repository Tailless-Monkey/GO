package main

import "fmt"

func main() {
	// 算术运算符
	var (
		a = 6
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// a++ b-- 都是单独的语句，不能被赋值，a = a++是错误的表达式

	// 关系运算符，go是强类型，相同类型的变量才能比较
	fmt.Println(a == b)
	fmt.Println(a != b)

	// 逻辑运算符
	age := 22
	if age > 18 && age < 30 {
		fmt.Println("年轻人")
	} else {
		fmt.Println("不是年轻人")
	}

	// not取反
	isMarried := false
	fmt.Println(!isMarried)

	// 位运算符，针对的为二进制数
	// &按位与,参与运算的两个数两位均为1为真，|两位或，两位有一个为1就是1
	fmt.Println(5 & 2)
}
