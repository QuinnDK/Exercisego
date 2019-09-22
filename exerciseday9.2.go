//学习结构体，感觉和c语言没有区别
package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.author = "董坤"
	Book1.title = "学习go"
	Book1.book_id = 1
	Book1.subject = "语言"

	Book2.author = "shax"
	Book2.title = "学习python"
	Book2.book_id = 2
	Book2.subject = "语言"

	printbook(Book1)
	printbook(Book2)
	fmt.Println("---------------------------------------------")

	printbook2(&Book1)
	printbook2(&Book2)

}
func printbook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
func printbook2(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
