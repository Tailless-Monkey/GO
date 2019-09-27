//hello.go，go语言以包作为管理单位，每个文件必须先声明包，main包必须存在
package main

import (
	"fmt"
)

//入口函数main，必须要有，同时大括号不能换行，必须紧跟在小括号后面
/*一个工程只能有一个main函数，在同一个文件夹下出现多个代码文件，不能使用同一个main，所以工程文件最好分开存放
go build  xx.go 编译go脚本，生成可执行程序,不会直接运行
go run xx.go  不生成程序，直接执行
*/
func main() {
	//Println()会自动换行,不支持格式化输出,Printf()不会自动换行，但支持格式化输出。
	//语句结尾不需要分号
	fmt.Println("Hello, World!")
	fmt.Println("Hello, GO!")
}
