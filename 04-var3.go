package main

import (
	"fmt"
)

func main() {
	//多个变量同时自动推导
	a, b, c := 10, 20, 30
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	//交换两个变量的值
	a, b = b, a
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
