#ifndef HELLO_H
#define HELLO_H

#ifdef __cplusplus
extern "C" {
#endif

// 暴露的函数：输出helloworld
void printHelloWorld();

void printAny(const char* s);

struct SdkConfig{
    int log_level;
    const char* log_path;
    const char* url;
};

void passSdkConfig(struct SdkConfig config);

void getGlobalConfig(struct SdkConfig* out);

int add(int a, int b);

#ifdef __cplusplus
}
#endif

#endif // HELLO_H
