package Assignment

import "strconv"

func evalRPN(A []string) int {
	var head *Node

	for i := range A {

		if (A[i] != "*") && (A[i] != "-") && (A[i] != "+") && (A[i] != "/") {
			node := newNode(A[i])
			node.next = head
			head = node
		} else {

			op := A[i]
			b := head.value
			head = head.next
			a := head.value
			head = head.next

			temp := operation(a.(string), b.(string), op)
			result := strconv.Itoa(temp)
			node := newNode(result)
			node.next = head
			head = node

		}
	}

	result, _ := strconv.Atoi((head.value).(string))
	return result
}

func operation(a, b, op string) int {

	num1, _ := strconv.Atoi(a)
	num2, _ := strconv.Atoi(b)

	switch op {
	case "*":
		return num1 * num2
	case "-":
		return num1 - num2
	case "/":
		return num1 / num2
	case "+":
		return num1 + num2
	}

	return 0
}
