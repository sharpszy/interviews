package middle

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	Key, Value int
	Prev, Next *DLinkNode
}

func initDLinkNode(k, v int) *DLinkNode {
	return &DLinkNode{
		Key:   k,
		Value: v,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLinkNode{},
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.Key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.Value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkNode) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) removeNode(node *DLinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}
