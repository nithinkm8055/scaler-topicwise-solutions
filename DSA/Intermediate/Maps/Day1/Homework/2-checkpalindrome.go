package Homework

func CheckPalindrome(A string) int {

	countMap := make(map[byte]int)

	for i := range A {
		if _, ok := countMap[A[i]]; ok {
			countMap[A[i]] += 1
		} else {
			countMap[A[i]] = 1
		}
	}

	count := 0
	for _, v := range countMap {
		if v%2 != 0 {
			count++
		}
		if count > 1 {
			return 0
		}
	}

	return 1
}
