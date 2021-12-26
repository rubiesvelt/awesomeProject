package goLearning

import "fmt"

func PointerTest() {
	var a = 10
	fmt.Println("变量的地址: %x", &a) // %x 以十六进制打印变量的地址

	// *int 表示 int 类型的指针；* 表示这是一个指针
	var ptr *int

	// &a 表示取变量 a 的地址；& 表示取地址；&a 是一个 *int 型变量
	ptr = &a

	fmt.Printf("a 变量的地址是: %x\n", &a)       // &a 表示取变量 a 的地址
	fmt.Printf("ptr 变量储存的指针地址: %x\n", ptr) // 打印 指针 的值（地址）
	fmt.Printf("*ptr 变量的值: %d\n", *ptr)    // *ptr 表示 指针所指 的值

	var ptr1 *int
	fmt.Printf("空指针 ptr1 的值为 : %x\n", ptr1)
}
