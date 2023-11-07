package Assignment

//There is a row of seats represented by string A. Assume that it contains N seats adjacent to each other.
//There is a group of people who are already seated in that row randomly. i.e. some are sitting together & some are scattered.
//
//An occupied seat is marked with a character 'x' and an unoccupied seat is marked with a dot ('.')
//
//Now your target is to make the whole group sit together i.e. next to each other, without having any vacant seat between them in such a way that the total number of hops or jumps to move them should be minimum.
//
//In one jump a person can move to the adjacent seat (if available).

func seats(A string) int {

	seats := make(map[int]int) // stores the number of occupied seat + index
	allocator := make([]int, len(A))

	count := 0
	for i := range A {
		if A[i] == 'x' {
			allocator[i] = 1
			seats[count] = i
			count++
		}
	}

	midIndex := seats[count/2]
	hopCounter := 0
	leftClosest := -1
	rightClosest := -1

	for i := range A {

		if allocator[i] == 1 && i != midIndex {
			if midIndex-i > 1 {
				allocator[i] = 0
				j := midIndex - 1

				if leftClosest != -1 {
					j = leftClosest
				}

				for ; j >= i; j-- {
					if allocator[j] == 0 {
						allocator[j] = 1
						leftClosest = j - 1
						hopCounter += (j - i)
						break
					}
				}
			} else if i-midIndex > 1 {
				allocator[i] = 0

				j := midIndex + 1

				if rightClosest != -1 {
					j = rightClosest
				}

				for ; j <= i; j++ {
					if allocator[j] == 0 {
						allocator[j] = 1
						rightClosest = j + 1
						hopCounter += (i - j)
						break
					}
				}
			}
		}
	}
	return hopCounter % 10000003
}
