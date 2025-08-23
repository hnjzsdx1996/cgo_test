package main

/*
#cgo CFLAGS: -I../cpp
#cgo LDFLAGS: -L../cpp/build -lhello_lib -lstdc++
#include "hello.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Println("Go程序开始运行...")

	// 调用C++库中的函数
	C.printHelloWorld()

	fmt.Println("Go程序运行完成!")
}
