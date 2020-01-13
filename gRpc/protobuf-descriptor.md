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
0a0c 4865 6c6c 6f52 6571 7565 7374 1212
0a04 6e61 6d65 1801 2001 2809 5204 6e61
6d65 2226 0a0a 4865 6c6c 6f52 6570 6c79
1218 0a07 6d65 7373 6167 6518 0120 0128
0952 076d 6573 7361 6765 328d 010a 0747
7265 6574 6572 123e 0a08 5361 7948 656c
6c6f 1218 2e68 656c 6c6f 776f 726c 642e
4865 6c6c 6f52 6571 7565 7374 1a16 2e68
656c 6c6f 776f 726c 642e 4865 6c6c 6f52
6570 6c79 2200 1242 0a0a 5361 7948 656c
6c6f 5632 1218 2e68 656c 6c6f 776f 726c
642e 4865 6c6c 6f52 6571 7565 7374 1a16
2e68 656c 6c6f 776f 726c 642e 4865 6c6c
6f52 6570 6c79 2200 3001 4230 0a1b 696f
2e67 7270 632e 6578 616d 706c 6573 2e68
656c 6c6f 776f 726c 6442 0f48 656c 6c6f
576f 726c 6450 726f 746f 5001 6206 7072
6f74 6f33 0ad1 040a 0a6c 616e 672e 7072
6f74 6f12 0a68 656c 6c6f 776f 726c 6422
370a 074c 6973 7452 6571 1216 0a06 6f66
6673 6574 1801 2001 2805 5206 6f66 6673
6574 1214 0a05 6c69 6d69 7418 0220 0128
0552 056c 696d 6974 2267 0a04 4c61 6e67
1212 0a04 6e61 6d65 1801 2001 2809 5204
6e61 6d65 121a 0a08 6269 7274 6864 6179
1802 2001 2803 5208 6269 7274 6864 6179
122f 0a08 7665 7273 696f 6e73 1803 2003
280b 3213 2e68 656c 6c6f 776f 726c 642e
5665 7273 696f 6e52 0876 6572 7369 6f6e
7322 9f01 0a07 5665 7273 696f 6e12 0e0a
0269 6418 0120 0128 0552 0269 6412 180a
0776 6572 7369 6f6e 1802 2001 2809 5207
7665 7273 696f 6e12 120a 0464 6573 6318
0320 0128 0952 0464 6573 6312 140a 0576
6933 3273 1804 2003 2805 5205 7669 3332
7312 140a 0576 7374 7273 1805 2003 2809
5205 7673 7472 7312 140a 0576 6936 3473
1806 2003 2803 5205 7669 3634 7312 140a
0576 6636 3473 1807 2003 2801 5205 7666
3634 7322 520a 084c 6973 7452 6573 7012
260a 056c 616e 6773 1801 2003 280b 3210
2e68 656c 6c6f 776f 726c 642e 4c61 6e67
5205 6c61 6e67 7312 1e0a 0a74 6f74 616c
436f 756e 7418 0220 0128 0552 0a74 6f74
616c 436f 756e 7422 200a 0a47 6574 4c61
6e67 5265 7112 120a 046e 616d 6518 0120
0128 0952 046e 616d 6532 750a 0b4c 616e
6753 6572 7669 6365 1231 0a04 4c69 7374
1213 2e68 656c 6c6f 776f 726c 642e 4c69
7374 5265 711a 142e 6865 6c6c 6f77 6f72
6c64 2e4c 6973 7452 6573 7012 330a 0747
6574 4c61 6e67 1216 2e68 656c 6c6f 776f
726c 642e 4765 744c 616e 6752 6571 1a10
2e68 656c 6c6f 776f 726c 642e 4c61 6e67
6206 7072 6f74 6f33 
```