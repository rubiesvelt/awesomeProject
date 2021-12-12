package goLearning

import "fmt"

// array 和 slice
func arrayTest(arr []int) {
	// 数组
	arr0 := [6]int{1, 2, 3, 5} // 初始化大小为 6 的 int[] 数组，未赋值元素的初始值为 0
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
	f := make([]int, 10)      // len = 10; 切片长度为10
	dp := make([]int, 10, 20) // len = 10, cap = 20; 切片长度为10，切片容量为20（最长可以达到20）
	printSlice("f", f)
	printSlice("dp", dp)

	// 创建一个 slice 叫 newSlice0
	var newSlice0 []int
	printSlice("newSlice0", newSlice0)

	// 向切片添加一个元素
	newSlice0 = append(newSlice0, 1)
	printSlice("newSlice0", newSlice0)

	// 同时添加多个元素
	newSlice0 = append(newSlice0, 2, 3, 4)
	printSlice("newSlice0", newSlice0)

	// 创建切片 newSlice1 是之前切片的两倍容量
	len0 := len(newSlice0)
	cap0 := cap(newSlice0)
	len1 := len0
	cap1 := cap0 << 1
	newSlice1 := make([]int, len1, cap1)

	// 拷贝 newSlice0 的内容到 newSlice1 ；func copy(dst, src []Type) int
	copy(newSlice1, newSlice0)
	printSlice("newSlice1", newSlice1)
}

func printSlice(name string, x []int) {
	fmt.Printf("name=%s len=%d cap=%d slice=%v\n", name, len(x), cap(x), x)
}
