package main

/*
#cgo CFLAGS: -I../cpp
#cgo LDFLAGS: -L../cpp/build -lhello_lib -lstdc++
#include "hello.h"
#include <stdlib.h>

// 声明一个回调类型，便于从 Go 传入 C 的函数指针
typedef void (*SdkConfigCallback)(struct SdkConfig);

// 声明一个由 Go 实现、供 C 调用的回调入口
void callbackGo(struct SdkConfig cfg);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// HelloLib 定义了一个接口，用于包装C++库的功能
type HelloLib interface {
	// PrintHelloWorld 输出helloworld
	PrintHelloWorld()

	// PrintAny 输出任意字符串
	PrintAny(string)

	// SetGlobalConfigCallback 设置在 C++ 更新配置时的回调
	SetGlobalConfigCallback(func(SdkConfig))

	// PassSdkConfig 传递SDK配置
	PassSdkConfig(SdkConfig)

	// GetGlobalConfig 获取全局配置
	GetGlobalConfig() SdkConfig

	// Add 加法
	Add(int, int) int
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

func (h *cppHelloLib) Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

func (h *cppHelloLib) GetGlobalConfig() SdkConfig {
	var cConfig C.struct_SdkConfig
	C.getGlobalConfig((*C.struct_SdkConfig)(unsafe.Pointer(&cConfig)))
	return SdkConfig{
		LogLevel: int(cConfig.log_level),
		LogPath:  C.GoString(cConfig.log_path),
		Url:      C.GoString(cConfig.url),
	}
}

var goConfigCallback func(SdkConfig)

//export callbackGo
func callbackGo(c C.struct_SdkConfig) {
	fmt.Println("callbackGo called")
	if goConfigCallback != nil {
		goConfigCallback(SdkConfig{
			LogLevel: int(c.log_level),
			LogPath:  C.GoString(c.log_path),
			Url:      C.GoString(c.url),
		})
	}
}

func (h *cppHelloLib) SetGlobalConfigCallback(cb func(SdkConfig)) {
	goConfigCallback = cb
	C.setGlobalConfigCallback((C.SdkConfigCallback)(C.callbackGo))
}
