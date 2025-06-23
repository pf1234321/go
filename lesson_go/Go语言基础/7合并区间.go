package main

import (
	"fmt"
	"sort"
)

// 以数组 intervals 表示若干个区间的集合，
// 其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，
// 该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
// 然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，
// 如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func main() {
	numbers := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	println(numbers)
	numbers2 := merge(numbers)
	for i, ints := range numbers2 {
		fmt.Printf("%d: [%d, %d]\n", i, ints[0], ints[1])
	}
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按区间的起始点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]} // 初始化合并后的区间数组

	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] { // 重叠则合并
			if interval[1] > last[1] {
				last[1] = interval[1] // 更新结束点
			}
		} else { // 不重叠则添加新区间
			merged = append(merged, interval)
		}
	}

	return merged
}
