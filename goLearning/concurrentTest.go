package goLearning

import "fmt"

func ConcurrentTest() { // 大写字母开头则为公共方法，可通过包访问
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[len(s)/2:], c) // 后三个的和
	// time.Sleep(100*time.Millisecond)
	go sum(s[:len(s)/2], c) // 前三个的和

	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}
