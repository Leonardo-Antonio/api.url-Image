package main

import (
	"github.com/Leonardo-Antonio/api.images/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	router.Home(e)
	router.File(e)
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":" + port))
}
