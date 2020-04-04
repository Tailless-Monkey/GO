package main

import (
	"fmt"
)

func main() {
	// 默认go语言中的浮点数都是float64类型
	f1 := 1.23456
	fmt.Printf("%T \n", f1)

	// 强制将f1的类型转为float32，不能直接赋值，因为是不同类型的，一个32位，一个64位
	f2 := float32(f1)
	fmt.Println(f2)
	fmt.Printf("%T \n", f2)
}
 