package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string][]string{
		"endpoints": {"api/v1/images [POST]", "api/v1/images/:imageID [GET]"},
	})
}
