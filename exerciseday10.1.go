//数组与切片的联系与切片
/*
1.切片是指针类型，数组是值类型
2.数组长度固定，切片是动态数组
3.切片比数组多一个cap(容量)类型
4.切片的底层是数组
*/
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}     //定义切片
	numbers1 := [...]int{1, 2, 3, 4, 5, 6} //在声明数组时你可以忽略数组的长度并用 ... 代替
	printnubers(numbers)
	printnubers1(numbers1)

}
func printnubers(numbers []int) {
	for i, e := range numbers { //range for返回数组的索引和range 索引对应的值    i:索引，e:值
		if i == len(numbers)-1 {
			numbers[0] += e
		} else {
			numbers[i+1] += e
		}
	}
	fmt.Println(numbers)
}
func printnubers1(numbers [6]int) {
	for i, e := range numbers {
		if i == len(numbers)-1 {
			numbers[0] += e
		} else {
			numbers[i+1] += e
		}
	}
	fmt.Println(numbers)
}
