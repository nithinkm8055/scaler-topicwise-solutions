package Homework

func zigzagLevelOrder(A *treeNode) [][]int {

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

	first := true
	second := false
	ans := make([][]int, 0)
	temp := make([]int, 0)
	for i := range result {

		if first {

			if result[i] != -1 {
				temp = append(temp, result[i])
			} else {
				first = false
				second = true
				ans = append(ans, temp)
				temp = make([]int, 0)
				continue
			}
		}

		if second {

			if result[i] != -1 {
				temp = append(temp, result[i])
			} else {
				first = true
				second = false
				ans = append(ans, reverse(temp))
				temp = make([]int, 0)
				continue
			}
		}
	}

	return ans
}

func reverse(temp []int) []int {

	for i := 0; i < len(temp)/2; i++ {
		temp[i], temp[len(temp)-i-1] = temp[len(temp)-i-1], temp[i]
	}

	return temp
}
