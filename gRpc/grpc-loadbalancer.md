# grpc负载均衡

## 路由解析

`google.golang.org/grpc/naming/naming.go`


`Resolver` 用于监控一个服务target，得到的Watcher可以跟踪服务地址的变化。维护此服务地址是负载均衡的基础。

```
// Resolver creates a Watcher for a target to track its resolution changes.
type Resolver interface {
	// Resolve creates a Watcher for target.
	Resolve(target string) (Watcher, error)
}
```

`Watcher` 用于轮训，定期检测变化的地址信息。地址变化包括新增和删除，见Operation的定义

```
// Watcher watches for the updates on the specified target.
type Watcher interface {
	// Next blocks until an update or error happens. It may return one or more
	// updates. The first call should get the full set of the results. It should
	// return an error if and only if Watcher cannot recover.
	Next() ([]*Update, error)
	// Close closes the Watcher.
	Close()
}
```

```
type Operation uint8

const (
	// Add indicates a new address is added.
	Add Operation = iota
	// Delete indicates an exisiting address is deleted.
	Delete
)
```

## 负载均衡

`google.golang.org/grpc/balancer.go` 定义了Balancer接口：
```
type Balancer interface {
	Start(target string, config BalancerConfig) error
	Up(addr Address) (down func(error))
	Get(ctx context.Context, opts BalancerGetOptions) (addr Address, put func(), err error)
	Notify() <-chan []Address
	Close() error
}
```

同时，提供了一种负载均衡的实现：

```
func RoundRobin(r naming.Resolver) Balancer {
	return &roundRobin{r: r}
}
```

## 调用

grpc的dail option中有一个option：WithBalancer
```
opts := []grpc.DialOption{
	grpc.WithInsecure(),
	grpc.WithBalancer(Resolver),
}
```

WithBalancer定义：
```
func WithBalancer(b Balancer) DialOption {
}
```

不同的服务发现有不同的`naming.Resolver`实现方法，我们只需要实现Resolver即可使用官方的RoundRobin负载均衡。

grpc.WithBalancer(NewMyResolver(cfg))
