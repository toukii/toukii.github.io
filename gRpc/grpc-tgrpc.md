# tgrpc

## doc

使用说明参照：https://github.com/tgrpc/doc

## descriptor

收集依赖fileDescriptorSet

1. protoc命令

生成代码时添加flag: --descriptor_set_out=.helloworld.Greeter.pbin 会保存proto描述。

2. [desc工具](https://github.com/tgrpc/desc)

参考之前的文章[protobuf-desc](protobuf-desc.html)

