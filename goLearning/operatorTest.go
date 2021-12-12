package goLearning

import "fmt"

// 运算符
func operatorTest() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int // ptr 是一个指针变量

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("ptr 为 %d\n", ptr)   // ptr 为 1374390755336
	fmt.Printf("*ptr 为 %d\n", *ptr) // *ptr 为 4
}

func learn() {
	var a, b int
	var c string
	a, b, c = 5, 7, "abc" // 变量 a，b 和 c 都已经被声明，否则的话应该这样使用：a, b, c := 5, 7, "abc"
	fmt.Println(a, b, c)
}
