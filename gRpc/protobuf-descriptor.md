# protobuf反射原理 [转载](https://www.cnblogs.com/swey/p/4733638.html)

## 获取message和service的属性和方法

protobuf通过Descriptor获取任意message或service的属性和方法，Descriptor主要包括了以下集中类型：
 
```
FileDescriptor 　　获取proto文件中的Descriptor和ServiceDescriptor
Descriptor 　　　  获取类message属性和方法，包括FieldDescriptor 和 EnumDescriptor等
FieldDescriptor 　 获取message中各个字段的类型、标签、名称等
EnumDescriptor   获取Enum中的各个字段名称、值等
ServiceDescriptor 获取service中的MethodDescriptor
MethodDescriptor 获取各个rpc中的request、response、名称等
```

当我们获得proto文件的FileDescriptor时，我们就可以获得所有的service的Descriptor和service的ServiceDescriptor，进而获取其相应的字段或rpc。也就是说，如果能获取到proto文件的FileDescriptor，就能所有的proto文件中的所有内容。

以go为例，使用protoc活动proto文件的Descriptor，基于该fileDescriptor可以调用服务端接口（参考tgrpc工具）。

## 获取fileDescriptor 方法1

```
protoc -I helloworld --gogo_out=plugins=grpc:rpc/helloworld helloworld/*.proto
```

greeter.pb.go 文件末尾有Greeter服务的fileDescriptor。

```
var fileDescriptorGreeter = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x48, 0xcd, 0xc9, 0xc9,
	0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x51, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0xf1, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x25, 0x35, 0x2e, 0x2e, 0xa8, 0x9a, 0x82, 0x9c, 0x4a, 0x21, 0x09, 0x2e,
	0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x22, 0x18, 0xd7, 0xa8, 0x97, 0x91, 0x8b, 0xdd,
	0x1d, 0x62, 0x93, 0x90, 0x1d, 0x17, 0x47, 0x70, 0x62, 0x25, 0x58, 0x9b, 0x90, 0x84, 0x1e, 0xc2,
	0x42, 0x3d, 0x64, 0xdb, 0xa4, 0xc4, 0xb0, 0xc8, 0x14, 0xe4, 0x54, 0x2a, 0x31, 0x08, 0x39, 0x71,
	0x71, 0xc1, 0xf4, 0x87, 0x19, 0x91, 0x63, 0x82, 0x01, 0xa3, 0x93, 0x01, 0x97, 0x74, 0x66, 0xbe,
	0x5e, 0x7a, 0x51, 0x41, 0xb2, 0x5e, 0x6a, 0x45, 0x62, 0x6e, 0x41, 0x4e, 0x6a, 0x31, 0x92, 0x6a,
	0x27, 0x7e, 0xb0, 0xf2, 0x70, 0x10, 0x3b, 0x00, 0x14, 0x2e, 0x01, 0x8c, 0x49, 0x6c, 0xe0, 0x00,
	0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xb6, 0xc9, 0xe0, 0x31, 0x01, 0x00, 0x00,
}
```

## 获取fileDescriptor 方法2

```
protoc --proto_path=helloworld --descriptor_set_out=helloworld.Greeter.pbin --include_imports helloworld/*.proto
```

helloworld.Greeter.pbin文件内容：

```
0ab1 020a 0d67 7265 6574 6572 2e70 726f
746f 120a 6865 6c6c 6f77 6f72 6c64 2222

....

726c 642e 4765 744c 616e 6752 6571 1a10
2e68 656c 6c6f 776f 726c 642e 4c61 6e67
6206 7072 6f74 6f33 
```

## Go获取proto描述

```
import "github.com/golang/protobuf/protoc-gen-go/descriptor"
```

```
type FileDescriptorSet struct {
	File                 []*FileDescriptorProto `protobuf:"bytes,1,rep,name=file" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

