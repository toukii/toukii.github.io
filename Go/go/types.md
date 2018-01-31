#	本节提纲

1. 基本类型
2. 关键字及用法
3. 初始化
4. 目录结构

1.

	类型	长度	默认值	说明
	bool	1	false	
	byte	1	0	uint8
	rune	4	0	Unicode Code Point, int32
	int, uint	4 或 8	0	32 或 64 位
	int8, uint8	1	0	-128 ~ 127, 0 ~ 255
	int16, uint16	2	0	-32768 ~ 32767, 0 ~ 65535
	int32, uint32	4	0	-21亿 ~ 21 亿, 0 ~ 42 亿
	int64, uint64	8	0	
	float32	4	0.0	
	float64	8	0.0	
	complex64	8		
	complex128	16		
	uintptr	4 或 8		足以存储指针的 uint32 或 uint64 整数
	array			值类型
	struct			值类型
	string		""	UTF-8 字符串
	slice		nil	引用类型
	map		nil	引用类型
	channel		nil	引用类型
	interface		nil	接口
	function		nil	函数

2.


	make
	len
	cap
	select
	for range
	func
	go
	defer mtx.Unlock
	revocer


#	Go 基础

## 基本类型

Go的变量都有默认值。

	（64位机器）
	类型	字节	默认值	别称

	bool  1	false

	uint8	1	0	byte
	uint16	2	0
	uint32	4	0
	uint64	8	0

	int8	1	0
	int16	2	0	rune
	int32 	4	0	int
	int64	8	0
	
	float32	4	0.0
	float64	8	0.0

	complex64	8	0+0i
	complex128	16	0+0i

	string	-	""

	uintptr	4/8	0 // uint64

初始化方式：

	var age byte
	year := 2015 // int32

真实的类型与别称的转换不需要显示转换，如：

	var r rune
	var i32 int32
	i32 = (int32)(r) // 或者 i32 = r

##	其它类型

	名称			解释			初始化
	array		数组				arr := [...]int{1,2,3}
	slice		切片，动态数组		slc := []int{1,2,3} // slc := make([]int, 3, 6)
	map 		键值对映射			mp := make(map[string]int,10)
	struct		结构体				stt := &User{Name:"shaalx"}
	chan		通道				c := make(chan bool, 3)
	interface	接口				*
	func		函数也可以作为变量 	*

##	关键字

	len 测试容器的长度，这里的容器包括：array,slice,map,chan
	cap 测试slice的容量	（PS:容量不小于len）
	make 初始化：array,slice,map,chan
	close 关闭chan的写
	for range 遍历：array,slice,map,chan
	for 循环
	select 非阻塞地选择chan
	switch 分支，可选类型有string,int,string,type等等
	go 开启一个goruntine
	defer 延迟执行,(栈结构FILO)
	panic 引发恐慌
	recover 恢复panic
	