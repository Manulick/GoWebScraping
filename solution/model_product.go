package main

type Product struct {
	Description string `json:"description"`

	Price string `json:"price"`

	OldPrice string `json:"oldPrice"`

	Brand string `json:"brand"`

	ImageURL string `json:"imageURL"`
}
