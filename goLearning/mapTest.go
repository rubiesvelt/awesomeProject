package goLearning

import "fmt"

func mapTest() {
	var mp map[string]string /*创建集合 */
	mp = make(map[string]string)

	mp["France"] = "巴黎"
	mp["Italy"] = "罗马"
	mp["Japan"] = "东京"
	mp["India"] = "新德里"

	for country := range mp {
		fmt.Println(country, "首都是", mp[country])
	}

	// contains key
	capital, ok := mp["American"]

	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	// 删除 map 中元素
	delete(mp, "France")

	// 遍历 map 中元素（随机遍历）
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	for k, v := range hash { // 打印的顺序随机
		println(k, v)
	}
}
