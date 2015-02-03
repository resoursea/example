package main

import (
	"errors"
	"fmt"
	"github.com/dchest/uniuri"
	"time"

	"github.com/extemporalgenome/slug"
	"github.com/resoursea/api"
)

// This is a representation of the Book Resource
type Book struct {
	BookID      string
	CategoryID  string
	Title       string
	Slug        string
	Description string
	LikeCount   int
	Creation    time.Time
	LastUpdate  time.Time
	Deleted     bool
}

// This error is returned when it can not recover the required book from the database
type BookNotFoundError error

// The creator for the Book Resource
// This method will be called whenever the Book resource is requested,
// it constructs the Book using the ID passed in the URI
// This method receives the the resources: DB, Category, ID or an error
// The book Resource is inside an category, so it should be injected to identify the Book the client is requesting
// It inject one Book or one BookNotFoundError in the pipeline
func (b *Book) New(db *DB, cat *Category, id api.ID, err error) (*Book, error) {
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, nil
	}
	err = db.Get(b, "SELECT * FROM book WHERE slug=? AND categoryid=?", id.String(), cat.CategoryID)
	if err != nil {
		return nil, BookNotFoundError(errors.New("Error searching for the book '" + id.String() + "': " + err.Error()))
	}
	return b, nil
}

// [GET] /api/categories/:CategorySlug/books/:BookID
// It receives a Book and an error and return it to the client
// When the book is requested the framework will call the Book's Creator method (Book.New)
// It's creator can inject or the Book itself, or an error in the pipeline
// This method just receives and returns it to the client
func (b *Book) GET(err error) (*Book, error) {
	return b, err
}

// [GET] /api/categories/:CategorySlug/books
// It receives the resources BookPOST, DB, Category or an error
// The BookPOST resource will retrieve the data sent by the client
// This method will be used by the client to add a new Book to this category
// If any error was injected, it just return it to the client
// If nothing goes wrong, this method return the Book saved in the database
func (_ *Books) POST(data *BookPOST, db *DB, cat *Category, err error) (*Book, error) {
	if err != nil {
		return nil, err
	}

	b := &Book{
		BookID:      uniuri.NewLen(5),
		CategoryID:  cat.CategoryID,
		Title:       data.Title,
		Slug:        slug.Slug(data.Title),
		Description: data.Description,
		LikeCount:   0,
		Creation:    time.Now(),
		LastUpdate:  time.Now(),
		Deleted:     false,
	}

	_, err = db.Exec("INSERT INTO book (bookid, categoryid, title, slug, description, likecount, creation, lastupdate, deleted) VALUES (?,?,?,?,?,?,?,?,?)",
		b.BookID, b.CategoryID, b.Title, b.Slug, b.Description, b.LikeCount, b.Creation, b.LastUpdate, b.Deleted)
	if err != nil {
		return nil, fmt.Errorf("Error inserting the book: %s", err.Error())
	}

	return b, nil
}

// [POST] /api/categories/:CategoryID/books/:BookID/like
// It receives a Book or an error and increments the like's counter of this book
// This method is an action that the Book resource can perform
// Actions are accessible by the route: [METHOD] resource/action
// Remember the Book resource can't have any children with this Action's name 'like'
// because it will conflict with this Action's route
func (b *Book) POSTLike(db *DB, err error) (*Book, error) {
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("UPDATE book SET likecount=likecount+1 WHERE bookid=?", b.BookID)
	if err != nil {
		return nil, fmt.Errorf("Error updating the book like: %s", err.Error())
	}
	b.LikeCount += 1
	return b, err
}
