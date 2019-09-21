package main

import "fmt"

func changc(num []int) {
	num[0]=100
}
func main() {
	var num  =[]int{1,3,5}
	changc(num)
	fmt.Println(num)//输出的数组是带有括号的的，原来的数组已经被函数所修改 ？
}
