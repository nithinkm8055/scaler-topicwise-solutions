package Homework

type Node struct {
	prev *Node
	next *Node
	key  int
	val  int
}

func newNode(key int, val int) *Node {
	newNode := new(Node)
	newNode.key = key
	newNode.val = val
	newNode.next = nil
	newNode.prev = nil

	return newNode
}

type LRUCache struct {
	size     int
	capacity int
	head     *Node
	tail     *Node
}

var addrMap map[int]*Node

func Constructor(capacity int) *LRUCache {
	lruCache := new(LRUCache)
	lruCache.capacity = capacity
	lruCache.size = 0

	dummyHead := new(Node)
	dummyTail := new(Node)

	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead

	lruCache.head = dummyHead
	lruCache.tail = dummyTail

	addrMap = make(map[int]*Node)

	return lruCache
}

func (this *LRUCache) Get(key int) int {

	if _, ok := addrMap[key]; ok {
		node := addrMap[key]

		// removing node from current position
		this.remove(node)
		this.addNodeToTail(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	// if node already exists, just update the value and move position to tail
	if _, ok := addrMap[key]; ok {
		node := addrMap[key]
		node.val = value
		this.remove(node)
		this.addNodeToTail(node)
		return
	}

	newNode := newNode(key, value)
	if this.size == this.capacity {
		delete(addrMap, this.head.next.key)
		this.remove(this.head.next)
	}
	this.addNodeToTail(newNode)
	addrMap[key] = newNode
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	this.size--
}

func (this *LRUCache) addNodeToTail(node *Node) {
	node.next = this.tail
	node.prev = this.tail.prev
	this.tail.prev = node
	node.prev.next = node
	this.size++
}

//Your LRUCache object will be instantiated and called as such:
//obj := Constructor(capacity);
//param_1 := obj.Get(key);
//obj.Put(key,value);
