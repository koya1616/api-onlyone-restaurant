package models

type Restaurant struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Gmap        string `json:"gmap,omitempty" validate:"required"`
	Iframe      string `json:"iframe,omitempty" validate:"required"`
}
