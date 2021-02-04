package order

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	CreateProductRequest struct {
		Product
	}


)


func decodeCreateProductReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
