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

func mergeTwoLists(A *listNode, B *listNode) *listNode {

	h3 := listNode_new(-1)
	tail := h3

	h1 := A
	h2 := B

	for h1 != nil && h2 != nil {
		if h1.value <= h2.value {
			tail.next = h1
			h1 = h1.next
			tail = tail.next
		} else {
			tail.next = h2
			h2 = h2.next
			tail = tail.next
		}
	}

	if h1 == nil {
		tail.next = h2
	} else {
		tail.next = h1
	}

	return h3.next
}
