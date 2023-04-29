package Homework

func specialsubseqAG(A string) int {

	gCount := 0
	ans := 0
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 'G' {
			gCount++
		} else if A[i] == 'A' {
			ans += gCount
		}
	}
	return ans

}
