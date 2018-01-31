#	Go工程结构

##	GOPATH

	go env
	可以看到go的环境，gopath为用户的工作路径，go get XXX 将把代码放到gopath中的src/XXX目录

###	目录

	---GOPATH
		---src pkg bin

		src 源码
		bin 编译的二进制文件
		pkg go install的包信息

##	导入包

	import "fmt" // 导入fmt包
	import golog "log" // 导入log包，并取别名golog
	import . "fmt" // 导入fmt包，使用时不必加前缀fmt.
	import _ "fmt" // 导入fmt包，并执行fmt包中的init函数

	ps: 先执行导入包的init函数，再执行本包的init函数，最后执行main函数

##	Test

	go的测试文件命名：XXX_test.go, 在该目录下运行 go test命令

###	单元测试

	import "testing"

	func TestAFunc(t *testing.T){
		ret := AFunc()
		if ret == nil{
			t.Error("ret is nil.")
		}
	}

	$go test -v

###	性能测试

	import "testing"

	func BenchmarkAFunc(b *testing.B){
		for i:=0;i<b.N;i++{
			_ = AFunc()
		}
	}

	$go test -bench=. -benchtime 2s

###	样例测试

	func ExampleAFunc(){
		println("Hello,Go.")
		// Output: Hello,Go.
	}

## CGO GOC

	package main

	/*
	#include "hello.h"
	#include <stdio.h>

	void sayHi() {
	  printf("Hi");
	  sayHello();
	}
	*/
	import "C"

	func main() {
		C.sayHi()
	}


## 关键字用法

### chan

请__[参考](https://github.com/everfore/gotest/tree/master/channel)__

	chan是同步的，带缓冲区的chan是非阻塞的，不带缓冲区的chan是阻塞的；
	从closed chan中读是非阻塞的，往closed chan中写会panic。
	往nil chan中写,从nil chan中读会panic。

		package main

		import (
			"time"
		)

		func main() {
			c := make(chan bool)
			go func() {
				for {
					time.Sleep(1e9)
					c <- true
					close(c)
					break
				}
			}()
			for {
				select {
				case cin, ok := <-c:
					println(cin, ok)
				}
			}
		}
	// closed chan可以非阻塞地读取到false,false;程序会不停地打印输出。

###	select 

		package main

		import (
			"time"
		)

		func main() {
			c := make(chan bool)
			go func() {
				for {
					time.Sleep(1e9)
					c <- true
				}
			}()
			for {
				select {
				case cin, ok := <-c:
					println(cin, ok)
					// default:
					// println("nothing")
				}
			}
		}
	// CPU使用率很低，基本为0
	// 若不注释default一行，CPU利用率很高，测试环境为4核，CPU使用率为25%左右。
	// 若不注释println一行，每次时钟周期到了会打印一行，CPU使用率同上。


**阻塞的写法**

		for {
			cin, ok := <-c
			println(cin, ok)
		}
