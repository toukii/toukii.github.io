# 内存对齐[转载](https://ms2008.github.io/2019/08/01/golang-memory-alignment/)

内存对齐可以更省内存，提高访问内存性能，特殊代码可以防止在不同平台程序异常。
主要有两个考量：

1. 平台的移植性
特定的硬件平台只允许在特定地址获取特定类型的数据，否则会导致异常情况

2. 性能
cpu访问内存，一次读取的是平台读取单位（“对齐系数”），对齐的内存可以减少cpu访问内存次数。
64位机的“对齐系数”是8, 32位机的“对齐系数”是4。

以下代码在64位机上编译为32位的二进制，运行会panic.

```
package main

import (
	"sync/atomic"
)

type T3 struct {
	b int64
	c int32
	d int64
}

func main() {
	a := T3{}
	atomic.AddInt64(&a.d, 1)
}
```

```
GOARCH=386 go build -o a
./a
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x30cc]
```

解决方法：

```
type T3 struct {
	b int64
	c int32
	_ int32
	d int64
}
```

或：

```
type T3 struct {
	b int64
	d int64
	c int32
}
```

## 内存对齐后的内存大小

```
type T1 struct {
	a int8
	b int64
	c int16
}

type T2 struct {
	a int8
	c int16
	b int64
}
```

64位机，对齐系数是8byte。
T1.a 8byte, T1.b 8byte, T1.c 8byte 共24byte。
T2.a+T2.b 8byte, T2.c 8byte  共16byte。

32位机，对齐系数是4byte。
T1.a 4byte, T1.b 8byte, T1.c 4byte 共16byte。
T2.a+T2.b 4byte, T2.c 8byte 共12byte。
