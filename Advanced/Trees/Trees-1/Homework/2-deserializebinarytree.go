package Homework

func DeserializeBinaryTree(A []int) *treeNode {

	if len(A) == 0 {
		return nil
	}

	var head *Node = Node_new(treeNode_new(A[0]))
	var tail *Node = head

	var node *Node = head

	i := 1
	for i < len(A) {

		if A[i] != -1 {
			temp1 := treeNode_new(A[i])
			node.val.left = temp1
			tail.next = Node_new(temp1)
			tail = tail.next
		}
		i++

		if i < len(A) && A[i] != -1 {
			temp2 := treeNode_new(A[i])
			node.val.right = temp2
			tail.next = Node_new(temp2)
			tail = tail.next

		}
		i++
		node = node.next
	}

	return head.val
}
