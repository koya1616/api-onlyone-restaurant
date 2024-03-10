package routes

import (
	"github.com/labstack/echo/v4"
	"koya/controllers"
)

func RestaurantRoute(e *echo.Echo) {
	e.GET("/restaurants/:id", controllers.GetRestaurant)
	e.GET("/restaurants", controllers.GetRestaurants)
	e.POST("/restaurant/request", controllers.RequestRestaurant)
}