// Describes a complete .proto file.
type FileDescriptorProto struct {
	Name    *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Package *string `protobuf:"bytes,2,opt,name=package" json:"package,omitempty"`
	// Names of files imported by this file.
	Dependency []string `protobuf:"bytes,3,rep,name=dependency" json:"dependency,omitempty"`
	// Indexes of the public imported files in the dependency list above.
	PublicDependency []int32 `protobuf:"varint,10,rep,name=public_dependency,json=publicDependency" json:"public_dependency,omitempty"`
	// Indexes of the weak imported files in the dependency list.
	// For Google-internal migration only. Do not use.
	WeakDependency []int32 `protobuf:"varint,11,rep,name=weak_dependency,json=weakDependency" json:"weak_dependency,omitempty"`
	// All top-level definitions in this file.
	MessageType []*DescriptorProto        `protobuf:"bytes,4,rep,name=message_type,json=messageType" json:"message_type,omitempty"`
	EnumType    []*EnumDescriptorProto    `protobuf:"bytes,5,rep,name=enum_type,json=enumType" json:"enum_type,omitempty"`
	Service     []*ServiceDescriptorProto `protobuf:"bytes,6,rep,name=service" json:"service,omitempty"`
	Extension   []*FieldDescriptorProto   `protobuf:"bytes,7,rep,name=extension" json:"extension,omitempty"`
	Options     *FileOptions              `protobuf:"bytes,8,opt,name=options" json:"options,omitempty"`
	// This field contains optional information about the original source code.
	// You may safely remove this entire field without harming runtime
	// functionality of the descriptors -- the information is needed only by
	// development tools.
	SourceCodeInfo *SourceCodeInfo `protobuf:"bytes,9,opt,name=source_code_info,json=sourceCodeInfo" json:"source_code_info,omitempty"`
	// The syntax of the proto file.
	// The supported values are "proto2" and "proto3".
	Syntax               *string  `protobuf:"bytes,12,opt,name=syntax" json:"syntax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

descriptor.FileDescriptorSet 描述了一个FileDescriptorProto的集合，包含了它自身的定义和一些依赖定义。
FileDescriptorProto 描述了单个FileDescProto的属性：名称、服务、定义等。

获取descriptor.FileDescriptorSet：

```
data, err := ioutil.ReadFile("helloworld.Greeter.pbin")
if err != nil {
	return nil, err
}
fileDescriptorSet := &descriptor.FileDescriptorSet{}
if err := proto.Unmarshal(data, fileDescriptorSet); err != nil {
	return nil, err
}
```

FileDescriptorProto检查：
fileDescriptorProto: 包含当前调用服务的fileDescriptorProto， fileDescriptorSet: 依赖的fileDescriptorProto集合。

```
func SortFileDescriptorSet(fileDescriptorSet *descriptor.FileDescriptorSet, fileDescriptorProto *descriptor.FileDescriptorProto) (*descriptor.FileDescriptorSet, error) {
	// best-effort checks
	names := make(map[string]struct{}, len(fileDescriptorSet.File))
	for _, iFileDescriptorProto := range fileDescriptorSet.File {
		if iFileDescriptorProto.GetName() == "" {
			return nil, fmt.Errorf("no name on FileDescriptorProto")
		}
		if _, ok := names[iFileDescriptorProto.GetName()]; ok {
			return nil, fmt.Errorf("duplicate FileDescriptorProto in FileDescriptorSet: %s", iFileDescriptorProto.GetName())
		}
		names[iFileDescriptorProto.GetName()] = struct{}{}
	}
	if _, ok := names[fileDescriptorProto.GetName()]; !ok {
		return nil, fmt.Errorf("no FileDescriptorProto named %s in FileDescriptorSet with names %v", fileDescriptorProto.GetName(), names)
	}
	newFileDescriptorSet := &descriptor.FileDescriptorSet{}
	for _, iFileDescriptorProto := range fileDescriptorSet.File {
		if iFileDescriptorProto.GetName() != fileDescriptorProto.GetName() {
			newFileDescriptorSet.File = append(newFileDescriptorSet.File, iFileDescriptorProto)
		}
	}
	newFileDescriptorSet.File = append(newFileDescriptorSet.File, fileDescriptorProto)
	return newFileDescriptorSet, nil
}
```

最终将fileDescriptorSet包装为grpcurl的

```
source := grpcurl.DescriptorSourceFromFileDescriptorSet(fileDescriptorSet)
```

```
import "github.com/tgrpc/grpcurl"

// DescriptorSource is a source of protobuf descriptor information. It can be backed by a FileDescriptorSet
// proto (like a file generated by protoc) or a remote server that supports the reflection API.
type DescriptorSource interface {
	// ListServices returns a list of fully-qualified service names. It will be all services in a set of
	// descriptor files or the set of all services exposed by a GRPC server.
	ListServices() ([]string, error)
	// FindSymbol returns a descriptor for the given fully-qualified symbol name.
	FindSymbol(fullyQualifiedName string) (desc.Descriptor, error)
	// AllExtensionsForType returns all known extension fields that extend the given message type name.
	AllExtensionsForType(typeName string) ([]*desc.FieldDescriptor, error)
}
```

最终，grpcurl调用grpc服务, source就是上面得到的DescriptorSource：

```
func InvokeRpc(ctx context.Context, source DescriptorSource, cc *grpc.ClientConn, methodName string,
	headers []string, handler InvocationEventHandler, requestData RequestMessageSupplier) error
```

