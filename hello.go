package main

import (
	"fmt"
)

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[len(s)/2:], c) // 后三个的和
	// time.Sleep(100*time.Millisecond)
	go sum(s[:len(s)/2], c) // 前三个的和

	x, y := <-c, <-c // 从通道 c 中接收

	e := RPCError{}
	e.Error()

	fmt.Println(x, y, x+y)

	type Test struct{}
	v := Test{}
	Print(v)
}

func Print(v interface{}) {
	println(v)
}

type throwShit interface {
	throw() string
}

type RPCError struct { // RPCError实现了error接口
	Code    int64
	Message string
}

func (e *RPCError) Error() string { // 意思是 此方法属于 RPCError 结构体
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func (e *RPCError) throw() string {
	return "shit was threw"
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

/*
 * 343. 整数拆分
 */
func integerBreak(n int) int {
	f := make([]int, n+1) // 创建切片
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			f[i] = max(f[i], max((i-j)*f[j], (i-j)*j))
		}
	}
	return f[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func learn() {
	var a, b int
	var c string
	a, b, c = 5, 7, "abc" // 变量 a，b 和 c 都已经被声明，否则的话应该这样使用：a, b, c := 5, 7, "abc"
	fmt.Println(a, b, c)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
