package main

import "fmt"

func main() {
	// 切片（slice）是数组类型的一层封装，是一个拥有相同类型元素的可变长度的序列
	var s1 []int    // 定义一个存放int类型元素的切片，没有长度限制
	var s2 []string //定义一个存放string类型元素的切片
	fmt.Println(s1, s2)

	// 切片初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"常州", "无锡", "苏州", "南京"}
	fmt.Println(s1, s2)

	// 切片的长度和容量，切片的长度是底层数组被切割后元素的数量，切片的容量是底层数组的所有元素数量
	fmt.Printf("len(s1): %d cap(s1): %d \n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d cap(s2): %d \n", len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11}
	s3 := a1[0:4]
	fmt.Println(s3)

	// 使用make函数构造切片
	s4 := make([]int, 5, 10)
	fmt.Printf("s4=%v  len(s4)=%d cap(s4)=%d\n", s4, len(s4), cap(s4))

	// 切片的赋值，切片中的内容都指向同一个底层数组
	s5 := []int{1, 3, 5}
	s6 := s5
	fmt.Println(s5, s6)
	s5[0] = 11
	s6[0] = 111
	fmt.Println(s5, s6)

	// 切片的遍历
	// 索引遍历
	for i := 0; i < len(s5); i++ {
		fmt.Println(s5[i])
	}
	// for range循环
	for i, v := range s5 {
		fmt.Println(i, v)
	}

	// 切片的append方法
	str := []string{"钟楼", "天宁", "新北", "武进"}
	fmt.Printf("str  len(str)=%d  cap(str)= %d \n", len(str), cap(str))
	// str[4] = "经开"，这样的写法会导致编译错误，索引越界，因为创建时切片的长度和容量都被限定
	str = append(str, "经开")
	fmt.Println(str)
	// 通过append增加元素后原来的数组长度和容量都会增加
	fmt.Printf("str  len(str)=%d  cap(str)= %d \n", len(str), cap(str))
	newstr := []string{"金坛", "溧阳"}
	str = append(str, newstr...) //...表示拆开前面的变量，不能直接添加一个数组
	fmt.Println(str)

	// copy()函数复制切片
	t1 := []int{1, 3, 5}
	t2 := t1
	t3 := make([]int, 3, 5)
	copy(t3, t1)
	fmt.Println(t1, t2, t3)
	t1[0] = 100
	fmt.Println(t1, t2, t3)

	// 从切片中删除元素，go语言不提供专门的删除切片元素的方法
	t4 := []int{1, 2, 3, 4, 5}
	t4 = append(t4[:1], t4[2:]...) // 删掉数组中的2，从头开始到第一个元素，然后拼接第二个开始到最后
	fmt.Println(t4)
}
