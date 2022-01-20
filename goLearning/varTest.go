package goLearning

import (
	"fmt"
	"strings"
)

func stringTest() {
	s := "abc,def,gh"
	var strs []string
	// 以逗号分隔成数组
	strs = strings.Split(s, ",")

	// 判断是否包含子数组
	b := strings.Contains(s, "abc")

	// starts with
	b = strings.HasPrefix("abcaaa", "abc")

	// 拼接
	s = "aa" + "bb" + "cc"

	// utf-8 遍历
	for i := 0; i < len(s); i++ {
		ch := s[i]
		fmt.Println(ch)
	}

	// unicode 遍历
	for _, ch1 := range s {
		fmt.Println(ch1)
	}
	fmt.Println(strs, b)
}

// 类型转换 test
func changeTest() {
	var sum = 17 // 与 sum := 17 相同
	var count = 5
	var mean float32

	mean = float32(sum) / float32(count) // 将整形转化为浮点型
	fmt.Printf("mean 的值为: %f\n", mean)
}

// 变量
func varTest() {
	var age int
	age = age + 1 // 变量定义了就必须使用

	// 以下三种定义变量的方法
	f1 := "Runoob"           // 相当于定义变量再赋值
	var f2 = "Runoob"        // 可以自动推断变量类型
	var f3 string = "Runoob" // 可以不写类型
	fmt.Println("hello world" + f1 + f2 + f3)

	// 交换 f2, f3 的值
	f2, f3 = f3, f2

	// 同时定义两个变量
	var u, c = 1, 2
	fmt.Println(u, c)

	g, h := 123, "hello"
	fmt.Println(g, h)

	// 变量的初始值
	var i int     // 0
	var f float64 // 0
	var b bool    // false
	var s string  // ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// 也可以这样定义方法
	// 只获取函数返回值的后两个
	_, numb, strs := numbers()
	fmt.Println(numb, strs)

	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}

// 常量
func constTest() {
	const (
		// iota 是一个自增的数
		a = iota // 0
		b        // 1
		c        // 2
		d = "ha" // "ha" 独立值，iota += 1
		e        // "ha"   iota += 1
		f = 100  // 100    iota +=1
		g        // 100    iota +=1
		h = iota // 7,恢复计数
		i        // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i) // 0 1 2 ha ha 100 100 7 8
}
