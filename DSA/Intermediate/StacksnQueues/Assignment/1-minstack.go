package Assignment

type MinStack struct {
	head *Node
}

type Node struct {
	value    int
	minValue int
	next     *Node
}

func Constructor() *MinStack {
	return new(MinStack)
}

func (this *MinStack) Push(val int) {

	newNode := &Node{
		value:    val,
		minValue: val,
	}

	newNode.next = this.head

	if this.head != nil {
		if this.head.minValue < newNode.minValue {
			newNode.minValue = this.head.minValue
		}
	}

	this.head = newNode

}

func (this *MinStack) Pop() {
	if this.head != nil {
		this.head = this.head.next
	}
}

func (this *MinStack) Top() int {
	if this.head != nil {
		return this.head.value
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.head != nil {
		return this.head.minValue
	}
	return -1
}

/**
* Your MinStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(val);
* obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.GetMin();
 */
