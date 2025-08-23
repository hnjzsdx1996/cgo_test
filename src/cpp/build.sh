#!/bin/bash

# 创建构建目录
mkdir -p build
cd build

# 配置cmake
cmake ..

# 编译
make

# 运行程序
echo "运行程序："
./bin/hello_main

echo ""
echo "构建完成！"
