package Assignment

type ListNode struct {
	value int
	next  *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	fast := head
	slow := head

	var interNode *ListNode

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			interNode = fast
			break
		}
	}

	if interNode != nil {

		slow = head

		for {

			if slow == fast {
				return slow
			}
			slow = slow.next
			fast = fast.next
		}
	}

	return nil
}
