package Homework

//The monetary system in DarkLand is really simple and systematic. The locals-only use coins. The coins come in different values. The values used are:
//
//1, 5, 25, 125, 625, 3125, 15625, ...
//Formally, for each K >= 0 there are coins worth 5K.
//
//Given an integer A denoting the cost of an item, find and return the smallest number of coins necessary to pay exactly the cost of the item (assuming you have a sufficient supply of coins of each of the types you will need).
//
//
//
//Problem Constraints
//1 <= A <= 2×10^9

func anotherCoinProblem(A int) int {

	arr := make([]int, 0)
	i := 0
	for {
		arr = append(arr, i)
		if pow(5, i) > A {
			break
		}
		i++
	}

	rem := A
	counter := 0

	for {

		rem -= binSearch(rem, arr)
		counter++
		if rem == 0 {
			break
		}

	}

	return counter

}

func binSearch(val int, arr []int) int {
	left := 0
	right := len(arr) - 1

	ans := -1

	for left <= right {
		mid := (left + right) / 2

		if val == pow(5, arr[mid]) {
			ans = pow(5, arr[mid])
			return ans
		} else if val > pow(5, arr[mid]) {
			ans = pow(5, arr[mid])
			left = mid + 1
		} else {

			right = mid - 1
		}

	}

	return ans
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}

	return a * pow(a, b-1)
}
