package main
//杨辉三角

import "fmt"

const LINES int = 10//定义常量

func showyanghuiTriangle() {
	nums:=[]int{}//定义一个空数组
	for i := 0; i < LINES; i++ {
		for j := 0; j < LINES-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value=1
			}else {
				value=nums[length-i]+nums[length-i-1]
			}
			nums=append(nums,value)
			fmt.Print(value," ")
		}
		fmt.Println("")
	}
}
func main() {
	showyanghuiTriangle()
}

