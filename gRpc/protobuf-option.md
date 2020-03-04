# protobuf扩展

protobuf 使用option扩展，option有如下几种类型：

```
FileOptions
ServiceOptions
MethodOptions
MessageOptions
FieldOptions
EnumOptions
EnumValueOptions
```

这些option定义在 google.protobuf/descriptor.proto 中。

利用option可以增加一些自定义的选项，比如：
使用ServiceOptions控制Service是否开放grpc、http；
使用MethodOptions控制是否为接口增加缓存，增加鉴权控制等。
这些自定义选项很活泛，可以满足针对不同级别的扩展。

我们以ServiceOptions为例，实现针对Service的扩展：http服务路由路径前缀(一般路由前缀已经是确定的并在代码里hard code的，不需要这种扩展，只是举个例子)，定义所有服务的路由前缀是：`/rest/v1/api`


## 服务扩展：


descriptor代码使用gogo库。

google.protobuf/descriptor.proto

```
syntax = "proto2";

package google.protobuf;
option go_package = "github.com/gogo/protobuf/protoc-gen-gogo/descriptor;descriptor";

message ServiceOptions {
	...
}
```

extension.proto

```
syntax = "proto3";
package extension;

option go_package = "github.com/tgrpc/ngrpc/helloworld/rpc/extension";

import "google.protobuf/descriptor.proto";

extend google.protobuf.ServiceOptions {
    HttpRestOption rest = 20200229;
}

message HttpRestOption {
	// 路由前缀
	string preffix = 1;
}
```

helloworld/greeter.proto


```
syntax = "proto3";
import "extension/extension.proto";
package helloworld;

service Greeter {
  option (extension.rest).preffix = "/rest/v1/api";
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}
```

生成扩展代码：

```
protoc -I . --gogo_out=plugins=grpc:extension extension/extension.proto
protoc -I . --gogo_out=plugins=grpc:. helloworld/greeter.proto
```
## 验证扩展

```
package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	ngrpc "github.com/tgrpc/ngrpc/extension"
	"github.com/tgrpc/desc"
)

var (
	rawDescs = []string{"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"}
)

func main() {
	fileset, err := desc.DecodeGoGoFileDescriptorSetByBase64Str("", rawDescs)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	for _, file := range fileset.File {
		fmt.Println(file.GetName())
		if file.GetName() != "helloworld/greeter.proto" {
			continue
		}
		for _, service := range file.Service {
			fmt.Println("service.Name: ", service.GetName())
			os1, err := proto.GetExtension(service.GetOptions(), ngrpc.E_Rest)
			fmt.Println(os1, err)
			opt := os1.(*ngrpc.HttpRestOption)
			fmt.Printf("option: %+v\n", opt)
		}
	}
}
```

输出结果：

```
helloworld/greeter.proto
service.Name:  Greeter
option: preffix:"/rest/v1/api"
```

结果取到了rest.preffix为protobuf里定义的内容。在正常的业务中，一般用于生成代码。
extension信息都可以在`file *generator.FileDescriptor`中得到，验证程序是通过desc工具得到的。
