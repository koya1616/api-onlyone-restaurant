package response

type PendingRestaurantResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Information string `json:"information"`
	Isapproved  bool   `json:"isapproved"`
}
