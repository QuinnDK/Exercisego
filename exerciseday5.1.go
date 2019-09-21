package main

import "fmt"

func main()  {
	PtintPrime()
	PrintPrime1()
}
func PtintPrime()  {
	var C, c int
	C=1//以免goto时每次都初始化C
	A:for C<100  {
		C++
		for c = 2; c < C/c; c++ {
			if C % c == 0 {
				goto A
			}
		}
		fmt.Println(C,"是素数")
	}
}
func PrintPrime1() {
	var a, b int
	for a = 2; a <= 100; a++ {
		for b = 2; b <= (a / b); b++ {
			if a%b == 0 {
				break
			}
		}
		if b > (a / b) {
			fmt.Printf("%d\t是素数\n", a)
		}
	}
}
