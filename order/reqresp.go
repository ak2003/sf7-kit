package order

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	AddToCartRequest struct {
		ProductID string `json:"product_id,omitempty"`
		CartID    string `json:"cart_id,omitempty"`
		Qty       int    `json:"qty,omitempty"`
		Options   []Options
	}

	Options struct {
		IndexOption   int `json:"i_option,omitempty"`
		IndexSelected int `json:"i_selected,omitempty"`
	}
)

func decodeAddToCartReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req AddToCartRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
