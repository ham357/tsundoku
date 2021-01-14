package main

import (
    "github.com/ham357/tsundoku/api/app"
    "github.com/ham357/tsundoku/api/middlewares"
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.CORS())
    e.Use(middlewares.Firebase())
    app.Init(e)
    e.Logger.Fatal(e.Start(":8080"))
}
