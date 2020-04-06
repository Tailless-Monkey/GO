package main

import "fmt"

func main() {
	// 数组，存放元素的容器，必须指定存放元素的类型和容量，数组的长度也是数组类型的一部分
	var a1 [3]bool //[true  false  true]
	var a2 [4]bool

	fmt.Printf("a1:%T  a2:%T \n", a1, a2)

	// 数组初始化，不初始化默认情况下元素都是零值
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 根据初始值自动推断数组的长度
	a100 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a100)

	// 根据索引初始化数组，其他没有索引的为零值
	a3 := [4]int{0: 1, 1: 2}
	fmt.Println(a3)

	// 数组的遍历
	citys := [...]string{"北京", "上海", "常州", "苏州", "南京", "无锡"}
	// 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组, 比如[[1,2][3,4}[5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
