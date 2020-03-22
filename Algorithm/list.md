# 链表

链表节点定义

```
type Node struct {
	K    string
	V    int
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("[%s-%v] .. ", n.K, n.V)
}

func (n *Node) Discply() {
	node := n
	for node != nil {
		fmt.Print(node)
		node = node.next
	}
	fmt.Println()
}
```

## 链表反转

普通反转递归实现：

```
func ReverseV1(root *Node) *Node {
	if root.next == nil {
		return root
	}
	// node -> node(next) ->...-> tail
	// node <- node(next) <-...<- tail
	next := root.next
	root.next = nil
	tail := ReverseV1(next)
	next.next = root

	return tail
}
```

普通反转非递归实现：

```

func Reverse(root *Node) *Node {
	if root == nil || root.next == nil {
		return root
	}
	// node -> node(next) ->...-> tail
	// node <- node(next) <-...<- tail
	head := root
	next := root.next
	root.next = nil

	for next != nil {
		head.Discply()

		next.next = head

		head = next
		next = nn
	}

	return head
}
```

每n个节点反转一次，不足n个不反转，递归实现：

```
func LeastLen(root *Node, least int) bool {
	n := root
	for n != nil {
		n = n.next
		least--
		if least <= 0 {
			return true
		}
	}
	return false
}

func ReverseN(cur, head *Node, n, idx int) *Node {
	if cur == head && !LeastLen(cur, n) {
		return head
	}
	if cur.next == nil {
		return cur
	}
	if idx%n != 0 { // 正常的翻转
		// node -> node(next) ->...-> tail
		// node <- node(next) <-...<- tail
		next := cur.next
		cur.next = nil
		tail := ReverseN(next, head, n, idx+1)
		next.next = cur
		return tail
	}
	// 需要翻转
	// node -> node(next) ->...-> tail
	// node -> tail ->...-> node(next)
	next := cur.next
	tail := ReverseN(next, next, n, idx+1)
	head.next = tail // 当前链表翻转后的head指向下一个翻转的新头（tail）

	tail.Discply()
	return cur
}
```

```
[1-1] .. [2-2] .. [3-3] .. [4-4] .. [5-5] .. [6-6] .. [7-7] .. [8-8] ..
[7-7] .. [8-8] ..
[6-6] .. [5-5] .. [4-4] .. [7-7] .. [8-8] ..
[3-3] .. [2-2] .. [1-1] .. [6-6] .. [5-5] .. [4-4] .. [7-7] .. [8-8] ..
```

从末尾开始每3个翻转，不足的不翻转

```
l := Len(root)
offset := l % 3
r := root.next
pr := root
for i := 0; i < offset-1; i++ {
	r = r.next
	pr = pr.next
}
rn := ReverseN(r, r, 3, 1)
pr.next = rn
root.Discply()
```

```
[1-1] .. [2-2] .. [5-5] .. [4-4] .. [3-3] .. [8-8] .. [7-7] .. [6-6] ..
```