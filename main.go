package main

import (
	"koya/configs"
	"koya/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:7776", "https://www.onlyone-restaurant.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	configs.ConnectDB()
	configs.InitNeonDB()

	routes.RestaurantRoute(e)
	routes.PendingRestaurantRoute(e)
	routes.AdminRoute(e)

	e.Logger.Fatal(e.Start(":" + configs.EnvPort()))
}
