package main

import (
	"fmt"
)

//定义一个带有返回值的函数,并且要求返回多个值，这个函数的目的是为了配合匿名函数使用
func test() (a, b, c int) {
	return 10, 20, 30
}

func main() {
	//匿名变量，丢弃数据不处理。匿名变量配合函数返回值使用，优势更大
	var num1, num2 int
	num1, num2 = 10, 20
	var tmp int
	tmp, _ = num1, num2

	fmt.Printf("tmp的值是：%d\n", tmp)

	//调用前面的函数，获取到abc三个变量的值
	var a, b, c int
	a, b, c = test()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	//通过匿名函数，我们可以只调用返回值的中某几个值
	a, _, c = test()
	fmt.Printf("a = %d, c = %d\n", a, c)
	_, b, _ = test()
	fmt.Printf("b = %d\n", b)

}
