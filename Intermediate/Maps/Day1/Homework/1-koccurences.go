package Homework

func GetSum(A int, B int, C []int) int {

	countMap := make(map[int]int)

	for i := range C {
		if _, ok := countMap[C[i]]; ok {
			countMap[C[i]] = countMap[C[i]] + 1
		} else {
			countMap[C[i]] = 1
		}
	}

	sum := 0
	flag := false
	for k, v := range countMap {
		if v == B {
			sum += k
			flag = true
		}
	}

	if !flag {
		return -1
	}

	return sum

}
