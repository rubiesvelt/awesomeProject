package goLearning

import "fmt"

func SetTest() {
	set := make(map[interface{}]bool) // go 语言没有 set，我们使用 map 做 set

	set[11] = true

	if set[12] {
		fmt.Println("set contains 12")
	}

	set[10] = false
	b := set[100]

	fmt.Println(b)
}
