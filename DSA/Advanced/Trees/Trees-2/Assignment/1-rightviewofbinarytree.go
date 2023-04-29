package Assignment

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

func RightViewBinaryTree(A *treeNode) []int {

	// construct the level order traversal for the tree

	var head *Node = Node_new(A)
	var tail *Node = Node_new(treeNode_new(-1))

	head.next = tail

	result := make([]int, 0)
	result = append(result, head.val.value)
	result = append(result, -1)

	for head != nil {
		if head.val.value == -1 && head.next == nil {
			break
		}

		if head.val.value == -1 {
			tail.next = Node_new(treeNode_new(-1))
			tail = tail.next
			head = head.next
			result = append(result, -1)
			continue
		}

		if head.val.left != nil {
			tail.next = Node_new(head.val.left)
			tail = tail.next
			result = append(result, head.val.left.value)
		}

		if head.val.right != nil {
			tail.next = Node_new(head.val.right)
			tail = tail.next
			result = append(result, head.val.right.value)
		}

		head = head.next
	}

	ans := make([]int, 0)
	// ans = append(ans, result[len(result-1)])

	for i := 0; i < len(result)-1; i++ {
		if result[i+1] == -1 {
			ans = append(ans, result[i])
		}
	}

	return ans
}
