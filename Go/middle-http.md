# http包中的扩展

面向扩展编程，比面向接口编程增加了可扩展的路径：函数。也可以理解为面向接口编程+面向方法编程。

## 两种ServeHTTP的扩展

```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

扩展Handler方法有两个途径：
1. 实现Handler接口
2. 扩展HandlerFunc方法，本质上和1没有区别，HandlerFunc实现了Handler接口，为扩展ServeHTTP增加了一条路径。

一般，使用func扩展func，使用struct扩展interface。

```
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

http包中提供了以下ServeHTTP的扩展：

```
HandlerFunc
Server
fileHandler(FileServer)
stringHandler
RedirectHandler
redirectHandler
ServeMux
serverHandler
TimeoutHandler
timeoutHandler
globalOptionsHandler
initALPNRequest
StripPrefix
```

还有 http/httputil.ReverseProxy, http/pprof.handler

<!-- `Counter Chan FlagServer ArgServer DateServer Logger` -->

## 路由

### 注册路由

```
// 还提供了ServeTLS
func Serve(l net.Listener, handler Handler) error {
	srv := &Server{Handler: handler}
	return srv.Serve(l)
}
```

```
// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	mux.Handle(pattern, HandlerFunc(handler))
}

// Handle registers the handler for the given pattern.
func (mux *ServeMux) Handle(pattern string, handler Handler) {
}
```

mux.HandleFunc的第二个参数也是可扩展的：可以定义一个函数，返回结果为`func(ResponseWriter, *Request)`, 函数体实现一些中间处理逻辑。
mux内部用map保存handler的路由映射：

```
e := muxEntry{h: handler, pattern: pattern}
mux.m[pattern] = e
```

### 解析路由

```
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```

mux.Handler解析request的host和URL，`h, pattern = mux.match(host + path)`, 返回该请求的 handler(Handler), 最后调用该handler的实现：ServeHTTP。
mux.Handler也会处理一些异常情况，如CONNECT，路由不匹配，这些异常情况会返回重定向handler：RedirectHandler。

ServeMux的路由规则：支持注册路径：host+path或path， path可以为 /xxx/ 或 /xxx。
path为 /xxx/的路由，会放到一个按路径长度降序的数组里（appendSorted），优先返回匹配最长的路由。


## 扩展

扩展 HandleFunc

### 链式拦截器扩展


```
func FooHttpInterceptor(rw http.ResponseWriter, req *http.Request, handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("FooHttpInterceptor", req.RequestURI)
		handler(rw, req)
	}
}
```

```
func Foo(h http.HandlerFunc) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		log.Println(req.RequestURI)

		h(rw, req)
	}
	return http.HandlerFunc(fn)
}
```

### 数组式拦截器扩展

```
var (
	_middlewares []Middleware
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func WebapiInceptor(middlewares ...Middleware) func(originFunc http.HandlerFunc) http.HandlerFunc {
	_middlewares = append(_middlewares, middlewares...)
	return WebapiWarp
}

func WebapiWarp(originFunc http.HandlerFunc) http.HandlerFunc {
	log.Println("WebapiWarp...")
	currentFunc := originFunc
	for _, it := range _middlewares {
		currentFunc = it(currentFunc)
	}
	return currentFunc
}
```

扩展 Handler

如 `beego/mux.Mux`, 'gorilla.mux.Router'。
