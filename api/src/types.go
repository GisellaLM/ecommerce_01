package src

type Installments struct {
	Quantity         uint8  `json:"quantity"`
	Fraction         string `json:"fraction"`
	DecimalSeparator string `json:"decimal_separator"`
	Cents            string `json:"cents"`
	Currency         string `json:"currency"`
}

type Price struct {
	Currency         string `json:"currency"`
	Symbol           string `json:"symbol"`
	DecimalSeparator string `json:"decimal_separator"`
	Fraction         string `json:"fraction"`
	Cents            string `json:"cents"`
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
	Id                    *uint8       `json:"id"`
	AvailableQuatity      *uint8       `json:"available_quatity"`
	AttributesCombination []*Attribute `json:"attributes_combination"`
}

type Shipping struct {
	IsFree      *bool   `json:"is_free"`
	Description *string `json:"description"`
}

type Product struct {
	Name              *string             `json:"name"`
	Description       *string             `json:"description"`
	ImageLink         *string             `json:"image_link"`
	CategoryId        *string             `json:"category_id"`
	AvailableQuantity *uint               `json:"available_quantity"`
	Variations        []*ProductVariation `json:"variations"`
	Shipping          *Shipping           `json:"shipping"`
	Price             *Price              `json:"price"`
	Installments      *Installments       `json:"installments"`
}
