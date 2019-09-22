package main

import "fmt"

func changc(num []int) {
	num[0] = 100
}

func prt(arr [][]float32) { //传入二位数组做参数
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i][0])
	}
}
func main() {
	var num = []int{1, 3, 5}
	changc(num)
	fmt.Println(num) //输出的数组是带有括号的的，原来的数组已经被函数所修改 ？
	var arr = [][]float32{{-1, -2}, {-3, -4}, {-5}}
	prt(arr)
}
