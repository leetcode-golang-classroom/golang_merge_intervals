package sol

import "sort"

type ByStart [][]int

func (a ByStart) Len() int {
	return len(a)
}
func (a ByStart) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func (a ByStart) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func merge(intervals [][]int) [][]int {
	result := [][]int{}
	sort.Sort(ByStart(intervals))
	overlapStart, overlapEnd := intervals[0][0], intervals[0][1]
	nIntervals := len(intervals)
	for pos := 1; pos < nIntervals; pos++ {
		if overlapEnd <= intervals[pos][1] && overlapEnd >= intervals[pos][0] {
			overlapEnd = intervals[pos][1]
		}
		if overlapEnd < intervals[pos][0] {
			result = append(result, []int{overlapStart, overlapEnd})
			overlapStart = intervals[pos][0]
			overlapEnd = intervals[pos][1]
		}
	}
	result = append(result, []int{overlapStart, overlapEnd})
	return result
}
