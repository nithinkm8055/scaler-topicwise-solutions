package Assignment

func TaskScheduling(A []int, B []int) int {

	var head *Node
	var tail *Node

	head = newNode(A[0])
	tail = newNode(A[1])

	head.next = tail

	for i := 2; i < len(A); i++ {
		node := newNode(A[i])
		tail.next = node
		tail = node
	}
	counter := 0
	for i := range B {
		for head != nil {

			if head.val != B[i] {
				temp := head
				head = head.next
				temp.next = nil

				tail.next = temp
				tail = tail.next
				counter++
			} else {
				head = head.next
				counter++
				break
			}

		}
	}

	return counter
}
