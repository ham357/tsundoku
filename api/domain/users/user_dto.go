package users

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/ham357/tsundoku/api/utils/errors"
)

type User struct {
	gorm.Model
	UID       string     `json:"-"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
}

func (u *User) Validate() *errors.ApiErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid Email")
	}
	return nil
}
