package main

import "fmt"

func changeTest() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count) // 将整形转化为浮点型
	fmt.Printf("mean 的值为: %f\n", mean)
}

func mapTest() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	// 删除 map 中元素
	delete(countryCapitalMap, "France")
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
func sliceTest() {
	// Go 语言切片是对数组的抽象
	// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
	f := make([]int, 10)      // len = 10; 切片长度为10
	dp := make([]int, 10, 20) // len = 10, cap = 20; 切片长度为10，切片最长可以达到20
	printSlice("f", f)
	printSlice("dp", dp)

	var numbers []int
	printSlice("numbers", numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice("numbers", numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice("numbers", numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice("numbers", numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice("numbers1", numbers1)
}

func printSlice(name string, x []int) {
	fmt.Printf("name=%s len=%d cap=%d slice=%v\n", name, len(x), cap(x), x)
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func structTest() {
	book := Books{"boob", "Bob", "boom", 10}
	ptr := &book

	var nptr *Books
	nptr = &book

	printBooks(book)
	printBooksPtr(ptr)
	printBooksPtr(nptr)
}

func printBooks(books Books) {
	fmt.Printf("book's title is %s", books.title)
	fmt.Printf("book's auther is %s", books.author)
	fmt.Printf("book's subject is %s", books.subject)
	fmt.Printf("book's book_id is %d", books.book_id)
}

func printBooksPtr(books *Books) {
	fmt.Printf("book's title is %s\n", books.title)
	fmt.Printf("book's auther is %s\n", books.author)
	fmt.Printf("book's subject is %s\n", books.subject)
	fmt.Printf("book's book_id is %d\n", books.book_id)
}

func pointerTest() {
	var a int = 10
	fmt.Println("变量的地址: %x", &a) // 以十六进制打印变量的地址

	var ip *int /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	// 指针变量的存储地址
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	// 使用指针访问值
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 空指针
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
}

func arrayTest(arr []int) {
	size := len(arr) // 获取数组大小
	fmt.Println(size)

	var i, j, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}

// 变量传递
func variableTest(t1 int, t2 string) {
	t1 = 9
	t2 = "zxc"
}

// 循环
func loopTest() {
	n := 10
	for i := 0; i < n; i++ {
	}

	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// fo each 循环 —— 使用 range
	strings := []string{"google", "runoob"}
	for i, s := range strings { // 打印 下标 和 字符串值
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5} // 初始化大小为 6 的 int[] 数组，未赋值元素的初始值为0
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

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
