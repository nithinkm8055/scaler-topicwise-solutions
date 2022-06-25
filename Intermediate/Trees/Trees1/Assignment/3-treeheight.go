package Assignment

//type TreeNode struct {
//	Left  *TreeNode
//	Value int
//	Right *TreeNode
//}

func TreeHeight(A *TreeNode) int {

	if A == nil {
		return 0
	}

	return 1 + max(TreeHeight(A.Left), TreeHeight(A.Right))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
