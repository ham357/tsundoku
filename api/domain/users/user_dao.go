package users

import (
	"github.com/ham357/tsundoku/api/datasources/mysql/books_db"
	"github.com/ham357/tsundoku/api/utils/errors"
	"github.com/ham357/tsundoku/api/utils/mysqlutils"
)

var (
	usersDB = make(map[uint]*User)
)

const (
	noSearchResult = "record not found"
)

// Get - user
func (u *User) Find() *errors.ApiErr {
	if result := books_db.Client.Table("users").Where("uid = ?",  u.UID).First(&u); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

// Save - user
func (u *User) Save() *errors.ApiErr {
	// https://gorm.io/ja_JP/docs/error_handling.html
	if result := books_db.Client.Table("users").Create(&u); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}
