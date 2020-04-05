package main

import "fmt"

func main() {
	age := 24

	// 单个条件判断
	if age > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("您还未成年")
	}

	// 多条件判断
	if age > 18 {
		fmt.Println("成年了")
	} else if age > 30 {
		fmt.Println("而立之年")
	} else {
		fmt.Println("您还未成年")
	}
}
