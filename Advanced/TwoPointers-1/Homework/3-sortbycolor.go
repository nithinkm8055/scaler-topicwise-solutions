package Homework

func sortColors(A []int) []int {
	c0 := 0
	c1 := 0
	c2 := 0

	for i := range A {
		if A[i] == 0 {
			c0++
		} else if A[i] == 1 {
			c1++
		} else {
			c2++
		}
	}
	c1 += c0
	c2 += c1
	counter := 0
	for i := range A {
		if counter < c0 {
			A[i] = 0
			counter++
		} else if counter >= c0 && counter < c1 {
			A[i] = 1
			counter++
		} else {
			A[i] = 2
		}
	}

	return A
}
