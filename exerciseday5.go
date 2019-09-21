package main

import "fmt"

func main() {
	print9x()
	gototag()
}
//嵌套for循环打印九九乘法表
func print9x()  {
	for m:=1;m<10;m++{
		for n := 1; n <= m; n++ {
			fmt.Printf("%d*%d=%d ",m,n,m*n)
		}
		fmt.Println("")

	}
}
//for循环配合goto打印九九乘法表
func gototag() {
	for m := 1; m < 10; m++ {
		n := 1
	LOOP: if n <= m {
		fmt.Printf("%dx%d=%d ",n,m,m*n)
		n++
		goto LOOP
	} else {
		fmt.Println("")
	}
		n++
	}
}
