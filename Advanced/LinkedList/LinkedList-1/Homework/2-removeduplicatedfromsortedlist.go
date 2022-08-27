package Homework

func deleteDuplicates(A *listNode) *listNode {

	curr := A
	var prev *listNode

	for curr != nil {

		if prev != nil && curr.value == prev.value {
			curr = prev.next.next
			prev.next = curr
			continue
		}
		prev = curr
		curr = curr.next

	}

	return A
}
