package Assignment

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

func TreeNode_new(val int) *TreeNode {
	var node *TreeNode = new(TreeNode)
	node.Value = val
	node.Left = nil
	node.Right = nil
	return node
}

func InorderTraversal(A *TreeNode) []int {

	var result []int
	var inorderTraversal func(*TreeNode)

	inorderTraversal = func(A *TreeNode) {
		if A == nil {
			return
		}

		inorderTraversal(A.Left)
		result = append(result, A.Value)
		inorderTraversal(A.Right)
	}

	inorderTraversal(A)

	return result

}

// Without Recursion

type Node struct {
	value *TreeNode
	next  *Node
}

var head *Node

func inorderTraversalwithoutRecursion(A *TreeNode) []int {

	result := make([]int, 0)
	temp := A

	for temp != nil || head != nil {

		for temp != nil {
			push(temp)
			temp = temp.Left
		}

		temp1 := pop()
		result = append(result, temp1.value.Value)
		temp = temp1.value.Right

	}

	return result
}

func push(A *TreeNode) {
	newNode := &Node{
		value: A,
	}

	newNode.next = head

	head = newNode
}

func pop() *Node {

	temp := head
	if head != nil {
		head = head.next
	}
	return temp
}
