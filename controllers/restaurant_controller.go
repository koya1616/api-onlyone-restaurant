package controllers

import (
	"fmt"
	"koya/configs"
	"koya/models"
	"koya/request"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var restaurantCollection *mongo.Collection = configs.GetCollection(configs.DB, "restaurants")
var validate = validator.New()

// func CreateUser(c echo.Context) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	var user models.User
// 	defer cancel()

// 	//validate the request body
// 	if err := c.Bind(&user); err != nil {
// 		return c.JSON(http.StatusBadRequest, responses.RestaurantResponse{Status: http.StatusBadRequest, Message: "error", Restaurant: &echo.Map{"data": err.Error()}})
// 	}

// 	//use the validator library to validate required fields
// 	if validationErr := validate.Struct(&user); validationErr != nil {
// 		return c.JSON(http.StatusBadRequest, responses.RestaurantResponse{Status: http.StatusBadRequest, Message: "error", Restaurant: &echo.Map{"data": validationErr.Error()}})
// 	}

// 	newUser := models.User{
// 		Id:       primitive.NewObjectID(),
// 		Name:     user.Name,
// 		Location: user.Location,
// 		Title:    user.Title,
// 	}

// 	result, err := userCollection.InsertOne(ctx, newUser)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.RestaurantResponse{Status: http.StatusInternalServerError, Message: "error", Restaurant: &echo.Map{"data": err.Error()}})
// 	}

// 	return c.JSON(http.StatusCreated, responses.RestaurantResponse{Status: http.StatusCreated, Message: "success", Restaurant: &echo.Map{"data": result}})
// }

func GetRestaurant(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var restaurant models.Restaurant
	defer cancel()

	err := restaurantCollection.FindOne(ctx, bson.M{"id": c.Param("id")}).Decode(&restaurant)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, &echo.Map{"error": fmt.Sprintf("IDが %s のお店ないよ！！", c.Param("id"))})
		} else {
			return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusOK, &echo.Map{"restaurant": restaurant})
}

func GetRestaurants(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var restaurants []models.Restaurant
	defer cancel()

	results, err := restaurantCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleRestaurant models.Restaurant
		if err = results.Decode(&singleRestaurant); err != nil {
			return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
		}

		restaurants = append(restaurants, singleRestaurant)
	}

	return c.JSON(http.StatusOK, &echo.Map{"restaurants": restaurants})
}

func RequestRestaurant(c echo.Context) error {
	var input request.RequestRestaurant
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{"error": err.Error()})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{"error": err.Error()})
	}

	restaurant := models.PendingRestaurant{Name: input.Name, Information: input.Information}

	result := configs.GetNeonDB().Create(&restaurant)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{"error": result.Error})
	}

	return c.JSON(http.StatusCreated, &echo.Map{"message": "投稿ありがとう！"})
}