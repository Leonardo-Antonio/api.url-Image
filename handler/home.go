package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string][]string{
		"endpoints": {
			"https://api-url-images.herokuapp.com/api/v1/images [POST]",
			"https://api-url-images.herokuapp.com/api/v1/images/:imageID [GET]",
		},
	})
}
