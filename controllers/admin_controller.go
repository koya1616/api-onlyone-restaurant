package controllers

import (
	"koya/configs"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AdminLoginRequest struct {
	WhoAreYou string `json:"whoareyou"`
	Password  string `json:"password"`
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func AdminController(c echo.Context) error {
	var req AdminLoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.WhoAreYou != configs.EnvAdminName() || req.Password != configs.EnvAdminPassword() {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "認証に失敗しました"})
	}

	claims := &jwtCustomClaims{
		configs.EnvAdminName(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	trustyou, err := token.SignedString([]byte(configs.EnvJWTSecretKey()))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{"token": trustyou})
}
