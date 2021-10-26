package controllers

import (
	"backend/api/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetPost(c echo.Context) error {
	var (
		id string

		res models.Response = models.CreateResponse()
		err error
	)

	id = c.QueryParam("id")

	if strings.TrimSpace(id) == "" {
		res.Data, err = models.PostGetAll()
	} else {
		res.Data, err = models.PostSearchID(id)
	}

	if err != nil {
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	return c.JSON(res.Status, res)
}
