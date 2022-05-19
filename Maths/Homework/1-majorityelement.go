package Homework

func MajorityElement(A []int) int {

	me := A[0]
	count := 1

	for i := 1; i < len(A); i++ {
		if me == A[i] {
			count++
		} else if count == 0 {
			me = A[i]
			count++
		} else {
			count--
		}
	}
	return me
}
