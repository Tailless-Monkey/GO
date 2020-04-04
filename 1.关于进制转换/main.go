package main

import "fmt"

// go语言中无法定义二进制的数
func main() {
	// 定义一个十进制数
	var a = 10
	fmt.Printf("%d \n", a)
	// 十进制转换为二进制，用%b
	fmt.Printf("%b \n", a)
	// 十进制转换为八进制,用%o
	fmt.Printf("%o \n", a)
	// 十进制转换为十六进制，用%x
	fmt.Printf("%x \n", a)

	// 定义一个八进制数，0开头的为八进制数
	b := 077
	fmt.Printf("%d \n", b)

	// 定义一个十六进制，0x开头
	c := 0x12
	fmt.Printf("%d \n", c)
}
