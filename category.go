package main

import (
	"errors"
	"github.com/resoursea/api"
)

// This Resource classifies the Books resource in categories
type Category struct {
	CategoryID   string
	CategoryName string
	CategorySlug string

	Books *Books
}

// This resource is represents a list of Categories
// It forces the framework to catch an api.ID whenever a single Category was requested
type Categories []Category

// This error is returned when it can not recover the required category from the database
// It will enter in the pipeline in order be treated by any subsequent method with is requesting for the errors
type CategoryNotFoundError error

// The creator method for the Category resource
// This method will be called whenever the resource is requested
// It inject one Category or one CategoryNotFoundError in the pipeline
func (c *Category) Init(id *api.ID, db *DB) (*Category, error) {
	// api.ID will never be null
	// Because all routes that requires the Category resource are children of Category
	// So to access them the client should use an route like: /api/category/:CategorySlug/...
	// and the framework will inject the :CategorySlug as the api.ID requested by the Category resource
	err := db.Get(c, "SELECT * FROM category WHERE categoryslug=?", id.String())
	if err != nil {
		return nil, CategoryNotFoundError(errors.New("Category '" + id.String() + "' not found!"))
	}
	return c, nil
}
