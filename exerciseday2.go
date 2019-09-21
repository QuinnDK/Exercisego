package main
/*
如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同
 */
import "fmt"

func numbers() (int,int,string) {
	a,b,c:=1,3,"str"//未申明的对象可以用这种形式定义   相当于var a,b int=1,3       var c string="str"
	return a,b,c

}
func main()  {
	var a string = "string"
	fmt.Println(a)

	var c, b int = 1,3
	fmt.Println(c,b)

	_,numb,strs:=numbers()//main外定义函数定义numbers,numbers函数返回参数
/*
   空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
   _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，
   但有时你并不需要使用从一个函数得到的所有返回值。
*/
	fmt.Println(numb,strs)

}