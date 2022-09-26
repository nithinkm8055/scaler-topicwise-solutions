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
	node.val = val
	node.next = nil
	return node
}

func LeftViewBinaryTree(A *treeNode) []int {

	var head *Node = Node_new(A)

	temp := treeNode_new(-1)
	var tail *Node = Node_new(temp)

	head.next = tail

	result := make([]int, 0)
	result = append(result, A.value)

	for head != nil {

		if head.val.value == -1 && head.next == nil {
			break
		}

		if head.val.value == -1 {
			tail.next = Node_new(treeNode_new(-1))
			tail = tail.next
			result = append(result, head.next.val.value)
		}

		if head.val.left != nil {
			tail.next = Node_new(head.val.left)
			tail = tail.next
		}

		if head.val.right != nil {
			tail.next = Node_new(head.val.right)
			tail = tail.next
		}

		head = head.next
	}

	return result
}
