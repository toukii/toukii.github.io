# go内置函数

## recover

recover与defer配合使用，而且！！！recover只能在defer的func(匿名或非匿名)里调用才有效。有且只能包装一次！

```
/*
	recover 只能包装在defer func（匿名或非匿名）中，只能包装一层。
*/

func TestRecover(t *testing.T) {
	resp, err := PerfectResp()
	t.Logf("resp:%+v, err:%+v", resp, err)
}

type Resp struct {
	Code string
	Msg  string
}

// resp:&{Code:success! Msg:}, err:something going wrong
func PerfectResp() (resp *Resp, err error) {
	defer func() {
		if cerr := recover(); cerr != nil {
			fmt.Printf("recover err:%+v\n", cerr)
			err = fmt.Errorf("something going wrong")
		}
	}()

	resp = &Resp{
		Code: "success!",
	}

	panic("go wrong!")
}

```

## [defer](https://mp.weixin.qq.com/s/t5tmqsjZ4y4Z_n6u4c9bMw)

1. defer 先装载好其func的参数，最后在func执行之前返回到主程。此过程为串行调用。

```
func main(){
	mainfunc()
}

func mywarpfunc(i int) func() {
	fmt.Println("mywarpfunc begine", i)
	return func() {
		i++
		fmt.Println("mywarpfunc 具体实现", i)
	}
}

/*
mywarpfunc begine 1
mainfunc 1
mywarpfunc 具体实现 2
*/
func mainfunc() {
	i := 1
	defer mywarpfunc(i)()
	fmt.Println("mainfunc", i)
}
```

2. defer func 局部变量未包含的变量，将会在上下文中查找。
3. 注意返回值的匿名或非匿名

```
	ret = deferFuncWithNamedReturnValue()
	fmt.Println("deferFuncWithNamedReturnValue ret: ", ret)
	ret = deferFuncWithNamedReturnValue2()
	fmt.Println("deferFuncWithNamedReturnValue2 ret: ", ret)

/*
retVal: 0
retVal: 1
deferFuncWithNamedReturnValue ret:  2
*/
func deferFuncWithNamedReturnValue() (retVal int) {
	r := 2
	defer func(retVal int) {
		fmt.Println("retVal:", retVal)
		retVal++
		fmt.Println("retVal:", retVal)
	}(retVal)
	return r
}

/*
retVal: 2
retVal: 3
deferFuncWithNamedReturnValue2 ret:  3
*/
func deferFuncWithNamedReturnValue2() (retVal int) {
	r := 2
	defer func() {
		fmt.Println("retVal:", retVal)
		retVal++
		fmt.Println("retVal:", retVal)
	}()
	return r
}

```

简而言之,延迟函数在声明时会收集相关参数赋值拷贝一份入栈,时机合适时再从入栈环境中寻找相关环境参数,如果找不到就扩大范围寻找外层函数是否包含所需变量,执行过程也就是延迟函数的出栈.

```
func test1() (x int) {
    defer fmt.Printf("in defer: x = %d\n", x)
    x = 7
    return 9
}

func test2() (x int) {
    x = 7
    defer fmt.Printf("in defer: x = %d\n", x)
    return 9
}

func test3() (x int) {
    defer func() {
        fmt.Printf("in defer: x = %d\n", x)
    }()

    x = 7
    return 9
}

func test4() (x int) {
    defer func(n int) {
        fmt.Printf("in defer x as parameter: x = %d\n", n)
        fmt.Printf("in defer x after return: x = %d\n", x)
    }(x)

    x = 7
    return 9
}

func main() {
    fmt.Println("test1")
    fmt.Printf("in main: x = %d\n", test1())
    fmt.Println("test2")
    fmt.Printf("in main: x = %d\n", test2())
    fmt.Println("test3")
    fmt.Printf("in main: x = %d\n", test3())
    fmt.Println("test4")
    fmt.Printf("in main: x = %d\n", test4())
}
```

公布答案：
```
test1
in defer: x = 0
in main: x = 9
test2
in defer: x = 7
in main: x = 9
test3
in defer: x = 9
in main: x = 9
test4
in defer x as parameter: x = 0
in defer x after return: x = 9
in main: x = 9
```


```
https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01

type database struct{}
func (db *database) connect() (disconnect func()) {
  fmt.Println("connect")
  return func() {
    fmt.Println("disconnect")
  }
}

func() {
  db := &database{} // 如果defer设计私有变量传递，需要用指针
  close := db.connect()
  defer close()
 
  fmt.Println("query db...")
}
```