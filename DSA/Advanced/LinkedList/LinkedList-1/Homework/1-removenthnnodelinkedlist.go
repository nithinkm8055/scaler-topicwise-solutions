package Homework

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
type listNode struct {
	value int
	next  *listNode
}

//func listNode_new(val int) *listNode {
//	var node *listNode = new(listNode)
//	node.value = val
//	node.next = nil
//	return node
//}

func removeNthFromEnd(A *listNode, B int) *listNode {
	temp1 := A
	temp2 := A
	size := 0
	for temp1 != nil {
		temp1 = temp1.next
		size++
	}
	if B >= size {
		A = A.next
	} else {
		for i := 0; i < size-B-1; i++ {
			temp2 = temp2.next

		}
		temp2.next = temp2.next.next
	}
	return A
}
