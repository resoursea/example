package main

import (
	"log"
	"time"

	"github.com/dchest/uniuri"
	"github.com/extemporalgenome/slug"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	dbname = "bookstore"
	dbuser = "app"
	dbpass = "SecretPassword!"
)

type DB struct {
	*sqlx.DB
}

// The only one DB instance
var db *DB = &DB{}

func init() {
	dbx, err := sqlx.Open("mysql", dbuser+":"+dbpass+"@/"+dbname+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalf("DB connection Error:", err.Error())
	}
	db.DB = dbx

	createCategoryTable(db)

	createBooksTable(db)
}

func createCategoryTable(db *DB) {
	_, err := db.Exec(categorySchema)
	if err != nil {
		log.Fatalf("Error creating the category schema. Err: %s\n", err.Error())
	}
	// Inserting the categories
	for i, categoryName := range categoryList {
		categorySlug := slug.Slug(categoryName)
		count := 0
		err := db.Get(&count, "select count(*) from category where categoryslug=?", categorySlug)
		if err != nil {
			log.Fatalf("Error searching for the category with categorySlug %s. Err: %s\n", categorySlug, err.Error())
		}

		if count == 0 {

			categoryId := uniuri.NewLen(20)

			if i == 0 { // "Uncategorized" is my default category
				categoryId = "default"
			}

			_, err := db.Exec("INSERT INTO category (categoryid, categoryname, categoryslug) VALUES (?,?,?)",
				categoryId, categoryName, categorySlug)
			if err != nil {
				log.Fatalf("Error when creating the category %s . Err: %s\n", categoryName, err)
			} else {
				log.Printf("Category %s created!\n", categoryName)
			}
		}
	}
}

func createBooksTable(db *DB) {
	_, err := db.Exec(booksSchema)
	if err != nil {
		log.Fatalf("Error creating the category schema. Err: %s\n", err.Error())
	}
	// Inserting the default book
	count := 0
	err = db.Get(&count, "select count(*) from book where bookid=?", "default")
	if err != nil {
		log.Fatalf("Error searching for the default book. Err: %s\n", err.Error())
	}

	b := bookDefault

	if count == 0 {
		_, err := db.Exec("INSERT INTO book (bookid, categoryid, title, slug, description, likecount, creation, lastupdate, deleted) VALUES (?,?,?,?,?,?,?,?,?)",
			b.BookID, b.CategoryID, b.Title, b.Slug, b.Description, b.LikeCount, b.Creation, b.LastUpdate, b.Deleted)
		if err != nil {
			log.Fatalf("Error when creating the default book. Err: %s\n", err)
		} else {
			log.Printf("Default book created!\n")
		}
	}
}

var categoryList = []string{
	"Uncategorized",
	"Animals",
	"Art and Culture",
	"Beauty and Style",
	"Cars and Motorbikes",
	"Interior Design",
	"Science and Technology",
	"Food & Drink",
	"Curiosities",
	"Education",
	"Entertainment",
	"Sport",
	"Events",
	"Movies",
	"Photos",
	"Football",
	"Humor",
	"Internet",
	"Games",
	"Woman",
	"Music",
	"Business",
	"News",
	"People and Blogs",
	"Health",
	"Videos",
}

var bookDefault = Book{
	BookID:      "default",
	CategoryID:  "default",
	Title:       "Default Book!",
	Slug:        "default-book",
	Description: "Default Book description",
	LikeCount:   0,
	Creation:    time.Unix(5239592839, 1041),
	LastUpdate:  time.Unix(5239592839, 1041),
	Deleted:     false,
}

const categorySchema = ` 
CREATE TABLE IF NOT EXISTS category (
  categoryid VARCHAR(40) NOT NULL,
  categoryname VARCHAR(40) NOT NULL,
  categoryslug VARCHAR(40) NOT NULL,
  PRIMARY KEY (categoryid),
  UNIQUE INDEX categoryslug_UNIQUE (categoryslug ASC))
`

const booksSchema = `
CREATE TABLE IF NOT EXISTS book (
  bookid VARCHAR(20) NOT NULL,
  categoryid VARCHAR(20) NOT NULL,
  title VARCHAR(255) NOT NULL,
  slug VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL DEFAULT '',
  likecount INT NULL DEFAULT 0,
  creation DATETIME NOT NULL,
  lastupdate DATETIME NOT NULL,
  deleted TINYINT(1) NULL DEFAULT 0,
  PRIMARY KEY (bookid),
  UNIQUE INDEX channelid_slug_unique (categoryid ASC, slug ASC),
  CONSTRAINT book_category
    FOREIGN KEY (categoryid)
    REFERENCES category (categoryid)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
`
