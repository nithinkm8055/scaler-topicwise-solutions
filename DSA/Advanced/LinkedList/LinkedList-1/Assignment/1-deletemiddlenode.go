package Assignment

type listNode struct {
	next  *listNode
	value int
}

func DeleteMiddleNode(A *listNode) *listNode {

	if A == nil || A.next == nil {
		return nil
	}

	temp := A

	count := 0 // stores length of LL

	for temp != nil {
		count++
		temp = temp.next
	}

	middle := count / 2

	count = 1
	temp = A
	for count < middle && temp != nil {
		count++
		temp = temp.next
	}

	temp.next = temp.next.next

	return A
}
