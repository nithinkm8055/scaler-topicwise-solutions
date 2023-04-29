package Assignment

type Pair struct {
	val *treeNode
	pos int
}

type QueueNode struct {
	val  Pair
	next *QueueNode
}

func QueueNode_new(val Pair) *QueueNode {
	node := new(QueueNode)
	node.val = val
	node.next = nil
	return node
}

func verticalOrderTraversal(A *treeNode) [][]int {

	var head *QueueNode = QueueNode_new(Pair{A, 0})
	var tail *QueueNode = head

	min := 0
	max := 0

	indexMap := make(map[int][]int) // key : level, value: slice of node values
	indexMap[0] = append(indexMap[0], A.value)

	for head != nil {

		if head.val.val.left != nil {
			tail.next = QueueNode_new(Pair{head.val.val.left, head.val.pos - 1})
			if min > head.val.pos-1 {
				min = head.val.pos - 1
			}
			if max < head.val.pos+1 {
				max = head.val.pos + 1
			}
			tail = tail.next
			indexMap[head.val.pos-1] = append(indexMap[head.val.pos-1], head.val.val.left.value)
		}

		if head.val.val.right != nil {
			tail.next = QueueNode_new(Pair{head.val.val.right, head.val.pos + 1})
			if max < head.val.pos+1 {
				max = head.val.pos + 1
			}
			if min > head.val.pos-1 {
				min = head.val.pos - 1
			}
			tail = tail.next
			indexMap[head.val.pos+1] = append(indexMap[head.val.pos+1], head.val.val.right.value)
		}

		head = head.next

	}

	result := make([][]int, 0)

	for i := min; i <= max; i++ {
		if len(indexMap[i]) > 0 {
			result = append(result, indexMap[i])
		}

	}

	return result
}
