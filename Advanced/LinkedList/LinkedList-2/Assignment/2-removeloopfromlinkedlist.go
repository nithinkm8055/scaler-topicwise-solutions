package Assignment

func RemoveLoopFromLinkedList(A *ListNode) *ListNode {
	slow := A
	fast := A

	var prev *ListNode

	// find intersection point
	for {
		prev = fast
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	slow = A

	for {

		if slow == fast {
			prev.next = nil
			break
		}
		prev = fast
		slow = slow.next
		fast = fast.next
	}

	return A
}
