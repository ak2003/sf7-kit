package order

import (
	"context"
	"encoding/json"
	"gitlab.com/dataon1/sf7-kit/pkg/order/model"
	"net/http"
)

func decodeAddToCartReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.AddToCartRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// decodeDetailProductReq request Method : GET
func decodeDelItemCartReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.DeleteItemCartRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
