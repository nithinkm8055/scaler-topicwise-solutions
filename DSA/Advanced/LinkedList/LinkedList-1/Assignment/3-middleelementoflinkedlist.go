package Assignment

//
//func main() {
//	node1 := listNode_new(46)
//	node2 := listNode_new(76)
//	node3 := listNode_new(35)
//
//	node1.next = node2
//	node2.next = node3
//
//	fmt.Println(solve(node1))
//}
//
//type listNode struct {
//	value int
//	next  *listNode
//}
//
//func listNode_new(val int) *listNode {
//	var node *listNode = new(listNode)
//	node.value = val
//	node.next = nil
//	return node
//}

// slow and fast pointers
func MiddleElementOfLinkedList(A *listNode) int {

	if A == nil {
		return 0
	}

	slow := A
	fast := A

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	if fast.next == nil {
		return slow.value
	}

	return slow.next.value
}

func MiddleElementOfLinkedList2(A *listNode) int {
	if A == nil {
		return 0
	}

	if A.next == nil {
		return A.value
	}

	count := 0
	temp := A
	for temp != nil {
		count++
		temp = temp.next
	}

	middle := count / 2

	temp2 := A
	count = 0
	for temp2 != nil {

		if count == middle {
			return temp2.value
		}
		count++
		temp2 = temp2.next
	}

	return 0
}
