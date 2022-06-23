package Homework

func findSubSeq(A string, B string) string {

	k := 0
	for i := range A {
		flag := false
		for j := k; j < len(B); j++ {

			if A[i] == B[j] {
				flag = true
				if i == len(A)-1 {
					return "YES"
				}
				k = j + 1
				break
			}

		}
		if !flag {
			return "NO"
		}

	}

	return "NO"
}
