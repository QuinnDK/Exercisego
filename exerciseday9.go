//go语言指针
package main

import "fmt"

func main() {
	var a int = 10
	var ip *int
	fmt.Printf("变量的地址：%x\n", &a)
	fmt.Println("变量的地址：", &a)
	ip = &a
	fmt.Println("ip 变量存储的指针地址:", ip)
	fmt.Println("ip 变量存储的指针地址的值:", *ip)
	fmt.Println("ip 变量存储的指针地址的地址:", &ip)
	var ptr *int
	if ptr != nil {
		if ip != nil {
			fmt.Println("ptr不是空指针")
			fmt.Println("ip不是空指针")
		} else {
			fmt.Println("ptr不是空指针")
			fmt.Println("ip是空指针")
		}
	} else {
		if ip != nil {
			fmt.Println("ptr是空指针")
			fmt.Println("ip不是空指针")
		} else {
			fmt.Println("ptr是空指针")
			fmt.Println("ip是空指针")
		}
	}

	//仅在 switch 处会对大括号内的各个 case 标签进行判断，而使用了 fallthrough 后 case 标签的内容将被无视
	// 无论条件是什么，都仅执行下一条 case 的内容
	switch {
	case ptr != nil:
		fmt.Println("ptr不是空指针")
		fallthrough
	case ptr == nil:
		fmt.Println("ptr是空指针")
		fallthrough
	case ip != nil:
		fmt.Println("ip不是空指针")
	default:
		fmt.Println("ip是空指针")
	}
}
