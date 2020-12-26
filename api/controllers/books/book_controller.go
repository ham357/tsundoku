package books

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/ham357/tsundoku/api/domain/books"
	"github.com/ham357/tsundoku/api/services"
	"github.com/ham357/tsundoku/api/utils/errors"
)

func CreateBook() echo.HandlerFunc{
	return func (c echo.Context) error {
		var book books.Book
		if err := c.Bind(&book); err != nil {
			apiErr := errors.NewBadRequestError("invalid json body")
			return c.JSON(apiErr.Status, apiErr)
		}
		newBook, saveErr := services.CreateBook(book)
		if saveErr != nil {
			return c.JSON(saveErr.Status, saveErr)
		}
		return c.JSON(http.StatusCreated, newBook)
	}
}

// GetBook - Get book by book id
func GetBook() echo.HandlerFunc{
	return func (c echo.Context) error {
		bookID, idErr := getBookID(c.Param("book_id"))
		if idErr != nil {
			return c.JSON(idErr.Status, idErr)
		}

		book, getErr := services.GetBook(bookID)
		if getErr != nil {
			return c.JSON(getErr.Status, getErr)
		}
		return c.JSON(http.StatusOK, book)
	}
}

// UpdateBook - Update book
func UpdateBook() echo.HandlerFunc{
	return func (c echo.Context) error {
		bookID, idErr := getBookID(c.Param("book_id"))
		if idErr != nil {
			return c.JSON(idErr.Status, idErr)
		}

		var book books.Book
		if err := c.Bind(&book); err != nil {
			apiErr := errors.NewBadRequestError("invalid json body")
			return c.JSON(apiErr.Status, apiErr)
		}

		book.ID = bookID

		isPartial := c.Request().Method == http.MethodPatch

		result, err := services.UpdateBook(isPartial, book)
		if err != nil {
			return c.JSON(err.Status, err)
		}

		return c.JSON(http.StatusOK, result)
	}
}

func getBookID(bookIDParam string) (uint, *errors.ApiErr) {
	bookID, bookErr := strconv.ParseUint(bookIDParam, 10, 64)
	if bookErr != nil {
		return 0, errors.NewBadRequestError("book id should be a number")
	}
	return uint(bookID), nil
}

// DeleteBook - Delete book
func DeleteBook() echo.HandlerFunc{
	return func (c echo.Context) error {
		bookID, idErr := getBookID(c.Param("book_id"))
		if idErr != nil {
			return c.JSON(idErr.Status, idErr)
		}

		if err := services.DeleteBook(bookID); err != nil {
			return c.JSON(err.Status, err)
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
	}
}
