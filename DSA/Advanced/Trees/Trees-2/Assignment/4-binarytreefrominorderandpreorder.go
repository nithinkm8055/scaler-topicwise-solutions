package Assignment

func buildTree(A []int, B []int) *treeNode {
	return constuctTree(A, B, 0, len(A)-1, 0, len(B)-1)
}

func constuctTree(pre []int, in []int, ps int, pe int, ins int, ine int) *treeNode {

	if ps > pe {
		return nil
	}

	c_root := pre[ps]
	x := treeNode_new(c_root)
	idx := -1
	for i := ins; i <= ine; i++ {
		if in[i] == c_root {
			idx = i
			break
		}
	}

	length := idx - ins
	x.left = constuctTree(pre, in, ps+1, ps+length, ins, idx-1)
	x.right = constuctTree(pre, in, ps+length+1, pe, idx+1, ine)

	return x

}
