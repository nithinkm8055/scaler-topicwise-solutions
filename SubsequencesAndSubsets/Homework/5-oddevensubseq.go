package Homework

func oddEvenSubSeq(A []int) int {

	evenFlag := false
	oddFlag := false

	count := 0
	for i := range A {

		if A[i]%2 == 0 && !evenFlag {
			evenFlag = true
			oddFlag = false
			count++
		}

		if A[i]%2 != 0 && !oddFlag {
			oddFlag = true
			evenFlag = false
			count++
		}

	}

	return count

}
