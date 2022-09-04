package Homework

type Node struct {
	next  *Node
	value byte
}

func newNode(val byte) *Node {
	node := new(Node)
	node.value = val
	node.next = nil
	return node
}

func braces(A string) int {
	// check if there is an operator between a pair of paranthesis
	// remove a paranthesis as soon you find an operator

	opMap := map[byte]bool{'+': true, '-': true, '*': true, '/': true}

	var head *Node

	for i := range A {

		if A[i] == '(' {
			node := newNode(')')
			node.next = head
			head = node
		} else if _, ok := opMap[A[i]]; ok {
			node := newNode(A[i])
			node.next = head
			head = node
		} else if A[i] == ')' {
			if _, ok := opMap[head.value]; ok {
				head = head.next.next
			} else {
				return 1
			}
		}
	}

	if head != nil && head.value == ')' {
		return 1
	}

	return 0

}
