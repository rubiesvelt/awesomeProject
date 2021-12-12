package goLearning

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func structTest() {
	book := Books{"boob", "Bob", "boom", 10}
	ptr := &book

	var nptr *Books
	nptr = &book

	printBooks(book)
	printBooksPtr(ptr)
	printBooksPtr(nptr)
}

func printBooks(books Books) {
	fmt.Printf("book's title is %s", books.title)
	fmt.Printf("book's auther is %s", books.author)
	fmt.Printf("book's subject is %s", books.subject)
	fmt.Printf("book's book_id is %d", books.book_id)
}

func printBooksPtr(books *Books) {
	fmt.Printf("book's title is %s\n", books.title)
	fmt.Printf("book's auther is %s\n", books.author)
	fmt.Printf("book's subject is %s\n", books.subject)
	fmt.Printf("book's book_id is %d\n", books.book_id)
}
