package goLearning

import "fmt"

// 循环
func forLoopTest() {

	// 经典三段式
	n := 10
	for i := 0; i < n; i++ {
	}

	// 三段式 —— 仅有退出条件；类似 while，go 语言没有 while
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 使用 range
	strings := []string{"google", "runoob"}

	for i := range strings { // 打印下标
		fmt.Println(i)
	}

	for i, s := range strings { // 打印 下标 和 字符串值
		fmt.Println(i, s)
	}
}

func rangeTest() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	// 以下两种写法一样
	// 下标遍历写法
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("sum:", sum)

	// for each 写法 (range)
	sum1 := 0
	for index, num := range nums {
		fmt.Println("index:", index)
		sum1 += num
	}
	fmt.Println("sum1:", sum1)

	// 在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
