cmake_minimum_required(VERSION 3.10)

# 设置项目名称
project(HelloWorld)

# 设置C++标准
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# 设置位置无关代码（PIC），这对于共享库很重要
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

# 创建库文件
add_library(hello_lib STATIC
    hello.cpp
)

# 设置库文件的输出目录
set_target_properties(hello_lib PROPERTIES
    ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}
)

# 创建可执行文件
add_executable(hello_main
    main.cpp
)

# 链接库到可执行文件
target_link_libraries(hello_main hello_lib)