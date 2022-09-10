package Assignment

import "strconv"

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

// TLE
func PerfectNumbers(A int) string {

	counter := 0

	var head *Node
	var tail *Node

	head = newNode(1)
	tail = newNode(2)
	head.next = tail

	for head != nil {

		temp := head.val
		head = head.next // remove one element

		node1 := newNode(temp*10 + 1)
		tail.next = node1
		tail = tail.next

		node2 := newNode(temp*10 + 2)
		tail.next = node2
		tail = tail.next

		if isPalindrome(node1.val) {
			counter++
			if counter == A {
				return strconv.Itoa(node1.val)
			}
		}
		if isPalindrome(node2.val) {
			counter++
			if counter == A {
				return strconv.Itoa(node2.val)
			}
		}
	}

	return ""

}

func isPalindrome(n int) bool {

	revNum := 0
	temp := n
	len := 0
	for temp > 0 {
		r := temp % 10
		revNum = revNum*10 + r
		temp = temp / 10
		len++
	}

	if n == revNum && (len%2) == 0 {
		return true
	}

	return false
}
