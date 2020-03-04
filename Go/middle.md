# go 

## 原理解读

[sync.Map](middle-syncMap.html)

## 实用技巧

[json 序列化tag](https://hanyajun.com/golang/int2sting/) 
	
	,string ,omitxxx 接口：MarshalJSON UnMarshalJSON

json 反序列化有个有趣的现象：{"123":{}}  可以用 map[string]interface{} 或 map[int64]interface{} 接收。


