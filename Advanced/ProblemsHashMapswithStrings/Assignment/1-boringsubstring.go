package Assignment

// TODO: resolve TLE
func BoringSubstring(A string) int {

	count := 0
	for i := range A {
		for j := i + 1; j < len(A); j++ {
			d := int(A[j]) - int(A[i])
			if d < 0 {
				d = d * -1
			}

			if d >= 2 {
				count++
			} else {
				count--
			}
		}
	}

	if count >= 0 {
		return 1
	}
	return 0

}
