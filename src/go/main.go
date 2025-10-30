package main

import "fmt"

func main() {
	fmt.Println("Go start")

	// 创建HelloLib实例
	lib := NewHelloLib()

	// 调用基本方法
	fmt.Println("调用基本方法:")
	lib.PrintHelloWorld()
	fmt.Println()

	fmt.Println("参数传递:")
	lib.PrintAny("hello any")
	fmt.Println()

	fmt.Println("Go程序运行完成！")
}
