package main

import "fmt"

func main()  {
	var a int = 4
	var b int32
	var c float32
	var ptr *int
//运算符实例
	fmt.Printf("第一行-a变量类型为=%T\n",a)
	fmt.Printf("第二行-b变量类型为=%T\n",b)
	fmt.Printf("第二行-c变量类型为=%T\n",c)

//& 和* 实例
	ptr=&a
	a++
	fmt.Printf("a的值为%d\n",a)
	fmt.Printf("*ptr为%d\n",*ptr)//需要格式化的输出用printf,其他的用println
	fmt.Println(ptr)
}