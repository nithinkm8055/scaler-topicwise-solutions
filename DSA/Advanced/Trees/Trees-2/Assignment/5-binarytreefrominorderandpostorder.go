package Assignment

func buildTreeFromInorderAndPostOrder(A []int, B []int) *treeNode {
	return constuctTreeFromInorderAndPostOrder(B, A, 0, len(A)-1, 0, len(B)-1)
}

func constuctTreeFromInorderAndPostOrder(pre []int, in []int, ps int, pe int, ins int, ine int) *treeNode {

	if ps > pe {
		return nil
	}

	c_root := pre[pe]
	x := treeNode_new(c_root)

	if ins == ine {
		return x
	}

	idx := -1
	for i := ins; i <= ine; i++ {
		if in[i] == c_root {
			idx = i
			break
		}
	}

	length := ine - idx
	x.left = constuctTreeFromInorderAndPostOrder(pre, in, ps, ps-ins+idx-1, ins, idx-1)
	x.right = constuctTreeFromInorderAndPostOrder(pre, in, pe-length, pe-1, idx+1, ine)

	return x

}
