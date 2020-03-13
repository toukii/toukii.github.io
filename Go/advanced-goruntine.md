# go 并发模型[转载](https://www.cnblogs.com/mokafamily/p/9975980.html)

## 并发与并行

并发是程序执行逻辑，并行是物理机支持的任务被同时进行处理。

```
Concurrency vs. parallelism
Concurrency is about dealing with lots of things at once.
Parallelism is about doing lots of things at once.
Concurrency is about structure, parallelism is about execution.
--- Rob Pikie
```

CPU密集的任务，高并发效果不明显；I/O密集型任务高并发效果明显。
高并发是指并发数大大多于并行使，想通过增加并发数来提高处理速度的收益。

## Goruntine

系统线程是内核级的线程，通常开启多线程的任务，需要在用户态与内核态间切换上下文，所以会对开启的线程数有限制。

goruntine是对系统线程封装并暴露出的用户级线程，goruntine的切换是用户态的切换。用户级线程到内核级线程的调度由golang的runtime负责，调度逻辑对外透明。所以golang可以开启成千上万的协程。

__G·P·M__

M（Machine） ：对内核级线程的封装，数量对应真实的CPU数（真正干活的对象）

P（Processor） ：即为G和M的调度对象，用来调度G和M之间的关联关系，其数量可通过GOMAXPROCS()来设置，默认为核心数

G（Goroutine） ：我们所说的协程，为用户级的轻量级线程，每个Goroutine对象中的sched保存着其上下文信息

```
All problems in computer science can be solved by another level of indirection.
计算机科学中遇到的所有问题都可通过增加一层抽象来解决。
--- David Wheeler(伟大的计算机科学家)
```

![](https://www.ardanlabs.com/images/goinggo/94_figure9.png)

LRQ（Local Run Queue）
GRQ（Global Run Queue ）

每个Processor对象都拥有一个LRQ（Local Run Queue），未分配的Goroutine对象保存在GRQ（Global Run Queue ）中，等待分配给某一个P的LRQ中，每个LRQ里面包含若干个用户创建的Goroutine对象，同时Processor作为桥梁对Machine和Goroutine进行了解耦，也就是说Goroutine如果想要使用Machine需要绑定一个Processor才行，上图中共有两个M和两个P也就是说我们可以同时并行处理两个goroutine。
