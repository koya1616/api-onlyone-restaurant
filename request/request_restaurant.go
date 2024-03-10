package request

type RequestRestaurant struct {
	Name        string `json:"name" validate:"required"`
	Information string `json:"information" validate:"required"`
}