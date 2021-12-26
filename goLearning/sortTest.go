package goLearning

import (
	"fmt"
	"sort"
)

type Dummy struct {
	a int
	b int
}

func SortTest() {
	arr := []int{4, 3, 1, 5, 2}
	sli := make([]int, 5)

	copy(sli, arr)
	// 升序排序
	// sort.Ints(arr)

	// 降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(sli)))

	// 降序排序
	sort.Slice(arr, func(i, j int) bool { // i, j 是下标；从大到小排序
		return arr[i] > arr[j]
	})

	dummySli := make([]Dummy, 0)
	dummySli = append(dummySli, Dummy{1, 2}, Dummy{a: 3, b: 6})
	sort.Slice(dummySli, func(i int, j int) bool {
		return dummySli[i].b-dummySli[i].a > dummySli[j].b-dummySli[j].a
	})

	fmt.Println(dummySli)
}
