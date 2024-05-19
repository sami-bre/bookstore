package bookstore_test

import(
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book {
		Title: "Spark joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}