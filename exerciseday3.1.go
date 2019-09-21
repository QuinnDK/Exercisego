/*
常量可以用作枚举
const{
Unknow=0
Female=1
Male=2
}
 */
//iota特殊常量（可以被编译器修改的常量）

package main

import "unsafe"

const(
	a="abc"
	b=len(a)
	c=unsafe.Sizeof(a)//字符串类型在go里是个结构，包含只想底层的指针和长度，这两部分各
	//8个字节，so字符串类型为16个字节
	k
	j//kj没有初始化使用上一行（c）的值
	)
const(
	d=iota
	e
	f
)
//第一个iota等于0，在新的一行被使用时，他的值恢复为本来的计数
//so d=0,e=1,f=2
//每当有新的const关键字时，iota计数会重新开始
func main(){
	println(a,b,c,k,j)
	println(d,e,f)
}