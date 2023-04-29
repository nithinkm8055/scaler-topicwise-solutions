package Homework

func plusOne(A []int) []int {

	zeroIndex := 0
	for i := range A {
		if A[i] != 0 {
			zeroIndex = i
			break
		}
	}

	A = A[zeroIndex:]
	carry := 1
	for j := len(A) - 1; j >= 0; j-- {
		temp := A[j] + carry
		A[j] = temp % 10

		carry = temp / 10

	}

	if carry > 0 {
		A = append([]int{1}, A...)
	}

	return A
}
