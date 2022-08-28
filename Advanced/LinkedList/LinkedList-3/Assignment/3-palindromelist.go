package Assignment

type listNode struct {
	value int
	next  *listNode
}

func lPalin(A *listNode) int {

	middle := findMiddle(A)

	h2 := reverse(middle.next)
	middle.next = nil

	temp := A

	for temp != nil && h2 != nil {
		if temp.value != h2.value {
			return 0
		}
		temp = temp.next
		h2 = h2.next
	}

	return 1
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

	return slow
}
