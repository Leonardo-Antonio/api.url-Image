package router

import (
	"github.com/Leonardo-Antonio/api.images/handler"
	"github.com/labstack/echo/v4"
)

func File(e *echo.Echo) {
	hand := handler.NewFile()
	g := e.Group("/api/v1/images")
	g.POST("", hand.Save)
	g.Static("", "public/images")
}
