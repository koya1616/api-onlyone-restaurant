package main

import (
	"koya/configs"
	"koya/routes"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{Validator: validator.New()}

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
