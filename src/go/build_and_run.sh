#!/bin/bash

echo "开始构建C++库..."

# 进入cpp目录构建库
cd ../cpp
mkdir -p build
cd build

# 配置和编译
cmake ..
make

if [ $? -eq 0 ]; then
    echo "C++库构建成功！"
else
    echo "C++库构建失败！"
    exit 1
fi

# 回到go目录
cd ../../go

echo "开始运行Go程序..."

# 设置库路径环境变量
export LD_LIBRARY_PATH=../cpp/build:$LD_LIBRARY_PATH

# 运行Go程序（包含当前目录下的所有Go文件）
go run .

echo "程序运行完成！"
