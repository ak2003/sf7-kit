package product

import (
	"context"
	"encoding/json"
	"gt-kit/product/model/protoc/model"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateProductRequest struct {
		ID          string        `json:"id,omitempty"`
		ProductName string        `json:"product_name,omitempty"`
		CategoryID  int           `json:"category_id,omitempty"`
		BrandID     int           `json:"brand_id,omitempty"`
		Description []Description `json:"description,omitempty"`
		Price       int           `json:"price,omitempty"`
		DiscPrice   int           `json:"disc_price,omitempty"`
		DiscPercent int           `json:"disc_percent,omitempty"`
		Options     []Options     `json:"options,omitempty"`
		Gallery     []string      `json:"gallery,omitempty"`
		SupplierID  int           `json:"supplier_id,omitempty"`
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
