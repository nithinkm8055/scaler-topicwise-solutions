package Assignment

func SubArraywithGivenSum(A []int, B int) []int {

	left := 0
	right := 0

	sum := A[left]
	for right < len(A) && left <= right {
		if sum == B {
			return A[left : right+1]
		}

		if sum < B {
			right++
			if right < len(A) {
				sum += A[right]
			}
		} else if sum > B {
			sum = sum - A[left]
			left++
			if left > right && left < len(A)-1 {
				right++
				sum += A[right]
			}
		}
	}

	return []int{-1}

}

// Brute force
func SubArraywithGivenSum2(A []int, B int) []int {

	for i := range A {
		sum := 0
		for j := i; j < len(A); j++ {
			sum += A[j]
			if sum == B {
				result := make([]int, 0)
				for k := i; k <= j; k++ {
					result = append(result, A[k])
				}
				return result
			} else if sum > B {
				break
			}
		}
	}

	return []int{-1}

}
