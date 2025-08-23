package main

/*
#cgo CFLAGS: -I../cpp
#cgo LDFLAGS: -L../cpp/build -lhello_lib -lstdc++
#include "hello.h"
*/
import "C"

import (
	"fmt"
	"time"
)

// HelloLib 定义了一个接口，用于包装C++库的功能
type HelloLib interface {
	// PrintHelloWorld 输出helloworld
	PrintHelloWorld()
	// PrintHelloWorldWithTimestamp 带时间戳输出helloworld
	PrintHelloWorldWithTimestamp()
	// GetLibraryInfo 获取库信息
	GetLibraryInfo() string
}

// cppHelloLib 实现了HelloLib接口，内部调用C++库
type cppHelloLib struct{}

// PrintHelloWorld 实现HelloLib接口，调用C++库中的printHelloWorld函数
func (h *cppHelloLib) PrintHelloWorld() {
	C.printHelloWorld()
}

// PrintHelloWorldWithTimestamp 带时间戳输出helloworld
func (h *cppHelloLib) PrintHelloWorldWithTimestamp() {
	fmt.Printf("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	C.printHelloWorld()
}

// GetLibraryInfo 获取库信息
func (h *cppHelloLib) GetLibraryInfo() string {
	return "C++ Hello Library v1.0 (wrapped by Go)"
}

// NewHelloLib 创建一个新的HelloLib实例
func NewHelloLib() HelloLib {
	return &cppHelloLib{}
}
