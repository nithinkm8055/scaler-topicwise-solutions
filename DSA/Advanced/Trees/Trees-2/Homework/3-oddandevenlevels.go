package Homework

func OddAndEvenLevels(A *treeNode) int {
	var head *Node = Node_new(A)
	var tail *Node = Node_new(treeNode_new(-1))

	head.next = tail

	result := make([]int, 0)
	result = append(result, A.value)
	result = append(result, -1)

	for head != nil {

		if head.val.value == -1 && head.next == nil {
			break
		}

		if head.val.value == -1 {
			tail.next = Node_new(treeNode_new(-1))
			tail = tail.next
			result = append(result, -1)
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

	sumEven := 0
	sumOdd := 0
	odd := false
	even := true
	for i := 0; i < len(result); i++ {

		if even && result[i] != -1 {
			sumEven += result[i]
		} else if even && result[i] == -1 {
			even = false
			odd = true
			continue
		}

		if odd && result[i] != -1 {
			sumOdd += result[i]
		} else if odd && result[i] == -1 {
			even = true
			odd = false
		}

	}

	return sumEven - sumOdd
}
