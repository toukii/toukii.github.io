# protobuf原理 [转载](https://blog.csdn.net/carson_ho/article/details/70568606)

## T - L - V 的数据存储方式
即 Tag - Length - Value，标识 - 长度 - 字段值 存储方式， 其中 Length可选存储，如 储存Varint编码数据就不需要存储Length。

### TLV存储的优点

不需要分隔符 就能 分隔开字段，减少了 分隔符 的使用
各字段 存储得非常紧凑，存储空间利用率非常高
若 字段没有被设置字段值，那么该字段在序列化时的数据中是完全不存在的，即不需要编码。


## 结论1： Protocol Buffer将 消息里的每个字段 进行编码后，再利用T - L - V 存储方式 进行数据的存储，最终得到的是一个 二进制字节流

protobuf编解码时，存储方式 Tag-Value-Tag-Value-...，Length（Varint编码无L）。

Tag：经过 Protocol Buffer采用Varint & Zigzag编码后 的消息字段 标识号 & 数据类型 的值, 用于标识字段。

```
存储了字段的标识号（field_number）和 数据类型（wire_type），即Tag = 字段数据类型（wire_type） + 标识号（field_number）
占用 一个字节 的长度（如果标识号超过了16，则占用多一个字节的位置）
解包时，Protocol Buffer根据 Tag 将 Value 对应于消息中的 字段。
```

Value: 经过 Protocol Buffer采用Varint & Zigzag编码后 的消息字段的值。

__由于其编码的特性，要求调用方参数的数据类型和标识号与服务端保持一致，若类型不一致，无法解码。__

## 结论2：Protocol Buffer对于不同数据类型 采用不同的 序列化方式（编码方式 & 数据存储方式）

```
对于存储Varint编码数据，就不需要存储字节长度 Length，所以实际上Protocol Buffer的存储方式是 T - V；
若Protocol Buffer采用其他编码方式（如LENGTH_DELIMITED）则采用T - L - V
```

## 结论3：因为 Protocol Buffer对于数据字段值的 独特编码方式 & T - L - V数据存储方式，使得 Protocol Buffer序列化后数据量体积如此小

## 使用建议

### 建议1：多用 optional或 repeated修饰符
因为若optional 或 repeated 字段没有被设置字段值，那么该字段在序列化时的数据中是完全不存在的，即不需要进行编码
相应的字段在解码时才会被设置为默认值

### 建议2：字段标识号（Field_Number）尽量只使用 1-15，且不要跳动使用
因为Tag里的Field_Number是需要占字节空间的。如果Field_Number>16时，Field_Number的编码就会占用2个字节，那么Tag在编码时也就会占用更多的字节；如果将字段标识号定义为连续递增的数值，将获得更好的编码和解码性能

### 建议3：若需要使用的字段值出现负数，请使用 sint32 / sint64，不要使用int32 / int64
因为采用sint32 / sint64数据类型表示负数时，会先采用Zigzag编码再采用Varint编码，从而更加有效压缩数据

### 建议4：对于repeated字段，尽量增加packed=true修饰
因为加了packed=true修饰repeated字段采用连续数据存储方式，即T - L - V - V -V方式
