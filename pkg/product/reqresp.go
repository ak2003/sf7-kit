package product

import (
	"context"
	"encoding/json"
	"gitlab.com/dataon1/sf7-kit/pkg/product/model/protoc/model"
	"net/http"

	"github.com/gorilla/mux"
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

// decodeDetailProductReq request Method : GET
func decodeDetailProductReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.ProductId
	vars := mux.Vars(r)

	req = model.ProductId{
		Id: vars["id"],
	}
	return req, nil
}
