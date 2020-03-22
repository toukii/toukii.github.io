# 二分查找

很好的一篇文章，讲述了如何确定[二分查找的边界 (转载)](https://www.cnblogs.com/kyoner/p/11080078.html)。

## 二分查找中的问题

二分mid溢出：low+high可能超出了int的最大长度，安全一点的写法：`mid := low + (high-low)>>1`。
go官方的写法：`mid := int(uint(low+hith)>>1)`。
要用一个表示范围更广的数字类型保存中间结果，除以2后再转为原来的数字类型。

go sort包的二分查找：

```
func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
```

查找的规律：将已有序列划分为左右两部分，使得左部分f[:i]都是false：f(x); 右部分f[i:]都是true。

如在数组 `[2 2 3 3 4 6 7 8]` 中查找3的最左侧位置，f为：
```
func(i int) bool {
	data[i]>= 3
}
```

若查找最右侧位置，f为：
```
func(i int) bool {
	data[i]> 3
}
```

查找并插入：

```
idx := sort.Search(len(data), func(i int) bool {
	return data[i] > x
})
data = append(data, x)
copy(data[idx+1:], data[idx:]) // copy(dst, src), on the way 切片操作是左开右闭[k:b)
data[idx] = x
```

## 搜索区间

`high = len(data)`
搜索区间是 [low, high), for条件是 `low<high`

最左查找：
```
data[m] == x, 搜索范围是 [low,m), high = m；
data[m] > x, 搜索范围是 [low, m-1], high = m
data[m] < x, 搜索范围是 [m+1, high), low = m+1
```

最右查找：
```
data[m] == x, 搜索范围是 [m+1,high), low = m+1；
data[m] > x, 搜索范围是 [low, m-1], high = m
data[m] < x, 搜索范围是 [m+1, high), low = m+1
```

```
func BinarySearchLeft(data []int, x int) int {
	low, high := 0, len(data)
	for low < high {
		m := (low + high) >> 1
		if data[m] == x {
			high = m
		} else if data[m] > x {
			high = m
		} else if data[m] < x {
			low = m + 1
		}
	}
	fmt.Println(x, low, high)
	return low
}

func BinarySearchRight(data []int, x int) int {
	low, high := 0, len(data)
	for low < high {
		m := (low + high) >> 1
		if data[m] == x {
			low = m + 1
		} else if data[m] > x {
			high = m
		} else if data[m] < x {
			low = m + 1
		}
	}
	fmt.Println(x, low, high)
	return low
}
```

第一个，最基本的二分查找算法：

```
因为我们初始化 right = nums.length - 1
所以决定了我们的「搜索区间」是 [left, right]
所以决定了 while (left <= right)
同时也决定了 left = mid+1 和 right = mid-1

因为我们只需找到一个 target 的索引即可
所以当 nums[mid] == target 时可以立即返回
```

第二个，寻找左侧边界的二分查找：

```
因为我们初始化 right = nums.length
所以决定了我们的「搜索区间」是 [left, right)
所以决定了 while (left < right)
同时也决定了 left = mid+1 和 right = mid

因为我们需找到 target 的最左侧索引
所以当 nums[mid] == target 时不要立即返回
而要收紧右侧边界以锁定左侧边界
```

第三个，寻找右侧边界的二分查找：

```
因为我们初始化 right = nums.length
所以决定了我们的「搜索区间」是 [left, right)
所以决定了 while (left < right)
同时也决定了 left = mid+1 和 right = mid

因为我们需找到 target 的最右侧索引
所以当 nums[mid] == target 时不要立即返回
而要收紧左侧边界以锁定右侧边界

又因为收紧左侧边界时必须 left = mid + 1
所以最后无论返回 left 还是 right，必须减一
```
