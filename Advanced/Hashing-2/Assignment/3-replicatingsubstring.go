package Assignment

func ReplicatingSubstring(A int, B string) int {

	countMap := make(map[byte]int)

	for i := range B {
		countMap[B[i]] = countMap[B[i]] + 1
	}

	for _, v := range countMap {
		if v%A != 0 {
			return -1
		}
	}

	return 1
}
