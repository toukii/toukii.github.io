# go pprof


```
go test -bench=Benchmark_001_short_NCGet -benchmem -cpuprofile prof.cpu
go test -bench=Benchmark_001_short_NCGet -benchmem -memprofile prof.mem
```

生成 `jsnm.test`, `prof.mem`, `prof.cpu`文件；

```
go tool pprof prof.mem

#top N：显示N个占用CPU时间最多的函数;默认10个
#top N -cum：显示N个占用CPU时间（包括其调用的子函数）最多的函数

top

top -cum

## 生成火焰图
go-torch jsnm.test prof.cpu
```



