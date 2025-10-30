#include "hello.h"
#include <iostream>
#include <string>

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
}