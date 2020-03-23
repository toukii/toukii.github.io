# 排序map

## 对key排序

```
import (
	"sort"
)

type SortedKeyMap struct {
	M          map[int]string
	SortedKeys []int
}

func InitSortedKeyMap(m map[int]string) *SortedKeyMap {
	sm := &SortedKeyMap{M: m}
	sks := make([]int, 0, len(m))
	for k, _ := range m {
		sks = append(sks, k)
	}
	sort.Slice(sks, func(i, j int) bool {
		return sks[i] < sks[j]
	})
	sm.SortedKeys = sks
	return sm
}
```

## 链表排序

可以支持二分查找(O(2logN))
快速插入(O(1))

```
type SortedLinkedList struct {
	M          map[int]*Node
	SortedKeys []int
}

func InitSortedLinkedList(m map[int]*Node) *SortedLinkedList {
	sm := &SortedLinkedList{M: m}
	sks := make([]int, 0, len(m))
	for k, _ := range m {
		sks = append(sks, k)
	}
	sort.Slice(sks, func(i, j int) bool {
		return sks[i] < sks[j]
	})
	sm.SortedKeys = sks
	return sm
}

```
