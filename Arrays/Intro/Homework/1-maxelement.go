package Homework

func MaxElement(A []int, B int) int {
	count := 0
	flag := false
	for i := range A {
		if A[i] > B {
			count++
		}
		if A[i] == B {
			flag = true
		}
	}
	if flag {
		return count
	}
	return -1
}
