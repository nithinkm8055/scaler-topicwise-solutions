package Assignment

func findCommon(A *listNode, B *listNode) int {

	count := 0
	for A != nil && B != nil {
		if A.value == B.value {
			count++
		} else {
			return count
		}
		A = A.next
		B = B.next
	}

	return count

}

func LongestPalindromicList(A *listNode) int {

	curr := A
	ans := 0

	var prev *listNode
	var temp *listNode

	for curr != nil {

		temp = curr.next
		curr.next = prev

		ans = max(ans, 2*findCommon(prev, temp)+1)

		ans = max(ans, 2*findCommon(curr, temp))
		prev = curr
		curr = temp
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
