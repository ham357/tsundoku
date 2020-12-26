package books

import (
	"github.com/ham357/tsundoku/api/datasources/mysql/books_db"
	"github.com/ham357/tsundoku/api/utils/errors"
	"github.com/ham357/tsundoku/api/utils/mysqlutils"
)

var (
	booksDB = make(map[uint]*Book)
)

const (
	noSearchResult = "record not found"
)

// Get - book
func (p *Book) Get() *errors.ApiErr {
	if result := books_db.Client.Where("id = ?", p.Model.ID).Find(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

// Save - book
func (p *Book) Save() *errors.ApiErr {
	// https://gorm.io/ja_JP/docs/error_handling.html
	if result := books_db.Client.Create(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

// Update - book
func (p *Book) Update() *errors.ApiErr {
	if result := books_db.Client.Save(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

// PartialUpdate - book
func (p *Book) PartialUpdate() *errors.ApiErr {
	if result := books_db.Client.
		Table("books").
		Where("id IN (?)", p.ID).
		Updates(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

// Delete - book
func (p *Book) Delete() *errors.ApiErr {
	if result := books_db.Client.Delete(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}
