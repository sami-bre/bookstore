package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
	Id     int
}

func Buy(book Book) (Book, error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	book.Copies--
	return book, nil
}

func GetAllBooks(books []Book) []Book {
	return books
}

func GetBook(books map[int]Book, id int) Book {
	return books[id]
}
