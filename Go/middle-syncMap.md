## sync.Map 源码解析 [转载](https://colobu.com/2017/07/11/dive-into-sync-Map/)

sync.Map提供了如下方法：

```
type Map
    func (m *Map) Delete(key interface{}) 删除
    func (m *Map) Load(key interface{}) (value interface{}, ok bool) 查询
    func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) setNx
    func (m *Map) Range(f func(key, value interface{}) bool) 遍历
    func (m *Map) Store(key, value interface{}) 更新
```

实现原理是在其底层有两份存在冗余的数据存储：`dirty`、`read`; 其中，对read只读，不需要加锁，在更新kv时会操作dirty，加锁。

Load：一般地，读取数据都可以从read中读到，若read中没有读到，查找非空dirty，则miss+1；当miss>len(dirty)时，需要将dirty转为read，dirty置空。
加锁查找dirty时，用到了著名的双检查（java单例模式常用方法）。


