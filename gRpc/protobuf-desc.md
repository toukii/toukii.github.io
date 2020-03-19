# protobuf descriptor

## Descriptor

go-grpc官方的Descriptor定义在 github.com/golang/protobuf/protoc-gen-go/descriptor/descriptor.pb.go中。
已第三方gogo为例，代码见：https://github.com/gogo/protobuf/blob/master/protoc-gen-gogo/descriptor/descriptor.pb.go

Descriptor有以下一些类型，从它的名称可以看出是描述什么的，也可以参考之前的文章[prodobuf-descriptor](prodobuf-descriptor.html)：

```
FileDescriptorProto 文件描述
ServiceDescriptorProto 服务描述
MethodDescriptorProto 方法描述
DescriptorProto message描述
FieldDescriptorProto 字段描述
OneofDescriptorProto
EnumDescriptorProto
EnumValueDescriptorProto
```

Descriptor的关系：

```
type FileDescriptorProto struct {
	MessageType []*DescriptorProto
	EnumType    []*EnumDescriptorProto
	Service     []*ServiceDescriptorProto
	Extension   []*FieldDescriptorProto
	Options     *FileOptions
}

type ServiceDescriptorProto struct {
	Method               []*MethodDescriptorProto
	Options              *ServiceOptions
}

type MethodDescriptorProto struct {
	Options    *MethodOptions
}
```

FileDescriptorProto包含了若干ServiceDescriptorProto，ServiceDescriptorProto又包含了若干MethodDescriptorProto。message的定义放在FileDescriptorProto中。
几个Options是扩展用的，参考[protobuf-option](protobuf-option.html)。
FieldDescriptorProto可能依赖其他FileDescriptorProto的FieldDescriptorProto，如果要得到完整的Descriptor，需要收集依赖的FieldDescriptorProto, 即 其他proto的FileDescriptorProto, 最终得到一个FileDescriptorSet：

```
type FileDescriptorSet struct {
	File                 []*FileDescriptorProto
}
```

## fileDescriptor

生成grpc代码后，在pb.go文件中可以拿到 FileDescriptorProto，这个FileDescriptorProto只描述了当前proto文件的Descriptor。

```
var fileDescriptorCharts = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0x48, 0x2c,
	0x2a, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0xcc, 0x4b, 0xcc, 0xa9,
	0x2c, 0x4e, 0x95, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x2b, 0xa9, 0x72, 0x71,
	0xba, 0xe5, 0x24, 0x16, 0xa5, 0x06, 0xa5, 0x16, 0x17, 0x08, 0x49, 0x70, 0xb1, 0x27, 0xe7, 0xe7,
	0x95, 0xa4, 0xe6, 0x95, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x46, 0xae, 0x5c,
	0x6c, 0xce, 0x60, 0xe3, 0x84, 0xac, 0xb9, 0x58, 0xc1, 0x1a, 0x84, 0xa4, 0xf5, 0x60, 0x46, 0xea,
	0xb9, 0xa5, 0x26, 0x96, 0x94, 0x16, 0xa5, 0x7a, 0xe6, 0x16, 0xe4, 0x17, 0x95, 0x24, 0xe6, 0x25,
	0xa7, 0x4a, 0x09, 0x23, 0x49, 0xc2, 0x8c, 0x57, 0x62, 0x48, 0x62, 0x03, 0x5b, 0x6a, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xa2, 0xaa, 0x43, 0x07, 0x9c, 0x00, 0x00, 0x00,
}
```

在decode set时（`grpcurl.DescriptorSourceFromFileDescriptorSet(p.fileDescriptorSet)`），可能会因为缺少依赖报错，如：no descriptor found for "common.proto"，这需要将缺少的依赖加进来。最终可以得到完整的FileDescriptorSet。

## desc

```
desc -m helloworld.Greeter/SayHello helloworld extension google.protobuf
```

收集`helloworld.Greeter/SayHello`的依赖, 收集成功后，可以收集一份完整的helloworld.Greeter fileDescriptorSet。
收集目录：helloworld extension google.protobuf

日志如下：

```
imports: [helloworld extension google.protobuf]
INFO[0000] collect 3 pb.go
extension/extension.pb.go
helloworld/lang.pb.go
helloworld/greeter.pb.go
DEBU[0000] search helloworld.Greeter/SayHello, pkg: helloworld, service: helloworld.Greeter
DEBU[0000] searching helloworld/greeter.pb.go
INFO[0000] helloworld.Greeter/SayHello --> /Users/toukii/PATH/GOPATH/ezbuy/goflow/src/github.com/tgrpc/ngrpc/helloworld/greeter.pb.go
DEBU[0000] no descriptor found for "extension/extension.proto" ==> extension/extension.proto extension/extension.pb.go
DEBU[0000] searching extension/extension.pb.go
INFO[0000] extension/extension.proto --> /Users/toukii/PATH/GOPATH/ezbuy/goflow/src/github.com/tgrpc/ngrpc/extension/extension.pb.go
DEBU[0000] searching extension/extension.pb.go
INFO[0000] extension/extension.proto --> /Users/toukii/PATH/GOPATH/ezbuy/goflow/src/github.com/tgrpc/ngrpc/extension/github.com/tgrpc/ngrpc/extension/extension.pb.go
WARN[0000] extension/extension.pb.go is duplicated
DEBU[0000] no descriptor found for "google.protobuf/descriptor.proto" ==> google.protobuf/descriptor.proto google.protobuf/descriptor.pb.go
extension/extension.pb.go google.protobuf/descriptor.pb.go
helloworld/lang.pb.go google.protobuf/descriptor.pb.go
helloworld/greeter.pb.go google.protobuf/descriptor.pb.go
WARN[0000] not found pb for:helloworld.Greeter/SayHello
data writing ==> helloworld.Greeter.desc
```