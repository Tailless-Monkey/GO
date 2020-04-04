package main

import (
	"fmt"
	"strings"
)

func main() {
	// go语言中的字符串只能使用双引号表示，单引号包裹的是字符
	fmt.Println("\"Hello GoLang\",I said")

	// 多行字符串，使用反引号包裹
	s1 := `
	十年生死两茫茫
	不思量
	自难忘
	千里孤坟无处话凄凉
	`
	fmt.Println(s1)

	// 答应字符串的长度，取的是byte字节的数量
	fmt.Println(len(s1))

	// 遍历打印字符串中每一个字符
	for _, c := range s1 {
		// 默认打印每一个字符的byte
		fmt.Println(c)
		// 将字节转换为字符
		fmt.Printf("%c \n", c)
	}

	// 字符串的拼接，也可以使用strings包中的join拼接
	name := "理想,"
	world := "很骨感"
	fmt.Println(name + world)
	// 通过占位符拼接
	fmt.Printf("%s%s", name, world)
	// 和Println不同的是Sprintln不会直接将内容打印出来，而是可以传递给一个新的变量名
	ss := fmt.Sprintf("%s%s", name, world)

	// 字符串的分割，通过strings这个包进行操作
	ret := strings.Split(ss, ",")
	fmt.Println(ret)

	// 判断字符串中的内容是否包含，有的话返回true，没有返回flase
	fmt.Println(strings.Contains(ss, "骨感"))

	// 判断字符串的开头和结尾是否包含所需内容，开头为HasPrefix，结尾为HasSuffix
	fmt.Println(strings.HasSuffix(ss, "骨感"))

	// 查看索引
	test := "abcdea"
	fmt.Println(strings.Index(test, "c"))

	// 字符串修改，字符串类型默认是不能修改的，必须转为rune类型，rune类型就是int32类型
	n1 := "黑色"
	// 将字符串强制转换为一个rune切片
	n2 := []rune(n1)
	// 修改的为字符，必须用单引号
	n2[0] = '蓝'
	// 把rune切片强制转为字符串
	fmt.Println(string(n2))
}
