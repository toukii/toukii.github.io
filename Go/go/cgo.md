需要自己设置 头文件 和 库文件的地址

比如 cgo参数的使用

//#cgo CFLAGS:  -I../../cinclude
//#cgo LDFLAGS: -L../../clib -lgsasl

//#include <gsasl.h>

......

CFLAGS

上边指示了头文件地址

LDFLAGS

下边的表明了库文件地址

都是当前文件的相对位置

-I (大写)指示了头文件目录

-L 指示了库文件目录 -l(L小写)指示所用的具体的某个库文件