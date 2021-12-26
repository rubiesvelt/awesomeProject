package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 8, 6, 5, 7}
	sort.Ints(arr) // 数组从小到大排序

	sli := make([]int, len(arr)) // 默认创出来都是 0

	sli = append(sli, 9, 9, 9)
	var kkqq []int
	kkqq = sli

	sort.Ints(kkqq) // 结构体从小到大排序
	fmt.Println("ok")

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

/**
5934. 找到和最大的长度为 K 的子序列

输入：nums = [2,1,3,3], k = 2
输出：[3,3]

输入：nums = [-1,-2,3,4], k = 3
输出：[-1,3,4]
*/
func maxSubsequence(nums []int, k int) []int {
	id := make([]int, len(nums)) // 初始化有大小的 slice；var id []int 初始化空slice
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(a, b int) bool {
		return nums[id[a]] > nums[id[b]]
	})
	sort.Ints(id[:k]) // 对 []id 的 [0, k) 进行排序(从小到大)
	ans := make([]int, k)
	for i, j := range id[:k] {
		ans[i] = nums[j]
	}
	return ans
}
