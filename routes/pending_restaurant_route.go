package routes

import (
	"koya/configs"
	"koya/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt/v4"
)

func PendingRestaurantRoute(e *echo.Echo) {
	r := e.Group("/pending-restaurants")

	r.Use(echojwt.JWT([]byte(configs.EnvJWTSecretKey())))

	r.GET("", controllers.GetPendingRestaurants)
	r.POST("/approve", controllers.ApprovePendingRestaurants)
}
