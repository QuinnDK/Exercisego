package main

import "fmt"

func main() {
	num :=[6]float32{5,8,9,4}//不足位自动补零

	for i, k:=range num{
		fmt.Printf("第 %d 位置为 %f\n",i,k)//默认小数点后六位    不足补零
		fmt.Printf("第 %d 位置为 %0.1f\n",i,k)//小数点后一位
		fmt.Printf("第 %d 位置为 %d\n",i,k)//不晓得是什么
		fmt.Printf("第 %f 位置为 %f\n",i,k)//三四行格式错误输出
		
	}
}
