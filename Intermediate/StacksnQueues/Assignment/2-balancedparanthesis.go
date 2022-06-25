package Assignment

type LLNode struct {
	value byte
	next  *LLNode
}

var head *LLNode

func BalancedParanthesis(A string) int {

	for i := range A {
		if A[i] == '(' {
			push(A[i])
		} else {
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

func push(val byte) {
	newNode := &LLNode{
		value: val,
	}

	newNode.next = head
	head = newNode
}

func pop() {
	head = head.next
}
