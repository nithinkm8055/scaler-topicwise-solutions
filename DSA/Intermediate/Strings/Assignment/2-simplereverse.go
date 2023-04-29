package Assignment

func SimpleReverse(A string) string {

	temp := ""
	for i := len(A) - 1; i >= 0; i-- {
		temp = temp + string(A[i])
	}

	return temp

}
