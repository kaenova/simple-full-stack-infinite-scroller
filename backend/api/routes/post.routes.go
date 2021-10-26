package routes

import (
	"backend/api/controllers"

	"github.com/labstack/echo/v4"
)

func Post(e *echo.Echo) *echo.Echo {
	e.GET("/post", controllers.GetPost)
	// e.PUT("/routes1", ...)
	// e.DELETE("/routes1", ...)
	// e.POST("/tipe", controllers.PostTipeUser) //Creating Regular Account

	return e
}
