package Assignment

func minWindow(A string, B string) string {

	// A bigger string

	if len(A) < len(B) {
		return ""
	}

	l := len(B)
	r := len(A)

	reqFreq := make(map[byte]int)
	for i := range B {
		reqFreq[B[i]]++
	}

	ansStr := ""

	for l <= r {
		mid := (l + r) / 2
		ok, str := checkWindow(A, B, mid, reqFreq)
		if ok {
			r = mid - 1
			ansStr = str
		} else {
			l = mid + 1
		}

	}

	return ansStr

}

func checkWindow(bigger, smaller string, windowSize int, reqFreq map[byte]int) (bool, string) {

	currFreq := make(map[byte]int)

	for j := 0; j < windowSize; j++ {
		currFreq[bigger[j]]++

	}

	if compareFreq(reqFreq, currFreq) {
		return true, bigger[:windowSize]
	}

	for j := 1; j <= len(bigger)-windowSize; j++ {

		currFreq[bigger[windowSize+j-1]]++

		currFreq[bigger[j-1]]--
		if currFreq[bigger[j-1]] == 0 {
			delete(currFreq, bigger[j-1])
		}

		if compareFreq(reqFreq, currFreq) {
			return true, bigger[j : windowSize+j]
		}

	}

	return false, ""
}

func compareFreq(a, b map[byte]int) bool {

	for k, _ := range a {
		if b[k] < a[k] {
			return false
		}
	}

	return true
}
