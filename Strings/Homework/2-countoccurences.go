package Homework

func CountOccurences(A string) int {

	count := 0
	for i := 0; i < len(A)-2; i++ {
		if A[i:i+3] == "bob" {
			count++
			i += 1
		}
	}
	return count

}
