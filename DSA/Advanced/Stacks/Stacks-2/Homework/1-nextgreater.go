package Homework

type Node struct {
	next *Node
	val  int
}

func newNode(val int) *Node {
	node := new(Node)
	node.val = val
	node.next = nil
	return node
}

func nextGreater(A []int) []int {

	rgIndex := findRightGreaterIndex(A)

	ans := make([]int, 0)
	for i := range rgIndex {
		if rgIndex[i] != -1 {
			ans = append(ans, A[rgIndex[i]])
		} else {
			ans = append(ans, rgIndex[i])
		}
	}

	return ans
}

func findRightGreaterIndex(A []int) []int {
	var rHead *Node
	rsIndex := make([]int, len(A))
	for j := len(A) - 1; j >= 0; j-- {
		if rHead == nil {
			rHead = insert(rHead, j)
			rsIndex[j] = -1
		} else if A[rHead.val] <= A[j] {
			for rHead != nil {
				if A[rHead.val] > A[j] {
					rsIndex[j] = rHead.val
					break
				}
				rHead = rHead.next
			}
			if rHead == nil {
				rsIndex[j] = -1
			}
			rHead = insert(rHead, j)
		} else if A[rHead.val] > A[j] {
			rsIndex[j] = rHead.val
			rHead = insert(rHead, j)
		}
	}
	return rsIndex
}

func insert(head *Node, val int) *Node {
	node := newNode(val)
	node.next = head
	head = node
	return head
}
