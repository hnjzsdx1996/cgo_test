package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Go start")

	// 创建HelloLib实例
	lib := NewHelloLib()

	// 调用基本方法
	fmt.Println("调用基本方法:")
	lib.PrintHelloWorld()
	fmt.Println()

	fmt.Println("设置全局配置回调:")
	lib.SetGlobalConfigCallback(func(config SdkConfig) {
		fmt.Println("全局配置回调:")
		b, _ := json.MarshalIndent(config, "", "  ")
		fmt.Println(string(b))
		fmt.Println()
	})

	fmt.Println("参数传递:")
	lib.PrintAny("hello any")
	fmt.Println()

	fmt.Println("参数传递:")
	lib.PassSdkConfig(SdkConfig{
		LogLevel: 1,
		LogPath:  "log_path",
		Url:      "url",
	})
	fmt.Println()

	fmt.Println("获取全局配置:")
	b, _ := json.MarshalIndent(lib.GetGlobalConfig(), "", "  ")
	fmt.Println(string(b))
	fmt.Println()

	fmt.Println("加法:")
	fmt.Println(lib.Add(1, 2))
	fmt.Println()

	fmt.Println("获取全局配置:")
	b, _ = json.MarshalIndent(lib.GetGlobalConfig(), "", "  ")
	fmt.Println(string(b))
	fmt.Println()

	fmt.Println("Go程序运行完成！")
}
