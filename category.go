package main

import (
	"errors"
	"github.com/resourcerest/api"
)

type CategoryNotFoundError error

type Category struct {
	CategoryID   string
	CategoryName string
	CategorySlug string

	Books *Books
}

type Categories []Category

func (c *Category) Init(id *api.ID, db *DB) (*Category, error) {
	err := db.Get(c, "SELECT * FROM category WHERE categoryslug=?", id.String())
	if err != nil {
		return nil, CategoryNotFoundError(errors.New("Category '" + id.String() + "' not found!"))
	}
	return c, nil
}
