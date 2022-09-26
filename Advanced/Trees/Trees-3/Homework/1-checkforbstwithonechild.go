package Homework

// type treeNode struct {
//     value int
//     left *treeNode
//     right *treeNode
// }

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

func solve(A []int) string {

	root := treeNode_new(A[0])
	temp := root

	for i := 1; i < len(A); i++ {
		if A[i] <= temp.value {
			temp.left = treeNode_new(A[i])
			temp = temp.left
		} else {
			temp.right = treeNode_new(A[i])
			temp = temp.right
		}
	}

	var inorder func(*treeNode)
	result := make([]int, 0)

	inorder = func(root *treeNode) {
		if root == nil {
			return
		}

		inorder(root.left)
		result = append(result, root.value)
		inorder(root.right)
	}

	inorder(root)
	// fmt.Println(result)
	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			return "NO"
		}
	}
	return "YES"

}
