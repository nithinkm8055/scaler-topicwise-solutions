package Assignment

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */

type Node struct {
	val  int
	next *Node
}

func newNode(val int) *Node {
	node := new(Node)
	node.val = val
	node.next = nil
	return node
}

func prevSmaller(A []int) []int {

	var head *Node

	result := make([]int, 0)

	for i := range A {

		if head == nil {
			result = append(result, -1)
			head = insert(head, A[i])
		} else if head.val >= A[i] {
			flag := false
			for head != nil {
				if head.val < A[i] {
					result = append(result, head.val)
					flag = true
					break
				}
				head = head.next
			}
			if !flag {
				result = append(result, -1)
			}

			head = insert(head, A[i])
		} else if head.val < A[i] {
			result = append(result, head.val)
			head = insert(head, A[i])
		}
	}

	return result
}

func insert(head *Node, val int) *Node {
	node := newNode(val)
	node.next = head
	head = node
	return head
}
