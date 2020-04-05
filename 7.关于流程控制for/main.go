package main

import "fmt"

func main() {
	// for循环基本表达式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 使用方式1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 使用方式2
	var j = 5
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// 无限循环,慎用，会卡死
	//for {
	//	fmt.Println("a")
	//}

	// for range循环
	s := "你好，世界"
	for i, v := range s {
		fmt.Printf("%d %c \n", i, v)
	}

	// 跳出for循环，当i为5时跳出
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 当i=5时，跳过当前循环（不执行当前循环语句）
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}