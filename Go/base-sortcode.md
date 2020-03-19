# 排序算法实现

## 插入排序

## 依次移动数组元素

```
func SortInsert(data []int) {
	size := len(data)
	for idx := 1; idx < size; idx++ {
		insert(data[:idx+1], idx)
	}
}

// 插入排序: 依次移动数组元素
func insert(data []int, idx int) {
	for idx > 0 && data[idx] < data[idx-1] {
		data[idx-1], data[idx] = data[idx], data[idx-1]
		idx--
	}
}
```

### 批量移动数组元素

```
func SortInsert2(data []int) {
	size := len(data)
	for idx := 1; idx < size; idx++ {
		insert2(data[:idx+1], idx)
	}
}

// 插入排序：批量移动数组元素
func insert2(data []int, idx int) {
	index := idx
	for i := idx - 1; i >= 0; i-- {
		if data[i] >= data[idx] {
			index = i
			continue
		} else {
			break
		}
	}
	v := data[idx]
	fmt.Println(idx, v, index, data)
	copy(data[index+1:idx+1], data[index:idx])
	data[index] = v
}
```

## 选择排序

```

// 选择排序：每次选择最大或最小的元素
func SortSelect(data []int) {
	size := len(data)
	for idx := 0; idx < size; idx++ {
		min := idx
		for j := idx + 1; j < size; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		if min != idx {
			data[idx], data[min] = data[min], data[idx]
		}
	}
}
```

## 冒泡排序

```
// 冒泡排序：比较相邻两个元素的大小，若不符合排序规则，则交互元素
func SortBubble(data []int) {
	swaped := false
	size := len(data)
	for i := 0; i < size; i++ {
		for j := 1; j < size; j++ {
			if !compare(data[j-1], data[j]) {
				data[j-1], data[j] = data[j], data[j-1]
				swaped = true
			}
		}
		if !swaped {
			break
		}
	}
}

func compare(i, j int) bool {
	return i < j
}
```

## 快速排序

### 取data[low]未pivot

```
func SortQuick(data []int) {
	// size := len(data)
	pivot := quick(data)
	// pivot := quickv2(data)
	if pivot < 0 {
		return
	}
	SortQuick(data[:pivot])
	SortQuick(data[pivot+1:])
}

// 取data[low]为pivot
func quick(data []int) int {
	size := len(data)
	low, high := 0, size-1
	if low >= high {
		return -1
	}
	pivot := data[low]
	for low < high {
		for low < high && data[high] >= pivot {
			high--
		}
		data[low] = data[high]
		for low < high && data[low] <= pivot {
			low++
		}
		data[high] = data[low]
	}
	data[low] = pivot
	return low
}
```

### 随机取或取中间数为pivot

```
func quickv2(data []int) int {
	size := len(data)
	if size <= 1 {
		return -1
	}
	low, high := 0, size-1
	pivot := (low + high) >> 1
	tmp := data[pivot]
	for low < high {
		for pivot < high && data[high] >= tmp {
			high--
		}
		data[pivot] = data[high]
		pivot = high

		for low < pivot && data[low] <= tmp {
			low++
		}
		data[pivot] = data[low]
		pivot = low
	}
	data[pivot] = tmp
	return pivot
}
```

### 傻瓜快排

```
func QuickSortv3(data []int) int {
	size := len(data)
	if size <= 1 {
		return -1
	}
	fmt.Println(data)
	low, high := 0, size-1
	pivot := (low + high) >> 1
	tmp := data[pivot]
	left, equal, right := make([]int, 0, size), make([]int, 0, size>>2+1), make([]int, 0, size)

	for i := 0; i < size; i++ {
		if data[i] < tmp {
			left = append(left, data[i])
		} else if data[i] > tmp {
			right = append(right, data[i])
		} else {
			equal = append(equal, data[i])
		}
	}

	QuickSortv3(left)
	idx := 0
	for _, it := range left {
		data[idx] = it
		idx++
	}

	for _, it := range equal {
		data[idx] = it
		idx++
	}

	QuickSortv3(right)
	for _, it := range right {
		data[idx] = it
		idx++
	}
	return pivot
}
```
