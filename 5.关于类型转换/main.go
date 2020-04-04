package main

import "fmt"

func main() {
	// 类型转换，将整型强制转换为浮点类型
	n := 10
	var f float64
	f = float64(n)
	fmt.Println(f)
	fmt.Println("%T \n", f)

}
