package main

import (
	"strings"
	"testing"
)

// mockHelloLib 是一个模拟实现，用于测试
type mockHelloLib struct{}

func (m *mockHelloLib) PrintHelloWorld() {
	// 模拟实现，不调用C++库
}

func (m *mockHelloLib) PrintHelloWorldWithTimestamp() {
	// 模拟实现，不调用C++库
}

func (m *mockHelloLib) GetLibraryInfo() string {
	return "Mock Hello Library v1.0 (for testing)"
}

// TestHelloLibInterface 测试HelloLib接口
func TestHelloLibInterface(t *testing.T) {
	var lib HelloLib

	// 测试真实实现
	lib = NewHelloLib()
	if lib == nil {
		t.Error("NewHelloLib() returned nil")
	}

	// 测试模拟实现
	lib = &mockHelloLib{}
	if lib == nil {
		t.Error("mockHelloLib is nil")
	}

	// 测试GetLibraryInfo方法
	info := lib.GetLibraryInfo()
	if !strings.Contains(info, "Mock") {
		t.Errorf("Expected mock info, got: %s", info)
	}
}

// TestNewHelloLib 测试NewHelloLib函数
func TestNewHelloLib(t *testing.T) {
	lib := NewHelloLib()
	if lib == nil {
		t.Error("NewHelloLib() returned nil")
	}

	// 验证返回的对象实现了HelloLib接口
	var _ HelloLib = lib
}
