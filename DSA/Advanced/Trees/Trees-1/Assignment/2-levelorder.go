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
	next *Node
	val  *treeNode
}

func Node_new(val *treeNode) *Node {
	node := new(Node)
	node.next = nil
	node.val = val
	return node
}

func levelOrder(A *treeNode) [][]int {

	maxInt := 100000000
	var head *Node = Node_new(A)
	newTreeNode := treeNode_new(maxInt)
	var tail *Node = Node_new(newTreeNode)

	result := make([][]int, 0)
	head.next = tail

	result = append(result, []int{A.value})
	temp := make([]int, 0)
	for head != nil {

		if head.val.value == maxInt && head.next == nil {
			break
		}

		if head.val.value == maxInt {
			newTreeNode := treeNode_new(maxInt)
			tail.next = Node_new(newTreeNode)
			tail = tail.next
			head = head.next
			if len(temp) > 0 {
				result = append(result, temp)
			}
			temp = make([]int, 0)
			continue
		}

		if head.val.left != nil {
			temp = append(temp, head.val.left.value)
			tail.next = Node_new(head.val.left)
			tail = tail.next
		}

		if head.val.right != nil {
			temp = append(temp, head.val.right.value)
			tail.next = Node_new(head.val.right)
			tail = tail.next
		}
		head = head.next
	}

	return result
}
