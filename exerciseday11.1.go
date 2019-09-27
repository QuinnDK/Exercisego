package main

import "fmt"

func main() {
	sliceTest()
	twoDimensionArray()
}

func sliceTest() {
	arr := []int{1, 2, 3, 4, 5}
	s := arr[:]
	for e := range s { //_,e,e为切片s中各个元素，e或者e,_表示索引值
		fmt.Println(s[e])
	}
	s1 := make([]int, 3)
	for e, _ := range s1 {
		fmt.Println(s1[e])
	}
}

func twoDimensionArray() {
	/* 数组 - 5 行 2 列*/
	var a = [][]int{{0, 0}, {1, 2}, {2}, {3, 6}, {4, 8}}
	var i, j int
	fmt.Println("------------------------------")
	fmt.Printf("%d", len(a))
	fmt.Println("\n")
	/* 输出数组元素 */
	for i = 0; i < len(a); i++ {
		fmt.Printf("%d\n", a[i]) //与c语言二维数组相似，先遍历行，a[i]表示这一行的所有元素
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
