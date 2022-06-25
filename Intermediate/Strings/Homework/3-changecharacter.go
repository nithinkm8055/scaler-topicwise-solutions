package Homework

import (
	"sort"
)

// find freq of each character in string
// sort it based on freq
// remove elements (thus its frequency) and decrement B
func ChangeCharacter(A string, B int) int {

	countMap := make(map[byte]int)
	for i := range A {
		if _, ok := countMap[A[i]]; ok {
			countMap[A[i]] = countMap[A[i]] + 1
		} else {
			countMap[A[i]] = 1
		}
	}

	type kv struct {
		k byte
		v int
	}

	var ss []kv

	for k, v := range countMap {
		ss = append(ss, kv{k, v})
	}

	sort.SliceStable(ss, func(i, j int) bool {
		return ss[i].v < ss[j].v
	})

	for i := range ss {
		if ss[i].v <= B && B > 0 {
			B = B - ss[i].v
			ss[i].v = 0
		} else if ss[i].v >= B && B > 0 {
			ss[i].v = ss[i].v - B
			B = 0
		}
	}
	count := 0
	for i := range ss {
		if ss[i].v > 0 {
			count++
		}
	}

	return count

}
