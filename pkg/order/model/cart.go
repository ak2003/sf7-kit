package model

import "time"

type (
	AddToCartRequest struct {
		ProductID string    `json:"product_id,omitempty"`
		CartID    string    `json:"cart_id,omitempty"`
		Qty       int       `json:"qty,omitempty"`
		Options   []Options `json:"options,omitempty"`
	}

	Options struct {
		IndexOption   int `json:"i_option,omitempty"`
		IndexSelected int `json:"i_selected,omitempty"`
	}

	DeleteItemCartRequest struct {
		CartID      string `json:"cart_id,omitempty"`
		IdxItemCart int    `json:"idx_item_cart,omitempty"`
	}
)

type ShoppingCart struct {
	ID        string
	UserID    string
	Items     []ItemCart
	MetaData  interface{}
	Total     int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ItemCart struct {
	ProductID       string
	Image           string
	Qty             int
	Price           int64
	OptionsItemCart []OptionsItemCart
}

type OptionsItemCart struct {
	Title        string
	ItemSelected string
	Price        int64
}
