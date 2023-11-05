package Assignment

//N children are standing in a line. Each child is assigned a rating value.
//
//You are giving candies to these children subjected to the following requirements:
//
//Each child must have at least one candy.
//Children with a higher rating get more candies than their neighbors.
//What is the minimum number of candies you must give?

func candy(A []int) int {

	candies := make([]int, len(A))
	totalCandies := len(A)
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for j := len(A) - 2; j >= 0; j-- {
		if j > 0 && A[j] > A[j+1] && A[j] > A[j-1] {
			candies[j] = max(candies[j-1], candies[j+1]) + 1
		} else if A[j] > A[j+1] {
			candies[j] = candies[j+1] + 1
		}
	}

	for i := range candies {
		totalCandies += candies[i]
	}

	return totalCandies
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
