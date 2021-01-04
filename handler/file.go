package handler

import (
	"github.com/Leonardo-Antonio/api.images/helper"
	"github.com/aidarkhanov/nanoid/v2"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type File struct{}

func NewFile() *File {
	return &File{}
}

func (f *File) Save(c echo.Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response := helper.NewResponseJSON("ERROR", err.Error(), true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	file, err := fileHeader.Open()
	if err != nil {
		response := helper.NewResponseJSON("ERROR", err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		response := helper.NewResponseJSON("ERROR", err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	ID, err := nanoid.New()
	if err != nil {
		response := helper.NewResponseJSON("ERROR", err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	fileName := strings.Split(fileHeader.Filename, " ")
	url := ID + strings.Join(fileName, "-")
	err = ioutil.WriteFile(
		"public/images/"+url,
		fileBytes, os.ModePerm,
	)
	if err != nil {
		response := helper.NewResponseJSON("ERROR", err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.NewResponseJSON(
		"MESSAGE",
		"the file was saved successfully",
		false,
		"https://api-url-images.herokuapp.com/api/v1/images/"+url,
	)
	return c.JSON(http.StatusCreated, response)
}
