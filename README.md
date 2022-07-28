# golang_merge_intervals

Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return *an array of the non-overlapping intervals that cover all the intervals in the input*.

## Examples

**Example 1:**

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

```

**Example 2:**

```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

```

**Constraints:**

- `1 <= intervals.length <= 104`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 104`

## 解析

給定一個 2D矩陣 intervals，

其中每個 intervals[i] = [$start_i, end_i]$ 代表一個區間

當兩個區間有重疊時， 可以合併位一個範圍較大的區間

要求寫一個演算法來把 intervals 內中所有可以合併的區間都做合併

假設 intervals[i] = [$start_i, end_i$]

        intervals[j] = [$start_j, end_j$]

if $start_i$ ≤ $start_j$ 代表 overlapStar

則如果兩個 區間重疊 代表 $end_i$ ≥ $start_j$

if $end_i$ ≤ $end_j$ 代表 overlapEnd = $end_j$

否則 overlapEnd = $end_i$

透過上述特性

初始化 result = []

只要先把 intervals 根據 start 來做 sort

就可以逐步以每個點 作為 overlapStart 逐步找到其 overlapEnd 

當發現遇到 overlapEnd < intervals[pos][0]

代表上一個 overlap interval 已結束 所以把 [overlapStart, overlapEnd] 加入 result

並且更新 overlapStart = intervals[pos][0], overlapEnd = intervals[pos][1]

當最後走完所有 intervals 時

需要把最後的 [overlapStart, overlapEnd] 加入 result

因為最後一個 overlapEnd 沒有 interval 可以比較

result 即為所求

![](https://i.imgur.com/imNiwpa.png)

時間複雜度是 O(n) 因為需要 loop 整個 intervals

空間複雜度是 O(n) 因為需要儲存回傳值

## 程式碼
```go
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
```
## 困難點

1. 要想出先透過 sort 方式以 start 作為 sort 基準

## Solve Point

- [x]  先對 intervals 以 start 做 sort
- [x]  初始化 overlapStart = intervals[0][0] , overlapEnd = intervals[0][1], result = []
- [x]  從 pos = 1 開始 遍歷 interval
- [x]  if overlapEnd ≥ intervals[pos][0] && overlapEnd ≤ intervals[pos][1]  更新 overlapEnd = intervals[pos][1]
- [x]  if overlapEnd < intervals[pos][0] , 把 [overlapStart, overlapEnd] 加入 result ，更新 overlapEnd = intervals[pos][1] , overlapStart = intervals[pos][0]
- [x]  當走完全部 把 [overlapStart, overlapEnd] 加入 result
- [x]  回傳 result