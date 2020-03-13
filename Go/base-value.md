# 值类型

类型： 值类型和引用类型讨论的是：函数传递参数时，函数内修改参数的值是否可以真正修改实参的值。

__在函数传值时，一切按值传递！__

```
func main(){
	p := []int{1, 2, 3, 4}
	Slice(p, 3)
	fmt.Println("Slice p-len", len(p))
	SlicePtr(&p, 3)
	fmt.Println("SlicePtr p-len", len(p))
}

func Slice(p []int, n int) {
	if len(p) > n {
		p = p[:n]
	}
	fmt.Println(len(p))
}

func SlicePtr(p *[]int, n int) {
	if len(*p) > n {
		*p = (*p)[:n]
	}
	fmt.Println(len(*p))
}
```

输出：

```
3
Slice p-len 4
3
SlicePtr p-len 3
```

Slice 实参是切片的地址的拷贝，对地址指向的内容修改是生效的；但是对地址的指向修改时无效的。
如果要修改地址的指向，需要传入地址的地址，如SlicePtr。

by the way, 切片修改长度,可以用：
p = append(p[:0], 1)
p = p[:1]
