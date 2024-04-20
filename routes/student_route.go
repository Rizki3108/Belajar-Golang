package routes

import (
	"Materi-Rest-API/controller"

	"github.com/labstack/echo/v4"
)

func StudentRoute(e *echo.Echo) {
	e.POST("/student", controller.CreateStudent)
}