package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	var b bookstore.Book = bookstore.Book{
		Title:  "some title",
		Author: "somebody",
		Copies: 12,
	}

	var want int = 11

	result, err := bookstore.Buy(b)

	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies

	if want != got {
		t.Errorf("Buy() failed to produce the correct number of copies. Want: %d, Got %d:", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "some title",
		Author: "some author",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Fatal("want error buying from zero copies, got nil")
	}
}

func TestShowAllBooks(t *testing.T) {
	t.Parallel()

	catalog := []bookstore.Book{
		{Title: "For the love of go"},
		{Title: "The power of go: Tools"},
	}

	want := []bookstore.Book{
		{Title: "For the love of go"},

		{Title: "The power of go: Tools"},
	}

	got := bookstore.GetAllBooks(catalog)

	// compare got and catalog
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}


func TestGetBook(t *testing.T) {
	t.Parallel()
	books := []bookstore.Book {
		{Id: 1, Title: "For the love of go"},
		{Id: 2, Title: "Go powered"},
		{Id: 3, Title: "The way to go"},
	}

	var want bookstore.Book = bookstore.Book{Id: 3, Title: "The way to go"}
	got := bookstore.GetBook(books, 3)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
