package Assignment

import "sort"

var a []int
var ans [][]int

func subset(temp []int, index int) {
	//insert
	ans = append(ans, temp)
	for i := index; i < len(a); i++ {
		temp = append(temp, a[i])
		subset(temp, i+1)
		var cur []int
		for i := 0; i < len(temp)-1; i++ {
			cur = append(cur, temp[i])
		}
		temp = cur
	}
	return
}
func subsets(A []int) [][]int {
	ans = nil
	sort.Ints(A)
	a = A
	index := 0
	var temp []int
	subset(temp, index)
	return ans
}

//  My solution
//func subsets(A []int )  ([][]int) {
//
//	result := make([][]int, 0)
//	for i := 0 ; i < (1 << len(A)); i++ {
//		test := make([]int, 0)
//		for j := 0 ;j < len(A); j++ {
//
//			if i & (1 << j) != 0 {
//				test = append(test, A[j])
//			}
//
//		}
//		sort.Ints(test)
//		result = append(result, test)
//
//		sort.SliceStable(result, func(i, j int) bool {
//			if len(result[i]) > 0 && len(result[j]) > 0 {
//				if len(result[i]) >= len(result[j]){
//					for k := 0; k < len(result[j]); k++ {
//						if result[i][k] < result[j][k] {
//							return true
//						} else if result[i][k] > result[j][k] {
//							return false
//						}
//
//					}
//
//					return false
//				} else if  len(result[j]) > len(result[i]){
//					for k := 0; k < len(result[i]); k++ {
//						if result[i][k] < result[j][k] {
//							return true
//						} else if result[i][k] > result[j][k] {
//							return false
//						}
//
//					}
//					return false
//				}
//
//			}
//
//			return false
//
//		})
//
//	}
//
//	return result
//
//}
