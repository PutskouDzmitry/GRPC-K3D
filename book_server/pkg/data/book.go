package data

import (
	"fmt"
	"gorm.io/gorm"
	"strings"

	dbConst "github.com/PutskouDzmitry/golang-training-library-grpc/book_server/pkg/const_db"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

//Entity in database
type Book struct {
	BookId            int64  // primary key
	AuthorId          int64  // foreign key
	PublisherId       int64  // foreign key
	NameOfBook        string // name of book
	YearOfPublication string // year of publication of the book
	BookVolume        int64  // book volume
	Number            int64  // number of book
}

//String output data in console
func (B Book) String() string {
	return fmt.Sprintln(B.BookId, B.AuthorId, B.PublisherId, strings.TrimSpace(B.NameOfBook), B.YearOfPublication, B.BookVolume, B.Number)
}

//BookData create a new connection
type BookData struct {
	db *gorm.DB // connection in db
}

//NewBookData it's imitation constructor
func NewBookData(db *gorm.DB) *BookData {
	return &BookData{db: db}
}

//ReadAll output all data with table books
func (B BookData) ReadAll() ([]Book, error) {
	log.Debug("Package data, method ReadAll")
	var books []Book
	result := B.db.Find(&books)
	log.Debug("Package data, method ReadAll. get data in db: ", result)
	if result.Error != nil {
		log.Error("We have an error :( ", result.Error)
		return nil, fmt.Errorf("can't read users from database, error: %w", result.Error)
	}
	log.Debug("Package data, method ReadAll, final")
	return books, nil
}

//Read read data in db
func (B BookData) Read(id int64) (Result, error) {
	var res Result
	result := B.db.Table(dbConst.Publishers).Select(dbConst.SelectBookAndPublisher).
		Joins(dbConst.ReadBookWithJoin).Where("book_id", id).Find(&res)
	if result.Error != nil {
		return Result{}, result.Error
	}
	return res, nil
}

//Add add data in db
func (B BookData) Add(book Book) (int64, error) {
	result := B.db.Create(&book)
	if result.Error != nil {
		return -1, fmt.Errorf(dbConst.CantAddDataError, result.Error)
	}
	return book.BookId, nil
}

//Update update number of books by the id
func (B BookData) Update(id int64, value int64) error {
	result := B.db.Table(dbConst.Books).Where(dbConst.BookId, id).Update("number", value)
	if result.Error != nil {
		return fmt.Errorf(dbConst.CantUpdateDataError, result.Error)
	}
	return nil
}

//Delete delete data in db
func (B BookData) Delete(id int64) error {
	result := B.db.Where(dbConst.BookId, id).Delete(&Book{})
	if result.Error != nil {
		return fmt.Errorf(dbConst.CantDeleteDataError, result.Error)
	}
	return nil
}
