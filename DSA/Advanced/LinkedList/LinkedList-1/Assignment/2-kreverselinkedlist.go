package Assignment

func KReverseLinkedList(A *listNode, B int) *listNode {

	if A == nil {
		return nil
	}

	curr := A
	var prev *listNode
	var temp *listNode

	k1 := B

	for curr != nil && B > 0 {

		temp = curr.next
		curr.next = prev
		prev = curr
		curr = temp
		B--
	}

	A.next = KReverseLinkedList(curr, k1)
	return prev
}
