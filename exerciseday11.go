/*
对于append()和copy函数
如果想增加切片的容量，需要创建容量更大的切片将原来的切片内容复制过来
用append向切片追加内容
*/
package main

import "fmt"

func printslice(x []int) {
	fmt.Printf("len:%d cap:%d slice:%v\n", len(x), cap(x), x) //输出格式d与D,v与V不同
}

func main() {
	var number []int
	printslice(number)

	number = append(number, 0)
	printslice(number)

	number = append(number, 1, 2, 3)
	printslice(number)

	//numbers:=make([]int,len(number),(cap(number))*2)
	numbers := make([]int, len(number), (cap(number))*2)
	copy(numbers, number)
	printslice(numbers)
	/*
	   cap的值取决于切片的前半部分到原切片末端的元素数
	   len的值取决于这个切片的元素数
	*/
	numbers1 := number[1:]
	printslice(numbers1)

	numbers1 = number[:3]
	printslice(numbers1)

}
