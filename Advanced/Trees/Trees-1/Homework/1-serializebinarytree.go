package Homework

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

type Node struct {
	val  *treeNode
	next *Node
}

func Node_new(val *treeNode) *Node {
	node := new(Node)
	node.next = nil
	node.val = val
	return node
}

func solve(A *treeNode) []int {
	var head *Node = Node_new(A)
	var tail *Node = head

	result := make([]int, 0)
	result = append(result, A.value)

	for head != nil {

		if head.val.left != nil {
			result = append(result, head.val.left.value)
			tail.next = Node_new(head.val.left)
			tail = tail.next
		} else {
			result = append(result, -1)
		}

		if head.val.right != nil {
			result = append(result, head.val.right.value)
			tail.next = Node_new(head.val.right)
			tail = tail.next
		} else {
			result = append(result, -1)
		}

		head = head.next
	}

	return result

}
