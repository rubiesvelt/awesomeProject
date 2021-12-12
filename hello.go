package main

import "fmt"
import "awesomeProject/goLearning"

func main() {
	fmt.Println()
	goLearning.ConcurrentTest()
}

// 容器
// Map, Set, Stack, Queue, Deque, List
// Set, SortedSet, TreeSet
// List, list.sort(), Comparator

/**
5935. 适合打劫银行的日子

security = [5,3,3,3,5,6,2], time = 2
-> [2,3]

*/
func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	left := make([]int, n) // 左边有几个大于等于 i 这个且递减的
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if security[i+1] >= security[i] {
			right[i] = right[i+1] + 1
		}
	}
	var ans []int // 定义不限大小的空数组
	for i := 0; i < n; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return ans
}
