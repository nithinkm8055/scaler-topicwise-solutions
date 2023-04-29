package Assignment

func reverseList(A *listNode) *listNode {

	curr := A
	var prev *listNode
	var temp *listNode

	for curr != nil {
		temp = curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}

	return prev
}
