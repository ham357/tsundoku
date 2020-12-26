package app

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ham357/tsundoku/api/controllers/books"
)

func Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
	})

	b := e.Group("/books")
	{
		b.GET("/:book_id", books.GetBook())
		b.POST("", books.CreateBook())
		b.PUT("/:book_id", books.UpdateBook())
		b.PATCH("/:book_id", books.UpdateBook())
		b.DELETE("/:book_id", books.DeleteBook())
	}
}
