package Assignment

//There is a row of seats represented by string A. Assume that it contains N seats adjacent to each other.
//There is a group of people who are already seated in that row randomly. i.e. some are sitting together & some are scattered.
//
//An occupied seat is marked with a character 'x' and an unoccupied seat is marked with a dot ('.')
//
//Now your target is to make the whole group sit together i.e. next to each other, without having any vacant seat between them in such a way that the total number of hops or jumps to move them should be minimum.
//
//In one jump a person can move to the adjacent seat (if available).

// Solution gets TLE
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

	for i := range A {

		if allocator[i] == 1 && i != midIndex {
			allocator[i] = 0
			if midIndex-i > 1 {
				for j := midIndex - 1; j >= i; j-- {
					if allocator[j] == 0 {
						allocator[j] = 1
						hopCounter += (j - i)
						break
					}
				}
			} else if i-midIndex > 1 {
				for j := midIndex + 1; j <= i; j++ {
					if allocator[j] == 0 {
						allocator[j] = 1
						hopCounter += (i - j)
						break
					}
				}
			}
		}

	}

	return hopCounter
}
