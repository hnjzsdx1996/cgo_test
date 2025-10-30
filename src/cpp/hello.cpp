#include "hello.h"
#include <iostream>
#include <string>

// 深拷贝持久化存储，避免悬垂指针
static std::string g_log_path;
static std::string g_url;
static int g_log_level = 0;

void printHelloWorld() {
    std::cout << "helloworld" << std::endl;
}

void printAny(const char* s) {
    std::cout << s << std::endl;
}

void passSdkConfig(struct SdkConfig config) {
    const std::string log_path_copy = config.log_path ? std::string(config.log_path) : std::string();
    const std::string url_copy = config.url ? std::string(config.url) : std::string();

    std::cout << "log_level: " << config.log_level << std::endl;
    std::cout << "log_path: " << log_path_copy.c_str() << std::endl;
    std::cout << "url: " << url_copy.c_str() << std::endl;

    // 更新全局持久化副本（深拷贝）
    g_log_level = config.log_level;
    g_log_path = log_path_copy;
    g_url = url_copy;
}

void getGlobalConfig(struct SdkConfig* out) {
    if (!out) return;
    out->log_level = g_log_level;
    out->log_path = g_log_path.c_str();
    out->url = g_url.c_str();
}

int add(int a, int b) {
    return a + b;
}