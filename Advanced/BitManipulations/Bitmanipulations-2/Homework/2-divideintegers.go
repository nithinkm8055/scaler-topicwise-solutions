package Homework

func divide(A int, B int) int {

	cnt := 0
	for i := 31; i >= 0; i-- {

		if A >= (B << i) {
			A = A - (B << i)
			cnt += (1 << i)
		}

	}
	return cnt
}
