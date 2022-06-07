package Assignment

import "fmt"

type Node struct {
	value int
	next  *Node
}

var root *Node

func length_ll() int {
	cnt := 0
	temp := root

	for temp != nil {
		temp = temp.next
		cnt++
	}
	return cnt
}

func Insert_node(position, value int) {

	newNode := &Node{
		value: value,
	}

	if position == 1 {
		newNode.next = root
		root = newNode
	} else if position == length_ll()+1 {
		temp := root
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	} else {
		temp := root

		cnt := 1
		for temp.next != nil {
			cnt++
			if cnt == position {
				newNode.next = temp.next
				temp.next = newNode
				return
			}
			temp = temp.next
		}
		if cnt == position {
			newNode.next = temp.next
			temp.next = newNode
			return
		}
	}

}

func Delete_node(position int) {

	if position > length_ll()+1 {
		return
	}

	if root == nil {
		return
	}

	cnt := 1
	temp := root

	if cnt == position {
		root = root.next
		return
	}

	for temp.next != nil {
		prev := temp
		temp = temp.next
		cnt++
		if cnt == position {
			prev.next = temp.next
			return
		}
	}
}

func Print_ll() {
	temp := root
	for temp.next != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Print(temp.value)
}
