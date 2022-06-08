package Homework

type Node struct {
	value byte
	next  *Node
}

var head *Node

func BP(A string) int {

	for i := range A {

		if A[i] == '{' || A[i] == '[' || A[i] == '(' {
			push(A[i])
		} else if A[i] == '}' {
			if head != nil && head.value == '{' {
				pop()
			} else {
				return 0
			}
		} else if A[i] == ']' {
			if head != nil && head.value == '[' {
				pop()
			} else {
				return 0
			}
		} else if A[i] == ')' {
			if head != nil && head.value == '(' {
				pop()
			} else {
				return 0
			}
		}

	}

	if head != nil {
		return 0
	}

	return 1
}

func push(char byte) {
	newNode := &Node{
		value: char,
	}

	newNode.next = head
	head = newNode
}

func pop() {
	if head != nil {
		head = head.next
	}
}
