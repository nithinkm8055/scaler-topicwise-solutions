package Assignment

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

func sortedArrayToBST(nums []int) *treeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, l, r int) *treeNode {
	if l > r {
		return nil
	}

	mid := (l + r) / 2
	x := treeNode_new(nums[mid])

	x.left = build(nums, l, mid-1)
	x.right = build(nums, mid+1, r)
	return x
}
