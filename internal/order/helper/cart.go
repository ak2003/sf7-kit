package helper

import (
	"context"
	"encoding/json"
	"gitlab.dataon.com/gophers/sf7-kit/internal/order/model"
	modelProtoc "gitlab.dataon.com/gophers/sf7-kit/internal/product/model/protoc/model"
	"strconv"
)

//var logOrderHelper = logger.MakeLogEntry("order", "helper")

func SetItemCart(ctx context.Context, itemCart model.AddToCartRequest, res *modelProtoc.ProductDetail, lastTotal int64, currentItems []model.ItemCart) ([]model.ItemCart, int64, error) {

	oi, total := DataOptionsItems(ctx, itemCart, res, lastTotal)

	// convert string to optionItemCart
	oiString, err := json.Marshal(oi)
	if err != nil {
		//level.Error(logOrderHelper).Log("err", err)
	}

	for _, v := range currentItems {
		oicString, _ := json.Marshal(v.OptionsItemCart)
		if v.ProductID == itemCart.ProductID && string(oiString) == string(oicString) {
			v.Qty = v.Qty + itemCart.Qty
		}
	}

	currentItems = append(currentItems, model.ItemCart{
		ProductID:       itemCart.ProductID,
		Image:           res.Gallery[0],
		Qty:             itemCart.Qty,
		Price:           res.Price,
		OptionsItemCart: oi,
	})

	return currentItems, total, nil
}

func DataOptionsItems(ctx context.Context, itemCart model.AddToCartRequest, res *modelProtoc.ProductDetail, lastTotal int64) ([]model.OptionsItemCart, int64) {
	var (
		oi    []model.OptionsItemCart
		total int64
	)

	total = res.Price + lastTotal

	// Get OptionsItems
	for _, v := range itemCart.Options {
		opt := res.Options[v.IndexOption]
		ops := opt.ItemOptions[v.IndexSelected]
		// convert to int64
		price, _ := strconv.ParseInt(ops.Price, 10, 64)

		oi = append(oi, model.OptionsItemCart{
			Title:        opt.Title,
			ItemSelected: ops.Value,
			Price:        price,
		})

		total = total + price
	}

	return oi, total
}
