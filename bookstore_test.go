package bookstore_test

import (
	"bookstore"
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
		Title: "some title",
		Author: "some author",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Fatal("want error buying from zero copies, got nil")
	}
}
