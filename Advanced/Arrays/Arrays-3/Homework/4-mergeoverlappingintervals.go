package Homework

import "sort"

type Interval struct {
	start, end int
}

func MergeOverlappingIntervals(A []Interval) []Interval {

	sort.SliceStable(A, func(i, j int) bool {
		return A[i].start < A[j].start
	})

	newIntervals := make([]Interval, 0)

	newInterval := A[0]
	flag := false

	for i := 1; i < len(A); i++ {

		if newInterval.start > A[i].end {
			newIntervals = append(newIntervals, A[i])
			flag = true
		}

		if newInterval.end < A[i].start {
			newIntervals = append(newIntervals, newInterval)
			newInterval.start = A[i].start
			newInterval.end = A[i].end
			// flag = true
		}

		c1 := newInterval.start > A[i].start && newInterval.end < A[i].end
		c2 := newInterval.start < A[i].start && newInterval.end > A[i].end
		c3 := newInterval.start > A[i].start && newInterval.start < A[i].end
		c4 := newInterval.end > A[i].start && newInterval.end < A[i].end

		if c1 || c2 || c3 || c4 {
			newInterval.start = min(A[i].start, newInterval.start)
			newInterval.end = max(A[i].end, newInterval.end)
		}
	}

	if !flag {
		newIntervals = append(newIntervals, newInterval)
	}

	return newIntervals
}

// 89 47 76 51 99 28 78 54 94 12 72 31 72 12 55 24 40 59 79 41 100 46 99 5 27 13 23 9 69 39 75 51 53 81 98 14 14 27 89 73 78 28 35 19 30 39 87 60 94 71 90 9 44 56 79 58 70 25 76 18 46 14 96 43 95 70 77 13 59 52 91 47 56 63 67 28 39 51 92 30 66 4 4 29 92 58 90 6 20 31 93 52 75 41 41 64 89 64 66 24 90 25 46 39 49 15 99 50 99 9 34 58 96 85 86 13 68 45 57 55 56 60 74 89 98 23 79 16 59 56 57 54 85 16 29 72 86 10 45 6 25 19 55 21 21 17 83 49 86 67 84 8 48 63 85 5 31 43 48 57 62 22 68 48 92 36 77 27 63 39 83 38 54 31 69 36 65 52 68
