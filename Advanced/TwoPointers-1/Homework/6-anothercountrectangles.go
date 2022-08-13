package Homework

func AnotherCountRectangles(A []int, B int) int {

	i := 0
	j := len(A) - 1
	count := 0

	for i < j {

		if A[i]*A[j] < B {
			for k := j; k > i; k-- {
				count += 2
			}
			i++

		} else if A[i]*A[j] >= B {
			j--
		}
	}

	for i := range A {
		if A[i]*A[i] >= B {
			break
		}
		count++
	}

	return count % (1e9 + 7)
}
