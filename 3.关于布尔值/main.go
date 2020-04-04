package main

import "fmt"

func main() {
	// 布尔值类型变量的默认值是false，布尔值不能参与数值运算
	b1 := true
	var b2 bool
	fmt.Printf("%T \n", b1)
	// %v这个占位符表示打印变量的值
	fmt.Printf("%T value:%v \n", b2, b2)
}
