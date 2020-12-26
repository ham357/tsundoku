package main

import (
	"github.com/ham357/tsundoku/api/app"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    app.Init(e)
    e.Logger.Fatal(e.Start(":8080"))
}
