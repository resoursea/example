package main

import (
	"errors"
)

// This resource represents a list of Books
// It forces the framework to catch an api.ID whenever a single Book was requested
type Books []Book

// [GET] /api/categories/:CategorySlug/books
// This method receives an Category resource or an error
// It return a list of books inside this ategory or an BookNotFoundError
// If no one Category was found by the Category creator method,
// so an error will be received and returned to the client
func (bs *Books) GET(cat *Category, err error) (*Books, error) {
	if err != nil {
		return nil, err
	}
	err = db.Select(bs, "SELECT * FROM book WHERE categoryid=?", cat.CategoryID)
	if err != nil {
		return nil, BookNotFoundError(errors.New("Error searching for books: " + err.Error()))
	}
	return bs, nil
}
