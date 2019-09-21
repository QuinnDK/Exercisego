package main

import "fmt"

func putstar() {
	x:=9
	y:=9
	row:=1
	for row < y  {
		count:=0
		if  row<(y/2+1) {
			count = 2*row + 1
		}else {
			count=2*(y-row)+1
		}
		row++
		text:=""
		star_min:=((x-row)/2)+1
		star_max:=((x-row)/2)+row
		for index := 0; index <= x; index++ {
			if  index>star_min && index<star_max{
				text+="*"
			}else {
				text+=""
			}
		}
		fmt.Println(text)
	}
}
func main() {
	putstar()
}
