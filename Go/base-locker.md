## Mutex [转载](https://learnku.com/articles/39577?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)

锁分为4大类：

```
独占锁：mutex.Lock  rwmuxt.Lock
共享锁：rwmutx.RLock
乐观锁：不提前获取锁，真正执行的代码段才lock
悲观锁：提前获取锁lock
```

使用锁的注意点：

```
同一个协程不能连续多次调用 Lock, 否则发生死锁
锁资源时尽量缩小资源的范围，以免引起其它协程超长时间等待
mutex 传递给外部的时候需要传指针，不然就是实例的拷贝，会引起锁失败
善用 defer 确保在函数内释放了锁
使用 - race 在运行时检测数据竞争问题，go test -race ....，go build -race ....
善用静态工具检查锁的使用问题
使用 go-deadlock 检测死锁，和指定锁超时的等待问题 (自己百度工具用法)
能用 channel 的场景别使用成了 lock
```

