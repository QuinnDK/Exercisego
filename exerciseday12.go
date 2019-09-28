/*
go语言中range关键字用于for循环中迭代数组，切片，通道，集合的元素

在数组和切片中返回索引及索引对应的值
*/

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { //range求切片的和与数组一致；对于不需要的值可以用_代替
		sum += num
	}
	fmt.Printf("%d\n", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range可以用在map的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana\n"}
	for k, v := range kvs {
		fmt.Printf("%s->%s", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
