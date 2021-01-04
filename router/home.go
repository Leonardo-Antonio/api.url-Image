package router

import (
	"github.com/Leonardo-Antonio/api.images/handler"
	"github.com/labstack/echo/v4"
)

func Home(e *echo.Echo) {
	e.GET("", handler.Home)
}
