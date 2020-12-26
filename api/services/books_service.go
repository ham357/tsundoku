package services

import (
	"github.com/jinzhu/gorm"
	"github.com/ham357/tsundoku/api/domain/books"
	"github.com/ham357/tsundoku/api/utils/errors"
)

// CreateBook - Service
func CreateBook(book books.Book) (*books.Book, *errors.ApiErr) {
	if err := book.Validate(); err != nil {
		return nil, err
	}

	if err := book.Save(); err != nil {
		return nil, err
	}

	return &book, nil
}

func GetBook(bookID uint) (*books.Book, *errors.ApiErr) {
	b := &books.Book{Model: gorm.Model{ID: bookID}}
	if err := b.Get(); err != nil {
		return nil, err
	}
	return b, nil
}

// UpdateBook - Service
func UpdateBook(isPartial bool, book books.Book) (*books.Book, *errors.ApiErr) {
	current, err := GetBook(book.ID)
	if err = current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if book.Name != "" {
			current.Name = book.Name
		}
		if book.Detail != "" {
			current.Detail = book.Detail
		}
		if book.Price != 0 {
			current.Price = book.Price
		}
		if book.Img != nil {
			current.Img = book.Img
		}
		if err := current.PartialUpdate(); err != nil {
			return nil, err
		}
	} else {
		// Change Book Info
		current.Name = book.Name
		current.Detail = book.Detail
		current.Price = book.Price
		current.Img = book.Img
		if err := current.Update(); err != nil {
			return nil, err
		}
	}
	return current, nil
}

// DeleteBook - Service
func DeleteBook(bookID uint) *errors.ApiErr {
	book := &books.Book{Model: gorm.Model{ID: bookID}}
	return book.Delete()
}
