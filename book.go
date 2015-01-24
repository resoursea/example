package main

import (
	"errors"
	"fmt"
	"github.com/dchest/uniuri"
	"time"

	"github.com/extemporalgenome/slug"
	"github.com/resoursea/api"
)

type BookNotFoundError error

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

func (b *Book) Init(id *api.ID) (*Book, error) {
	if id == nil {
		return nil, nil
	}
	err := db.Get(b, "SELECT * FROM book WHERE slug=?", id.String())
	if err != nil {
		return nil, BookNotFoundError(errors.New("Error searching for the book '" + id.String() + "': " + err.Error()))
	}
	return b, nil
}

func (b *Book) GET(err error) (*Book, error) {
	return b, err
}

func (b *Book) POSTLike(err error) (*Book, error) {
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

	return b, nil //fmt.Errorf("Error inserting the book!")
}
