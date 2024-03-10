package controllers

import (
	"koya/configs"
	"koya/models"
	"koya/request"
	"koya/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPendingRestaurants(c echo.Context) error {
	var pending_restaurants []models.PendingRestaurant

	result := configs.GetNeonDB().Unscoped().Order("id desc").Find(&pending_restaurants)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "データの取得中にエラーが発生しました"})
	}

	var pending_restaurant_responses []response.PendingRestaurantResponse
	for _, pr := range pending_restaurants {
		pending_restaurant_responses = append(pending_restaurant_responses, response.PendingRestaurantResponse{
			ID:          pr.ID,
			Name:        pr.Name,
			Information: pr.Information,
			Isapproved:  pr.Isapproved,
		})
	}

	return c.JSON(http.StatusOK, &echo.Map{"pendingRestaurants": pending_restaurant_responses})
}

func ApprovePendingRestaurants(c echo.Context) error {
	var req request.ApprovePendingRestaurantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストの形式が正しくありません"})
	}

	var pending_restaurant models.PendingRestaurant
	result := configs.GetNeonDB().Unscoped().Where("id = ?", req.ID).First(&pending_restaurant)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "データの取得中にエラーが発生しました"})
	}

	pending_restaurant.Isapproved = true
	result = configs.GetNeonDB().Unscoped().Save(&pending_restaurant)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "データの更新中にエラーが発生しました"})
	}

	return c.JSON(http.StatusOK, &echo.Map{"message": "承認しました"})
}