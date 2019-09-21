package main

import "fmt"

func geteverage(arr [5]int,size int)float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum+=arr[i]
	}
	avg=float32(sum)/float32(size)
	return avg//返回的参数需指定类型，才不会报错

}

func main() {
	var balance=[5]int{5,6,8,4,8}
	avger:=geteverage(balance,5)
	fmt.Println(avger)
	fmt.Printf("%d",int(avger))//无法通过格式化输出想要的格式，需提前进行强制转换

}
