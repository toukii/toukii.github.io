# 基础语法

## 1. 关键字 keyword

```
if else switch case default fallthrough
for range
new make len cap append delete copy close panic recover
go select defer
return goto break continue
complex imag real
map chan func interface
package import var const type struct
print println
```

#### fallthrough

case 满足条件时，强制进入下一个case；

#### for range

可遍历的类型有：slice(i,t)，array(i,t)，map(k,v)，string(i,rune)，chan(i,t)

#### new

返回指针

#### make

申请内存，slice(len,cap)，map(cap)，chan(cap)

make 语法：

```
s := make([]T,s_len,s_cap)
len(s) = s_len
cap(s) = s_cap

m := make(map[T1]T2,m_cap)
len(m) <> m_cap
// cap(m) is wrong.

c := make(chan T,c_cap)
len(c) <> c_cap
cap(c) = c_cap
```

```
m := make(map[string]bool, 10)
m["map-key1"] = true
fmt.Println(len(m)) // 1

c := make(chan bool, 9)
c <- false
c <- true
fmt.Println(len(c), cap(c)) // 2 9
```

#### append

往slice里续写，若`cap(slice)>len(slice)`, `len(slice)`会+1;
否则，会扩容：cap`增加(2倍，但不是无限2倍扩容)`，len(slice)增加1


[延伸阅读](http://sharecore.net/2014/01/09/%E5%AF%B9Go%E7%9A%84Slice%E8%BF%9B%E8%A1%8CAppend%E7%9A%84%E4%B8%80%E4%B8%AA%E5%9D%91/)

#### delete

删除map里的某个值

## 2. 类型

 - 1 基本类型
```
bool
byte int8
complex(64|128) real imga
error
float(32|64)
rune
string
u?int(8|16|32|64)?
int int32
uintptr
```

 - 2 其他类型

```
struct
map
slice|array
interface
func
```

## 3. 声明&初始化

```
var i int
j := 1
```

[诡异:区别于其他语言](base2.md)

