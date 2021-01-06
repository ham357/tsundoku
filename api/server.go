package main

import (
    "github.com/ham357/tsundoku/api/app"
    "github.com/ham357/tsundoku/api/middlewares"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.Use(middlewares.Firebase())
    app.Init(e)
    e.Logger.Fatal(e.Start(":8080"))
}
