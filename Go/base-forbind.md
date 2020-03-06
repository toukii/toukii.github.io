# 循环变量绑定[转载](https://rainbowmango.gitbook.io/go/appendix/unpinned)

题目1

```
func Process1(tasks []string) {
    for _, task := range tasks {
        // 启动协程并发处理任务
        go func() {
            fmt.Printf("Worker start process task: %s\n", task)
        }()
    }
}
```

题目2

```
func Process2(tasks []string) {
    for _, task := range tasks {
        // 启动协程并发处理任务
        go func(t string) {
            fmt.Printf("Worker start process task: %s\n", t)
        }(task)
    }
}
```

题目3

```
func TestDouble(t *testing.T) {
    var tests = []struct {
        name         string
        input        int
        expectOutput int
    }{
        {
            name:         "double 1 should got 2",
            input:        1,
            expectOutput: 2,
        },
        {
            name:         "double 2 should got 4",
            input:        2,
            expectOutput: 4,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            if test.expectOutput != Double(test.input) {
                t.Fatalf("expect: %d, but got: %d", test.input, test.expectOutput)
            }
        })
    }
}
```

题目4
```
type Car struct {
	Name string
}

cars := []Car{
	Car{Name: "Name1"},
	Car{Name: "Name2"},
	Car{Name: "Name3"},
	Car{Name: "Name4"},
}

newcars := make([]*Car, 0, len(cars))
newcars2 := make([]*Car, 0, len(cars))
for _, car := range cars {
	newcars = append(newcars, &car)
	tmp := car
	newcars2 = append(newcars2, &tmp)
}
```


1. for循环里的变量，每次循环都会给重新复制，如果要在协程闭包里使用, 需要用新(临时)变量绑定值：称为显示绑定，题目4中的tmp。
2. 也可以通过参数传递规避上面这种情况，其实参数传递也是用新变量(实参)绑定值，如题目2。
3. for循环里的变量是__同一个变量__，所以取变量的地址，所有循环都是一样的。
特别是range 后跟非指针数组（值数组）, 如果需要获得该数组里某些item的地址，需要使用新(临时)变量绑定值（显示绑定）；值数组的情况同样适用于值map。


## 循环变量是易变的

循环变量实际上只是一个普通的变量。
语句`for index, value := range xxx`中，__每次循环index和value都会被重新赋值（并非生成新的变量）__。

如果循环体中会启动协程（并且协程会使用循环变量），就需要格外注意了，因为很可能循环结束后协程才开始执行， 此时，__所有协程使用的循环变量有可能已被改写__。

## 循环变量需要绑定

__协程从被创建到被调度执行期间循环变量极有可能被改写。__

在题目一中，协程函数体中引用了循环变量task，， 这种情况下，我们称之为变量没有绑定。 所以，题目一打印结果是混乱的。很有可能（随机）所有协程执行的task都是列表中的最后一个task。

__实参传递，实际上也生成了新的变量，也即间接完成了绑定__

在题目二中，协程函数体中并没有直接引用循环变量task，而是使用的参数。而在创建协程时，循环变量task 作为函数参数传递给了协程。参数传递的过程实际上也生成了新的变量，也即间接完成了绑定。 所以，题目二实际上是没有问题的。

在题目三中，测试用例名字test.name通过函数参数完成了绑定，而test.input 和 test.expectOutput则没有绑定。 然而题目三实际执行却不会有问题，因为t.Run(...)并不会启动新的协程，也就是循环体并没有并发。 些时，即便循环变量没有绑定也没有问题。 但是风险在于，如果t.Run(...)执行的测试体有可能并发（比如通过t.Parallel()），此时就极有可能引入问题。
对于题目三，建议显式的绑定。