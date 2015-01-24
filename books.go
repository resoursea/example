package main

import (
	"errors"
	"github.com/resoursea/api"
)

type Books []Book

func (bs *Books) GET(id *api.ID) (*Books, error) {
	err := db.Select(bs, "SELECT * FROM book")
	if err != nil {
		return nil, BookNotFoundError(errors.New("Error searching for books: " + err.Error()))
	}
	return bs, nil
}
