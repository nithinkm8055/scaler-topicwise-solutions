package Assignment

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

func reorderList(A *listNode) *listNode {

	// find middle node of the linkedlist
	// reverse second part
	// merge in cyclic fashion

	mid := findMiddle(A)

	curr := mid.next
	mid.next = nil

	h2 := reverse(curr) // returns the head of the reversed linkedlist

	return merge(A, h2)
}

func merge(A *listNode, B *listNode) *listNode {

	h1 := A
	h2 := B

	var temp1 *listNode

	for h2 != nil {
		temp1 = h2.next
		h2.next = h1.next
		h1.next = h2
		h1 = h2.next
		h2 = temp1
	}
	return A
}

func reverse(A *listNode) *listNode {

	var prev *listNode
	var temp *listNode
	curr := A

	for curr != nil {
		temp = curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func findMiddle(A *listNode) *listNode {
	slow := A
	fast := A

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	if fast.next == nil {
		return slow
	}
	return slow.next
}
