package main

import "fmt"

// Go 语言中 函数可以返回新函数
func newFunctionTest() func(n int) int {
	return func(n int) int {
		return n << 1
	}
}

// 定义新函数结构体
type demoFunction func(a int, b int) int

func createDemoFunction(c int) demoFunction {

	index := c << 1
	fmt.Println(index)
	return func(a int, b int) int {
		return a + b + c
	}
}
