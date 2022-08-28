package Homework

func flip(A string) []int {
	count := 0
	for i := range A {
		if A[i] == '1' {
			count++
		}
	}

	if count == len(A) {
		return []int{}
	}

	l := -1
	r := -1
	max := 0
	count = 0
	ansLIndex := -1
	ansRIndex := -1

	reset := false
	for i := range A {
		if A[i] == '1' {
			count--
		} else if A[i] == '0' {
			count++
		}

		if count >= 0 {

			if l == -1 || reset {
				l = i
				r = i
			} else {
				r = i
			}

			reset = false
		} else if count < 0 {
			count = 0
			reset = true
		}

		if count >= max && !reset {
			if count > max {
				ansLIndex = l
				ansRIndex = r
			}

			max = count
		}
	}

	return []int{ansLIndex + 1, ansRIndex + 1}
}
