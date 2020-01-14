## 分布式事务 [转载](https://juejin.im/post/5b5a0bf9f265da0f6523913b)

分布式事务的解决方案：

```
1. XA
2. 消息中间件 mq 幂等实现
3. try confirm cancel TCC
```


CAP定律 : 一致性，可用性，分区容错， 3个考量只能取其2.

Base理论 基本可用，软状态，最终一致性. 对实际分布式事务有指导意义。
