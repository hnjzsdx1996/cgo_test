package main

/*
#cgo CFLAGS: -I../cpp
#cgo LDFLAGS: -L../cpp/build -lhello_lib -lstdc++
#include "hello.h"
*/
import "C"

// HelloLib 定义了一个接口，用于包装C++库的功能
type HelloLib interface {
	// PrintHelloWorld 输出helloworld
	PrintHelloWorld()

	// PrintAny 输出任意字符串
	PrintAny(string)
}

// cppHelloLib 实现了HelloLib接口，内部调用C++库
type cppHelloLib struct{}

// PrintHelloWorld 实现HelloLib接口，调用C++库中的printHelloWorld函数
func (h *cppHelloLib) PrintHelloWorld() {
	C.printHelloWorld()
}

func (h *cppHelloLib) PrintAny(s string) {
	C.printAny(C.CString(s))
}

// NewHelloLib 创建一个新的HelloLib实例
func NewHelloLib() HelloLib {
	return &cppHelloLib{}
}
