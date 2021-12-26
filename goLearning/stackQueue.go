package goLearning

import "fmt"

// go 没有内置 stack 和 queue，需要自己模拟
func stack() {
	//堆栈
	//先进后出
	var st []string
	//push
	//append
	st = append(st, "a")
	st = append(st, "b")
	st = append(st, "c")
	//pop
	//后面移除
	x := st[len(st)-1]
	st = st[:len(st)-1] // 切片，下标[0, len(st)-1)
	fmt.Println("1: ", x)

	x = st[len(st)-1]
	st = st[:len(st)-1]
	fmt.Println("2: ", x)

	x = st[len(st)-1]
	st = st[:len(st)-1]
	fmt.Println("3: ", x)
}

func queue() {
	//队列
	//先进先出
	var q []string
	//push
	//append
	q = append(q, "a", "b")
	q = append(q, "c")
	//pop
	x := q[0]
	q = q[1:]
	fmt.Println("1: ", x)

	x = q[0]
	q = q[1:]
	fmt.Println("2: ", x)

	x = q[0]
	q = q[1:]
	fmt.Println("3: ", x)
}
