

```go
package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

// 调用方法：km(arr,0,len(arr))
func km(arr []int, k, m int) {
	// fmt.Println(arr)
	if k > m {
		for i := 0; i <= m; i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Println()
	} else {
		for i := k; i <= m; i++ {
			swap(&arr[k], &arr[i])
			km(arr, k+1, m)
			swap(&arr[k], &arr[i])
		}
	}
}

func main() {
	arr := []int{1, 2, 3}
	km(arr, 0, len(arr)-1)
}

```


## DeduplicationId

```
// 去重，并保证顺序
// new cap is equal the old len : 新的slice的cap等于旧slice的len
func DeduplicationId(ids []int64) []int64 {
	idMap := make(map[int64]bool)
	var index int
	for i, id := range ids {
		if _, ex := idMap[id]; !ex {
			idMap[id] = true
			if i != index {
				ids[index] = id
			}
			index++
		}
	}
	return ids[:index]
}
```