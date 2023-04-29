package Homework

func EvenSubArrays(A []int) string {

	if len(A)%2 == 0 && A[0]%2 == 0 && A[len(A)-1]%2 == 0 {
		return "YES"
	}
	return "NO"
}
