package main

import "fmt"

type Duck interface {
	Quack()
}

type Cat struct{}

// Quack
// 我们使用指针实现接口 (c *Cat) 时，该方法属于指针 *Cat
// 当我们使用结构体实现接口 (c Cat) 时，该方法属于结构体 Cat
func (c *Cat) Quack() {
	fmt.Println("meow")
}

func QuackTest() {
	/**
	go语言在参数传递时，都是传递值
	作为指针的 &Cat{} 变量能够隐式地获取到指向的结构体

	Quack方法是属于 *Cat{} 指针，不属于 Cat{} 结构体

	使用 &Cat{} 指针初始化变量时，复制出一个 &Cat{} 指针，两个指针都是指向 Cat{} 结构体，作为指针的 &Cat{} 变量能够隐式地获取到指向的结构体
	使用 Cat{} 使用结构体初始化变量时，复制出一个 Cat{} 结构体，编译器不会凭空产生指针
	*/
	var z Duck = &Cat{}
	z.Quack()
}
