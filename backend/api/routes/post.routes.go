package routes

import (
	"backend/api/controllers"

	"github.com/labstack/echo/v4"
)

func Post(e *echo.Echo) *echo.Echo {
	e.GET("/post", controllers.GetPost)
	// e.POST("/post", controllers.PostPost)

	return e
}
