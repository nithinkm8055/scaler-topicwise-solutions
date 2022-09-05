package Assignment

//type Node struct {
//	next *Node
//	val  int
//}
//
//func newNode(val int) *Node {
//	node := new(Node)
//	node.val = val
//	node.next = nil
//	return node
//}

func largestRectangleArea(A []int) int {

	max := 0

	lsIndex := findLeftSmallerIndex(A)
	rsIndex := findRightSmallerIndex(A)

	for i := range A {
		area := (rsIndex[i] - lsIndex[i] - 1) * A[i]

		if area > max {
			max = area
		}
	}

	return max
}

func findLeftSmallerIndex(A []int) []int {
	var lHead *Node
	lsIndex := make([]int, 0)
	for i := range A {
		if lHead == nil {
			lHead = insert(lHead, i)
			lsIndex = append(lsIndex, -1)
		} else if A[lHead.val] >= A[i] {
			for lHead != nil {
				if A[lHead.val] < A[i] {
					lsIndex = append(lsIndex, lHead.val)
					break
				}
				lHead = lHead.next
			}
			if lHead == nil {
				lsIndex = append(lsIndex, -1)
			}
			lHead = insert(lHead, i)
		} else if A[lHead.val] < A[i] {
			lsIndex = append(lsIndex, lHead.val)
			lHead = insert(lHead, i)
		}
	}
	return lsIndex
}

func findRightSmallerIndex(A []int) []int {
	var rHead *Node
	rsIndex := make([]int, len(A))
	for j := len(A) - 1; j >= 0; j-- {
		if rHead == nil {
			rHead = insert(rHead, j)
			rsIndex[j] = len(A)
		} else if A[rHead.val] >= A[j] {
			for rHead != nil {
				if A[rHead.val] < A[j] {
					rsIndex[j] = rHead.val
					break
				}
				rHead = rHead.next
			}
			if rHead == nil {
				rsIndex[j] = len(A)
			}
			rHead = insert(rHead, j)
		} else if A[rHead.val] < A[j] {
			rsIndex[j] = rHead.val
			rHead = insert(rHead, j)
		}
	}
	return rsIndex
}

//func insert(head *Node, val int) *Node {
//	node := newNode(val)
//	node.next = head
//	head = node
//	return head
//}
