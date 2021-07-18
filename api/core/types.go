package core

import (
	"context"
	"net/http"
	"time"
)

type Category struct {
	Id        *uint   `json:"id"`
	Name      *string `json:"name"`
	TimesSeen *uint   `json:"times_seen"`
}

type Installments struct {
	Quantity         *uint   `json:"quantity"`
	Fraction         *string `json:"fraction"`
	DecimalSeparator *string `json:"decimal_separator"`
	Cents            *string `json:"cents"`
	Currency         *string `json:"currency"`
}

type Price struct {
	Currency         *string `json:"currency"`
	Symbol           *string `json:"symbol"`
	DecimalSeparator *string `json:"decimal_separator"`
	Fraction         *string `json:"fraction"`
	Cents            *string `json:"cents"`
}

type Discount struct {
	OriginalPrice *Price  `json:"original_price"`
	Offer         *string `json:"offer"`
}

type Attribute struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type ProductVariation struct {
	Id                    *uint        `json:"id"`
	AvailableQuatity      *uint        `json:"available_quatity"`
	AttributesCombination []*Attribute `json:"attributes_combination"`
}

type Shipping struct {
	IsFree      *bool   `json:"is_free"`
	Description *string `json:"description"`
}

type Product struct {
	Id                *uint               `json:"id"`
	Name              *string             `json:"name"`
	Description       *string             `json:"description"`
	ImageLink         *string             `json:"image_link"`
	CategoryId        *uint               `json:"category_id"`
	AvailableQuantity *uint               `json:"available_quantity"` //Sum of available quantity of each variation
	Stars             [5]*uint            `json:"stars"`
	Variations        []*ProductVariation `json:"variations"`
	Shipping          *Shipping           `json:"shipping"`
	Price             *float64            `json:"price"`
	Installments      *Installments       `json:"installments"`
	Link              *string             `json:"link"`
	TimesSeen         *uint               `json:"times_seen"` //Global count of time it has been seeing
	LastSeen          time.Time           `json:"last_seen"`  //Last time the current user have seen it
}

type Service interface {
	GetLastSeenProducts(ctx context.Context) []*Product
	GetTrendingProducts(ctx context.Context) []*Product
	GetPopularCategories(ctx context.Context) []*Category
	GetFilterCategories(ctx context.Context) []*Category
}

type IdentityProvider interface {
	Authenticate(w http.ResponseWriter, r *http.Request)
	Authorize(w http.ResponseWriter, r *http.Request)
}
