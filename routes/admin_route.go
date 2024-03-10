package routes

import (
	"github.com/labstack/echo/v4"
	"koya/controllers"
)

func AdminRoute(e *echo.Echo) {
	e.POST("/admin/login", controllers.AdminController)
}
