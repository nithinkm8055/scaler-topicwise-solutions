package Homework

func CheckPalindrome(A string) int {

	countMap := make(map[byte]int)

	for i := range A {
		countMap[A[i]] = countMap[A[i]] + 1
	}

	count := 0
	for _, v := range countMap {
		if v%2 != 0 {
			count++
		}
	}

	if count <= 1 {
		return 1
	}

	return 0
}
