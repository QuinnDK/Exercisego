/*
go语言集合Map
无需键值对的集合，通过Key快速查找对应的值，由于map是集合所以可以用迭代的放法去遍历
但由于其是无序的，所以无法决定他的返回顺序（map是用hash实现的）
*/
package main

import "fmt"

func main() {
	var countrycaptialmap map[string]string
	countrycaptialmap = make(map[string]string)

	//插入键值对
	countrycaptialmap["France"] = "巴黎"
	countrycaptialmap["Italy"] = "罗马"
	countrycaptialmap["Japan"] = "东京"
	countrycaptialmap["India"] = "新德里"

	for country := range countrycaptialmap {
		fmt.Println(country, "首都是", countrycaptialmap[country])
	}
	fmt.Println("-----------------------------")
	delete(countrycaptialmap, "France")

	for country := range countrycaptialmap {
		fmt.Println(country, "首都是", countrycaptialmap[country])
	}
	fmt.Println()
	//查看元素是否在集合中
	captial, ok := countrycaptialmap["American"]
	if ok {
		fmt.Println("American的首都是", captial)
	} else {
		fmt.Println("American的首都不存在")
	}
}
