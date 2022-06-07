package Assignment

import "fmt"

type listNode struct {
	Value int
	Next  *listNode
}

func listNode_new(val int) *listNode {
	var node *listNode = new(listNode)
	node.Value = val
	node.Next = nil
	return node
}

var head *listNode
var size int

func DesignLinkedList(A [][]int) *listNode {

	for i := range A {

		op := A[i][0]
		newNode := listNode_new(A[i][1])

		switch op {
		case 0:
			newNode.Next = head
			head = newNode
			size++
		case 1:
			temp := head
			if head == nil {
				head = newNode
			} else {
				for temp.Next != nil {
					temp = temp.Next
				}
				temp.Next = newNode
			}
			size++
		case 2:
			temp := head
			cnt := 0

			if cnt == A[i][2] {
				newNode.Next = head
				head = newNode
				size++
				continue
			}

			for temp.Next != nil {
				cnt++
				if cnt == (A[i][2]) {
					newNode.Next = temp.Next
					temp.Next = newNode
					size++
					continue
				} else {
					temp = temp.Next
				}
			}
			if cnt == (A[i][2]) {
				newNode.Next = temp.Next
				temp.Next = newNode
			}
			size++
		case 3:
			temp := head
			cnt := 0

			if cnt == A[i][1] && head != nil {
				head = head.Next
				size--
				continue
			}

			if temp != nil {
				for temp.Next != nil {
					cnt++
					prev := temp
					if cnt == A[i][1] {
						temp = temp.Next
						prev.Next = temp.Next
						size--
						continue
					} else {
						temp = temp.Next
					}
				}
			}
		}

	}

	fmt.Println(size)
	return head

}
