package Arrays_3

type Interval struct {
	start, end int
}

func insert(intervals []Interval, newInterval Interval) []Interval {

	if newInterval.start > newInterval.end {
		newInterval.start, newInterval.end = newInterval.end, newInterval.start
	}

	flag := false
	resultInterval := make([]Interval, 0)
	for i := range intervals {

		if newInterval.start > intervals[i].end {
			resultInterval = append(resultInterval, intervals[i])
		}

		if intervals[i].start > newInterval.end {

			resultInterval = append(resultInterval, newInterval)
			for i < len(intervals) {
				resultInterval = append(resultInterval, intervals[i])
				i++
			}
			flag = true

			break
		}

		if intervals[i].start < newInterval.end && newInterval.start < intervals[i].end {
			newInterval.start = min(intervals[i].start, newInterval.start)
			newInterval.end = max(intervals[i].end, newInterval.end)
		}
	}

	if !flag {
		resultInterval = append(resultInterval, newInterval)
	}

	return resultInterval

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
