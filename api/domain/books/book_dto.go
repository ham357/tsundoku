package books

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/ham357/tsundoku/api/utils/errors"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Price  uint64 `json:"price"`
	Img    []byte `json:"img"`
}

func (b *Book) Validate() *errors.ApiErr {
	b.Name = strings.TrimSpace(strings.ToLower(b.Name))
	if b.Name == "" {
		return errors.NewBadRequestError("invalid product name")
	}
	return nil
}
