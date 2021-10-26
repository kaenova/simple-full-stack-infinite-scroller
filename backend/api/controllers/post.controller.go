package controllers

import (
	"backend/api/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetPost(c echo.Context) error {
	var (
		page     int             = 1
		pageSize int             = 10
		res      models.Response = models.CreateResponse()
		err      error
	)

	pageCtx := c.QueryParam("page")
	pageSizeCtx := c.QueryParam("page_size")

	if strings.TrimSpace(pageCtx) != "" {
		page, err = strconv.Atoi(pageCtx)
		if err != nil {
			page = 1
		}
	}

	if strings.TrimSpace(pageSizeCtx) != "" {
		pageSize, err = strconv.Atoi(pageSizeCtx)
		if err != nil {
			pageSize = 20
		}
	}

	if pageSize > 35 {
		pageSize = 35
	}

	res.Data, err = models.PostGetPage(page, pageSize)

	if err != nil {
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	return c.JSON(res.Status, res)
}
