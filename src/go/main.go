package main

import "fmt"

func main() {
	fmt.Println("Go start")

	// 创建HelloLib实例
	lib := NewHelloLib()

	// 显示库信息
	fmt.Printf("库信息: %s\n", lib.GetLibraryInfo())

	// 调用基本方法
	fmt.Println("\n调用基本方法:")
	lib.PrintHelloWorld()

	// 调用带时间戳的方法
	fmt.Println("\n调用带时间戳的方法:")
	lib.PrintHelloWorldWithTimestamp()

	fmt.Println("\nGo程序运行完成！")
}
