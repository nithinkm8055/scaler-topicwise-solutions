package Assignment

func Bulbs(A []int) int {

	flag := false
	count := 0
	for i := range A {
		if A[i] == 0 && !flag {
			flag = true
			count++
		} else if A[i] == 1 && flag {
			flag = false
			count++
		}
	}

	return count
}
