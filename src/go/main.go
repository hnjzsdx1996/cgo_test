package main

import "fmt"

func main() {
	fmt.Println("Go start")

	// 创建HelloLib实例
	lib := NewHelloLib()

	// 调用基本方法
	fmt.Println("调用基本方法:")
	lib.PrintHelloWorld()

	fmt.Println("Go程序运行完成！")
}
