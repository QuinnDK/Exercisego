package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func changebooktitle(books Books) {
	books.title = "gogogo"
}
func changebooktitle1(books *Books) {
	books.title = "gogogo"
}
func main() {
	var book1 Books
	book1.title = "book1"
	book1.author = "zuozhe"
	book1.book_id = 1
	changebooktitle(book1)
	fmt.Println(book1)
	changebooktitle1(&book1)
	fmt.Println(book1)
}
