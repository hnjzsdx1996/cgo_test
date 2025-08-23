# Go调用C++库示例

这个项目演示了如何使用Go语言调用C++库。

## 项目结构

```
src/
├── cpp/           # C++库源码
│   ├── hello.h    # 头文件
│   ├── hello.cpp  # 实现文件
│   ├── main.cpp   # C++测试程序
│   └── CMakeLists.txt
└── go/            # Go程序
    ├── main.go    # Go主程序
    ├── go.mod     # Go模块文件
    └── build_and_run.sh  # 构建和运行脚本
```

## 使用方法

### 方法1：使用构建脚本（推荐）

```bash
cd src/go
./build_and_run.sh
```

这个脚本会自动：
1. 编译C++库
2. 运行Go程序

### 方法2：手动构建

1. 首先编译C++库：
```bash
cd src/cpp
mkdir -p build
cd build
cmake ..
make
```

2. 然后运行Go程序：
```bash
cd ../../go
go run main.go
```

## 技术要点

- 使用`extern "C"`将C++函数包装成C函数
- 使用CGO的`#cgo`指令指定编译和链接参数
- 设置正确的库路径和头文件路径
- 链接C++标准库（`-lstdc++`）

## 注意事项

- 确保C++库已经编译完成
- 如果遇到链接错误，检查库文件路径是否正确
- 在Linux系统上，可能需要设置`LD_LIBRARY_PATH`环境变量
