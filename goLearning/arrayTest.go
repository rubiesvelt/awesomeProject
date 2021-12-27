package goLearning

import "fmt"

// array 和 slice
func arrayTest(arr []int) {
	// 数组
	// 初始化大小为 6 的 int[] 数组，未赋值元素的初始值为 0
	// [6]代表初始化了一个数组，而不是slide
	arr0 := [6]int{1, 2, 3, 5}
	for i, u := range arr0 {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, u)
	}

	arr1 := []int{1, 2, 3, 4} // 未指定大小则 大小为 4
	for i, u := range arr1 {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, u)
	}

	balance := [5]int{1: 2.0, 3: 7.0} // 将索引为 1 和 3 的元素初始化，初始化后的数组为 {0, 2.0, 0, 3.0, 0}
	for k := 0; k < 5; k++ {
		fmt.Printf("balance[%d] = %f\n", k, balance[k])
	}

	size := len(arr) // 获取数组大小
	fmt.Println(size)
}

func sliceTest() {
	// Go 语言切片是对数组的抽象
	// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
	t := make([]int, 0)       // len = 0; 切片长度为0
	f := make([]int, 10)      // len = 10; 切片长度为10
	dp := make([]int, 10, 20) // len = 10, cap = 20; 切片长度为10，切片容量为20（最长可以达到20）

	// 创建一个空 slice 叫 slice；多种方法效果相同
	// var slice []int
	// slice := []int{}
	// slice := make([]int, 0)
	//
	// 创建有初始值的 slice
	// slice := []int{1,2,3}
	var slice []int

	// 向切片添加一个元素
	slice = append(slice, 1)

	// 同时添加多个元素
	slice = append(slice, 2, 3, 4)

	// 创建切片 newSlice1 是之前切片的两倍容量
	len0 := len(slice)
	cap0 := cap(slice)
	len1 := len0
	cap1 := cap0 << 1
	slice1 := make([]int, len1, cap1)

	// 拷贝 slice 的内容到 newSlice1 ；func copy(dst, src []Type) int
	copy(slice1, slice)

	fmt.Println(t, f, dp)
}
