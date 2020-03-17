# ratelimiter

限流有多种实现，常用的方法有：

1. 计数器算法 每秒控制n次请求，先到先得到服务，拒绝多余的请求
2. 漏桶算法 每秒控制n次请求，匀速控制。
3. 令牌桶算法 综合1,2两种方法的折中，基本是匀速控流，当有多个可用令牌是允许瞬时处理多个请求。

## 漏桶算法

istio中的限流实现：

github.com/istio/istio/pkg/mcp/rate

```
import "golang.org/x/time/rate"

type Limit interface {
	Wait(ctx context.Context) (err error)
}

func (c *Limiter) Create() Limit {
	return rate.NewLimiter(
		rate.Every(c.connectionFreq),
		c.connectionBurstSize)
}
```

## 令牌桶算法

github.com/tgrpc/interceptor/ratelimit

可以参考源代码，已使用atomic提升一倍性能。

```
// 环形队列实现令牌桶限流
type TokenBucket struct {
	Size       int32
	head, rear int32
	sync.Mutex
}

func NewTokenBucket(size int32) *TokenBucket {
	tb := &TokenBucket{
		Size: size + 1,
	}
	go tb.loopPut()
	return tb
}

func (tb *TokenBucket) next(cur int32) int32 {
	return (cur + 1) % tb.Size
}

func (tb *TokenBucket) Length() int32 {
	return (tb.rear + tb.Size - tb.head) % tb.Size
}

func (tb *TokenBucket) Limit() bool {
	tb.Lock()
	defer tb.Unlock()
	// 无令牌可用
	if tb.head == tb.rear {
		return false
	}
	tb.head = tb.next(tb.head)
	fmt.Println("Limit", tb.head, tb.rear, tb.Length())
	return true
}

// 将令牌放入桶中
func (tb *TokenBucket) put() bool {
	tb.Lock()
	defer tb.Unlock()
	// 桶已满，不需要放入
	next := tb.next(tb.rear)
	if next == tb.head {
		return false
	}
	tb.rear = next
	return true
}

func (tb *TokenBucket) loopPut() {
	dur := time.Duration(int64(time.Second) / int64(tb.Size-1))
	fmt.Printf("duration:%+v\n", dur)
	ticker := time.NewTicker(dur)
	for {
		select {
		case <-ticker.C:
			if tb.put() {
				fmt.Println("put ", tb.head, tb.rear, tb.Length())
			}
		}
	}
}
```

## grpc限流中间件


```
import "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"

type Limiter interface {
	Limit() bool
}

func UnaryServerInterceptor(limiter Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if limiter.Limit() {
			return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later.", info.FullMethod)
		}
		return handler(ctx, req)
	}
}
```
