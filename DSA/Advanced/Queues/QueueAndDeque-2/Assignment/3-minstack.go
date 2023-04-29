package Assignment

type Node struct {
	val  int
	next *Node
}

type MinStack struct {
	minhead *Node // pointer to min stack
	head    *Node // pointer to value stack
}

/** initialize your data structure here. */
func Constructor() *MinStack {
	minStack := new(MinStack)

	dummyHead := new(Node)
	dummyHead.next = nil

	dummyMinHead := new(Node)
	dummyMinHead.next = nil

	minStack.head = dummyHead
	minStack.minhead = dummyMinHead

	return minStack
}

func newNode(val int) *Node {
	node := new(Node)
	node.val = val
	node.next = nil
	return node
}

func (this *MinStack) Push(val int) {
	nNode := newNode(val)
	nNode.next = this.head.next
	this.head.next = nNode

	if this.minhead.next != nil {
		var minNode *Node
		if this.minhead.next.val >= val {
			minNode = newNode(val)
		} else {
			minNode = newNode(this.minhead.next.val)
		}
		minNode.next = this.minhead.next
		this.minhead.next = minNode
	} else {
		this.minhead.next = nNode
	}
}

func (this *MinStack) Pop() {
	if this.head.next != nil {
		this.head.next = this.head.next.next
		this.minhead.next = this.minhead.next.next
	}
}

func (this *MinStack) Top() int {
	if this.head.next != nil {
		return this.head.next.val
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.minhead.next != nil {
		return this.minhead.next.val
	}
	return -1
}
