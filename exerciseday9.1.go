//指针数组
package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{1, 3, 4}
	var i int

	var ptr [MAX]*int //定义指针数组
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d]=%d\n", i, a[i])
		fmt.Println(*ptr[i])
	}
}
