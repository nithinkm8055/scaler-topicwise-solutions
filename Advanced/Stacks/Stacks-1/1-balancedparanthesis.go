package Stacks_1

type Node struct {
	value byte
	next  *Node
}

func newNode(val byte) *Node {
	newNode := new(Node)
	newNode.value = val
	newNode.next = nil
	return newNode
}

func solve(A string) int {

	closures := []byte(A)

	var head *Node

	pairMap := make(map[byte]byte)
	pairMap['{'] = '}'
	pairMap['['] = ']'
	pairMap['('] = ')'

	for i := range closures {

		if closures[i] == '{' || closures[i] == '[' || closures[i] == '(' {
			node := newNode(pairMap[closures[i]])
			node.next = head
			head = node
		} else if head == nil && (closures[i] == ')' || closures[i] == ']' || closures[i] == '}') {
			return 0
		} else if (closures[i] == ')' && head.value != ')') || (closures[i] == '}' && head.value != '}') || (closures[i] == ']' && head.value != ']') {
			return 0
		} else {
			head = head.next
		}
	}

	if head == nil {
		return 1
	}

	return 0

}
