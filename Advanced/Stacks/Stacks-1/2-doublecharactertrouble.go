package Stacks_1

func DoubleCharacterTrouble(A string) string {

	chars := []byte(A)
	var head *Node

	for j := len(A) - 1; j >= 0; j-- {

		node := newNode(chars[j])
		if head == nil {
			node.next = head
			head = node
		} else {
			if head.value == chars[j] {
				head = head.next
			} else {
				node.next = head
				head = node
			}
		}
	}

	result := make([]byte, 0)
	temp := head

	for temp != nil {
		result = append(result, (temp.value).(byte))
		temp = temp.next
	}

	return string(result)

}
