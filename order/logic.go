package order

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"gt-kit/order/helper"
	"gt-kit/order/model"
	pModel "gt-kit/product/model/protoc/model"
	"gt-kit/shared/utils/logger"
)

var logOrder = logger.MakeLogEntry("order", "serviceProduct")

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func serviceProduct() pModel.ProductsClient {
	port := ":7000"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		level.Error(logOrder).Log("could not connect to"+port, err)
	}

	return pModel.NewProductsClient(conn)
}

// @todo Validate if product_id not exist
// @todo Validate qty is empty
// @todo Validate options mandatory not select
func (s service) AddToCart(ctx context.Context, itemCart model.AddToCartRequest) (interface{}, error) {

	var (
		currentItems []model.ItemCart
		total        int64
		lastTotal    int64
		oi           []model.OptionsItemCart
	)

	// Call service product to get information product
	prod := serviceProduct()
	res, err := prod.DetailProduct(context.Background(), &pModel.ProductId{Id: itemCart.ProductID})
	if err != nil {
		level.Error(logOrder).Log("err", err.Error())
	}

	// Read Item Cart from db
	dataCart, _ := s.repository.GetShoppingCart(ctx, itemCart.CartID)

	// Update shopping cart
	if dataCart != nil {
		for _, v := range dataCart.Items {
			currentItems = append(currentItems, v)
		}
		lastTotal = dataCart.Total

		oi, total = helper.DataOptionsItems(ctx, itemCart, res, lastTotal)

		oiString, err := json.Marshal(oi)
		if err != nil {
			level.Error(logOrder).Log("err", err)
		}

		updateQty := false
		for i, v := range currentItems {
			oicString, _ := json.Marshal(v.OptionsItemCart)
			if v.ProductID == itemCart.ProductID && string(oiString) == string(oicString) {
				v.Qty = v.Qty + itemCart.Qty
				currentItems[i] = model.ItemCart{
					ProductID:       v.ProductID,
					Image:           v.Image,
					Qty:             v.Qty,
					Price:           v.Price,
					OptionsItemCart: v.OptionsItemCart,
				}
				updateQty = true
			}
		}

		if updateQty == false {
			currentItems, total, _ = helper.SetItemCart(ctx, itemCart, res, lastTotal, currentItems)
		}

		err = s.repository.UpdateItemShoppingCart(ctx, itemCart.CartID, currentItems, total)
		if err != nil {
			level.Error(logOrder).Log("err", err)
			return nil, err
		}
		return currentItems, nil
	}

	// Insert new shopping cart
	currentItems, total, _ = helper.SetItemCart(ctx, itemCart, res, lastTotal, currentItems)

	// save cart
	sc := model.ShoppingCart{
		ID:       itemCart.CartID,
		UserID:   "6993fad7-dc54-40ba-bca2-aaa2ac3854e0",
		Items:    currentItems,
		MetaData: nil,
		Total:    total,
	}
	err = s.repository.SaveShoppingCart(ctx, sc)
	if err != nil {
		level.Error(logOrder).Log("err", err)
		return nil, err
	}

	return currentItems, nil
}

func (s service) DeleteItemCart(ctx context.Context, param model.DeleteItemCartRequest) (*[]model.ItemCart, error) {
	// Get Data Item Cart
	dataCart, err := s.repository.GetShoppingCart(ctx, param.CartID)
	if err != nil {
		level.Error(logOrder).Log("err", err)
		return nil, err
	}

	var (
		itemCart []model.ItemCart
		tPrice int64
		tPriceOption int64
	)

	tPrice = 0
	tPriceOption = 0
	for i, v := range dataCart.Items {

		// get price itemOptions
		if i == param.IdxItemCart {
			for _, y := range v.OptionsItemCart {
				tPriceOption = tPriceOption + y.Price
			}
			tPrice = v.Price + tPriceOption
		}

		if i != param.IdxItemCart {
			itemCart = append(itemCart, v)
		}
	}

	total := dataCart.Total - tPrice

	err = s.repository.UpdateItemShoppingCart(ctx, param.CartID, itemCart, total)
	if err != nil {
		level.Error(logOrder).Log("err", err)
		return nil, err
	}

	return &itemCart, nil
}
