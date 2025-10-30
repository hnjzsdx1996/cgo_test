package main

/*
#cgo CFLAGS: -I../cpp
#cgo LDFLAGS: -L../cpp/build -lhello_lib -lstdc++
#include "hello.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

// HelloLib 定义了一个接口，用于包装C++库的功能
type HelloLib interface {
	// PrintHelloWorld 输出helloworld
	PrintHelloWorld()

	// PrintAny 输出任意字符串
	PrintAny(string)

	// PassSdkConfig 传递SDK配置
	PassSdkConfig(SdkConfig)
}

// cppHelloLib 实现了HelloLib接口，内部调用C++库
type cppHelloLib struct{}

// PrintHelloWorld 实现HelloLib接口，调用C++库中的printHelloWorld函数
func (h *cppHelloLib) PrintHelloWorld() {
	C.printHelloWorld()
}

func (h *cppHelloLib) PrintAny(s string) {
	cs := C.CString(s)
	C.printAny(cs)
	C.free(unsafe.Pointer(cs))
}

// NewHelloLib 创建一个新的HelloLib实例
func NewHelloLib() HelloLib {
	return &cppHelloLib{}
}

func (h *cppHelloLib) PassSdkConfig(config SdkConfig) {
	cLogPath := C.CString(config.LogPath)
	cUrl := C.CString(config.Url)
	cConfig := C.struct_SdkConfig{
		log_level: C.int(config.LogLevel),
		log_path:  cLogPath,
		url:       cUrl,
	}
	C.passSdkConfig(cConfig)
	C.free(unsafe.Pointer(cLogPath))
	C.free(unsafe.Pointer(cUrl))
}

type SdkConfig struct {
	LogLevel int
	LogPath  string
	Url      string
}
