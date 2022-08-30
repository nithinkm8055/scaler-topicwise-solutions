package Homework

type listNode struct {
	value int
	next  *listNode
}

func listNode_new(val int) *listNode {
	var node *listNode = new(listNode)
	node.value = val
	node.next = nil
	return node
}

func partition(A *listNode, B int) *listNode {

	if A == nil || A.next == nil {
		return A
	}

	dummyNode := listNode_new(-1)
	tail := dummyNode

	clone := cloneList(A)

	count := 0
	count1 := 0
	temp := clone
	for temp != nil {
		if temp.value < B {
			tail.next = temp
			tail = tail.next
			count1++
		}
		temp = temp.next
		count++
	}

	tail.next = nil

	if count != count1 {
		temp = A
		for temp != nil {
			if temp.value >= B {
				tail.next = temp
				tail = tail.next
			}
			temp = temp.next
		}
	}
	tail.next = nil

	return dummyNode.next
}

func cloneList(A *listNode) *listNode {

	temp := A
	dummyNode := listNode_new(-1)
	tail := dummyNode

	for temp != nil {
		newNode := listNode_new(temp.value)
		tail.next = newNode
		tail = tail.next
		temp = temp.next
	}

	return dummyNode.next
}
