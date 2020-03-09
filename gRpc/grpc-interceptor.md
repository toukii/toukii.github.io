# grpc 拦截器

grpc/interceptor.go中定义了4中拦截器：


```
type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error

type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)

type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)

type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
```

总体上，可以分为服务端拦截器和客户端拦截器；再细分可以分为Unary和Stream两种。本文以UnaryInterceptor为例，分别实现服务端和客户的拦截器。

## UnaryServerInterceptor

先看一下grpc.NewServer方法，返回一个grpc Server，入参是ServerOption.

```
func NewServer(opt ...ServerOption) *Server {
}
```

ServerOption中的options是非导出结构，所以扩展包中不能能返回ServerOption。

```
type ServerOption func(*options)
```

grpc中提供了若干种将扩展（中间件）转换为ServerOption的方法：

```
func UnaryInterceptor(i UnaryServerInterceptor) ServerOption {
}
func StreamInterceptor(i StreamServerInterceptor) ServerOption {
}
```

此处提供了多种转为ServerOption的方法，为中间件提供了接口。

所以，加载服务的的拦截器是通过 grpc.UnaryInterceptor(MyUnaryServerInterceptor)。


以下是实现一个鉴权服务端拦截器：

```
func AuthInterceptor(ctx context.Context, req interface{}, 
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, 
		err error) {
	cookie, err := extractHeader(ctx, "cookie")
	if err != nil {
		log.Errorf("%+v", err)
		return nil, errors.Cause(err)
	}
	log.Infof("cookie:%+v", cookie)

	if err := biz.Auth(cookie, conf.Conf.AESKeybs, conf.Conf.AESSalt); 
	err != nil {
		log.Errorf("%+v", err)
		return nil, errors.Cause(err)
	}

	return handler(ctx, req)
}
```

一般，服务端的拦截器有很多。我们需要一个链式拦截器, [`github.com/grpc-ecosystem/go-grpc-middleware`](https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/chain.go#L19):

```
func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, info, currentHandler)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, req)
	}
}
```

服务端注册中间件的最终代码：

```
grpc.NewServer(
	grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(AuthInterceptor, XXInterceptor)))
```

## UnaryClientInterceptor

客户端的中间件添加在 `grpc.DialOption` 中，使用`grpc.WithUnaryInterceptor(MyClientInterceptor)` 包装。

[`github.com/grpc-ecosystem/go-grpc-middleware`](https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/chain.go#L66)也提供了Client链式拦截器：ChainUnaryClient。

