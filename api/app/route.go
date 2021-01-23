package app

import (
	"net/http"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/ham357/tsundoku/api/controllers/books"
	"github.com/ham357/tsundoku/api/middlewares"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func Init(e *echo.Echo) {
	// tmpl := template.Must(template.ParseGlob("api/views/*.html"))
	t := &Template{
			templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
	})

	b := e.Group("/books", middlewares.FirebaseGuard())
	{
		b.GET("/:book_id", books.GetBook())
		b.POST("", books.CreateBook())
		b.PUT("/:book_id", books.UpdateBook())
		b.PATCH("/:book_id", books.UpdateBook())
		b.DELETE("/:book_id", books.DeleteBook())
	}

	u := e.Group("/users")
	{
		u.GET("/signup", func(c echo.Context) error {
			return c.Render(http.StatusOK, "signup.html", "")
		})
		u.GET("/success", func(c echo.Context) error {
			return c.Render(http.StatusOK, "success.html", "")
		})
	}
}
