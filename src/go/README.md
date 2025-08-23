# Go调用C++库示例

这个项目演示了如何使用Go语言调用C++库，并通过Go接口进行包装，提供更好的IDE支持和测试能力。

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
    ├── hello_wrapper.go  # C++库的Go接口包装
    ├── hello_wrapper_test.go  # 接口测试
    ├── go.mod     # Go模块文件
    └── build_and_run.sh  # 构建和运行脚本
```

## 接口设计优势

### 1. **IDE支持**
- 可以自动跳转到Go接口方法
- 提供完整的代码补全
- 支持重构和重命名

### 2. **测试友好**
- 可以创建Mock实现进行单元测试
- 不依赖C++库就能测试Go逻辑
- 支持接口测试

### 3. **扩展性**
- 可以在Go层面添加新功能
- 支持多种实现（真实C++库、Mock等）
- 便于添加日志、监控等功能

## 使用方法

### 方法1：使用构建脚本（推荐）

```bash
cd src/go
./build_and_run.sh
```

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
go run .
```

3. 运行测试：
```bash
go test
```

## 接口方法

```go
type HelloLib interface {
    PrintHelloWorld()                    // 基本输出
    PrintHelloWorldWithTimestamp()       // 带时间戳输出
    GetLibraryInfo() string              // 获取库信息
}
```

## 技术要点

- **接口分离**：将CGO代码和Go接口分离到不同文件
- **依赖注入**：通过接口实现依赖注入模式
- **测试驱动**：支持Mock测试，提高测试覆盖率
- **错误处理**：在Go层面添加错误处理和日志记录

## 注意事项

- 确保C++库已经编译完成
- CGO代码必须在import "C"之前
- 接口设计遵循Go语言的最佳实践
- 测试时可以使用Mock实现避免依赖C++库
