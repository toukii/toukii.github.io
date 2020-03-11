# 内存分配[转载](https://tonybai.com/2020/03/10/visualizing-memory-management-in-golang/)

Go使用线程本地缓存(thread local cache)来加速小对象分配，并维护着scan/noscan的span来加速GC。这种结构以及整个过程避免了碎片，从而在GC期间无需做紧缩处理。

![](https://tonybai.com/wp-content/uploads/visualizing-memory-management-in-golang-3.png)

mheap: 页堆page heap
大对象（大小> 32kb的对象）直接从mheap分配。这些大对象申请请求是以获取中央锁(central lock)为代价的，因此在任何给定时间点只能满足一个P的请求。
mheap通过将页归类为不同结构进行管理的：mspan, mcentral, mcache, arena

arena:
堆在已分配的虚拟内存中根据需要增长和缩小。当需要更多内存时，mheap从虚拟内存中以每块64MB（对于64位体系结构）为单位获取新内存， 这块内存被称为arena。这块内存也会被划分页并映射到span。

mspan: mspan是mheap中管理的内存页的最基本结构。这是一个双向链接列表，其中包含起始页面的地址，span size class和span中的页面数量。像TCMalloc一样，Go将内存页按大小分为67个不同类别，大小从8字节到32KB，如下图所示

mcentral: 
mcentral将相同大小级别的span归类在一起。如果mcentral没有可用的span，它将向mheap请求新页。

mcache:
这是一个非常有趣的构造。mcache是提供给P（逻辑处理器）的高速缓存，用于存储小对象（对象大小<= 32Kb）。尽管这类似于线程堆栈，但它是堆的一部分，用于动态数据。所有类大小的mcache包含scan和noscan类型mspan。Goroutine可以从mcache没有任何锁的情况下获取内存，因为一次P只能有一个锁G。因此，这更有效。mcache从mcentral需要时请求新的span。

总结：
mheap有若干个arena(64MB)组成；每个arena由若干个mcentral组成；每个mcentral有相同大小的mspan组成，这些不同大小的mcentral共同组成arena；mspan的大小有不同类别：像TCMalloc一样，Go将内存页按大小分为67个不同类别，大小从8字节到32KB。
mcache是为P提供的高速缓存，<= 32Kb；大于32Kb直接分配在arena。

## GC
