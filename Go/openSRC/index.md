# go开源代码学习

## 1. [groupcache](https://github.com/golang/groupcache)

 - [参考中文博客](http://blog.csdn.net/npc_hp110/article/details/48036765)

groupcache提供了一个分布式k／v缓存解决方案。根据数据的k，计算hash（一致性哈希），读取数据：缓存在local节点或远端节点（Peer）。groupcache设计了相关接口，可以对方法进行扩展，如：
 
 1. 数据编码器，可以选择高性能的protobuf；

 2. 扩展Cache为LRU，LFU等Cache；

另外，groupcache还提供了热数据扩散、防止Load“惊群效应”、控制缓存内存大小（一般地，仅做到限制缓存的数量）等。

## 2. [grpc](grpc.html)

ProtoBuf 是一套接口描述语言（IDL）和相关工具集（主要是 protoc，基于 C++ 实现），类似 Apache 的 Thrift）。用户写好 .proto 描述文件，之后使用 protoc 可以很容易编译成众多计算机语言（C++、Java、Python、C#、Golang 等）的接口代码。这些代码可以支持 gRPC，也可以不支持。

gRPC 是 Google 开源的 RPC 框架和库，已支持主流计算机语言。底层通信采用 gRPC 协议，比较适合互联网场景。gRPC 在设计上考虑了跟 ProtoBuf 的配合使用。

两者分别解决的不同问题，可以配合使用，也可以分开。

典型的配合使用场景是，写好 .proto 描述文件定义 RPC 的接口，然后用 protoc（带 gRPC 插件）基于 .proto 模板自动生成客户端和服务端的接口代码。
