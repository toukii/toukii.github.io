# 性能优化

```
It you can’t measure it, you can’t manage it
你如果无法度量它，就无法管理它。
--- 彼得德鲁克(管理学大师)
```


## 性能优化方法论 [转载](https://mp.weixin.qq.com/s/snQ3T86usv4rXz0MMQvFfQ)

性能优化与解bug不同，对工程师的技术广度和技术深度有所要求。性能问题不仅包含代码问题，还包括容器、操作系统、网络、存储、文件系统等问题。

1. 了解应用的的总体架构，比如应用的外部依赖和核心接口有哪些，使用了哪些组件和框架，哪些接口、模块的使用率较高，上下游的数据链路是怎么样的等；
2. 了解应用对应的服务器信息，如服务器所在的集群信息、服务器的 CPU/内存信息、安装的 Linux 版本信息、服务器是容器还是虚拟机、所在宿主机混部后是否对当前应用有干扰等；

性能优化的痛点：

```
无法确定性能瓶颈点
性能优化的分析思路不明确
如何选择性能优化的分析工具
```

性能优化流程


准备阶段：主要工作是是通过性能测试，了解应用的概况、瓶颈的大概方向，明确优化目标；

分析阶段：通过各种工具或手段，初步定位性能瓶颈点；

调优阶段：根据定位到的瓶颈点，进行应用性能调优；

测试阶段：让调优过的应用进行性能测试，与准备阶段的各项指标进行对比，观测其是否符合预期，如果瓶颈点没有消除或者性能指标不符合预期，则重复步骤2和3。

__通过压测或系统日常监控确定优化目标，如数据库、内存等；通过采集系统指标确定性能瓶颈点；执行调优然后验证优化效果；重复执行优化，直到达到优化目标。__

注意事项

1. 性能瓶颈点通常呈现 2/8 分布，即80%的性能问题通常是由20%的性能瓶颈点导致的，2/8 原则也意味着并不是所有的性能问题都值得去优化；
2. 性能优化是一个渐进、迭代的过程，需要逐步、动态地进行。记录基准后，每次改变一个变量，引入多个变量会给我们的观测、优化过程造成干扰；
3. 不要过度追求应用的单机性能，如果单机表现良好，则应该从系统架构的角度去思考; 不要过度追求单一维度上的极致优化，如过度追求 CPU 的性能而忽略了内存方面的瓶颈；
4. 选择合适的性能优化工具，可以使得性能优化取得事半功倍的效果；
5. 整个应用的优化，应该与线上系统隔离，新的代码上线应该有降级方案。

## Go性能优化

### http/pprof

```
import _ "net/http/pprof"

http.ListenAndServe("localhost:8080", nil)
```

```
import "net/http/pprof"

mux.HandleFunc("/debug/pprof/", pprof.Index)
mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
```

### 指标采集

1. go tool pprof

```
go tool pprof http://127.0.0.1:8080/debug/pprof/profile\?seconds\=30
go tool pprof -alloc_space http://127.0.0.1:8080/debug/pprof/heap
```

日志里有有采集数据的保存路径
Saved profile in /Users/xxx/pprof/pprof.xxx.samples.cpu.002.pb.gz

```
go-torch /Users/xxx/pprof/pprof.xxx.samples.cpu.002.pb.gz
```

2. curl

```
curl http://127.0.0.1:8080/debug/pprof/profile >> app_cpu.profile
curl http://localhost:8080/debug/pprof/heap >> app_heap.profile
```

3. go-torch

```
go-torch -u http://127.0.0.1:8080 -p > torch.svg
```

### 火焰图工具 go-torch

CPU的火焰图，需要优化的点是比较平摊的区域。
