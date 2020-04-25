package LRU_cache

type ListNode struct {
	Key  int
	Val  int
	Pre  *ListNode
	Next *ListNode
}

type LRUCache struct {
	farthest *ListNode
	last     *ListNode
	m        map[int]*ListNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		farthest: &ListNode{},
		last:     &ListNode{},
		m:        make(map[int]*ListNode),
		capacity: capacity,
	}
	cache.farthest.Next = cache.last
	cache.last.Pre = cache.farthest
	return cache
}

// 裁出节点
func (this *LRUCache) cut(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) moveToLast(node *ListNode) {
	this.last.Pre.Next = node
	node.Pre = this.last.Pre
	node.Next = this.last
	this.last.Pre = node
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.m[key]
	if !ok {
		return -1
	}
	if this.last.Pre != v {
		this.cut(v)
		this.moveToLast(v)
	}
	return this.m[key].Val
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		this.m[key].Val = value // 更新值
		if this.last.Pre != v {
			this.cut(v)
			this.moveToLast(v)
		}
	} else {
		node := &ListNode{Key: key, Val: value}
		this.moveToLast(node)
		if len(this.m) == this.capacity {
			delete(this.m, this.farthest.Next.Key)
			this.cut(this.farthest.Next)
		}
		this.m[key] = node
	}
}
