package goLearning

import "fmt"

// go 没有内置 stack 和 queue，需要自己模拟
func stack() {

	// 用切片模拟堆栈
	var st []string

	// push
	st = append(st, "a")
	st = append(st, "b")
	st = append(st, "c")

	// pop
	x := st[len(st)-1]  // 取倒数第一个元素
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
