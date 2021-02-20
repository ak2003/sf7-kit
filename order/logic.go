package order

import (
	"context"
	"gt-kit/product/model/protoc/model"
	"gt-kit/shared/utils/logger"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

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

func serviceProduct() model.ProductsClient {
	port := ":7000"
	logLogin := logger.MakeLogEntry("order", "serviceProduct")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		level.Error(logLogin).Log("could not connect to"+port, err)
	}

	return model.NewProductsClient(conn)
}

func (s service) AddToCart(ctx context.Context, itemCart AddToCartRequest) (interface{}, error) {
	logLogin := logger.MakeLogEntry("order", "AddToCart")

	// Call service product to get information product
	prod := serviceProduct()
	res, err := prod.DetailProduct(context.Background(), &model.ProductId{Id: itemCart.ProductID})
	if err != nil {
		level.Error(logLogin).Log("err", err.Error())
	}

	// Read Item Cart from db
	var (
		currentItems []ItemCart
		total        int64
		lastTotal    int64
	)
	dataCart, _ := s.repository.GetShoppingCart(ctx, itemCart.CartID)
	if dataCart != nil {
		for _, v := range dataCart.Items {
			currentItems = append(currentItems, v)
		}
		lastTotal = dataCart.Total

		// function
		currentItems, total, _ = s.setItemCart(ctx, itemCart, res, lastTotal, currentItems)
		err = s.repository.UpdateItemShoppingCart(ctx, itemCart.CartID, currentItems, total)
		if err != nil {
			level.Error(logLogin).Log("err", err)
			return nil, err
		}
		return res, nil
	}

	currentItems, total, _ = s.setItemCart(ctx, itemCart, res, lastTotal, currentItems)

	// save cart
	sc := ShoppingCart{
		ID:       itemCart.CartID,
		UserID:   "6993fad7-dc54-40ba-bca2-aaa2ac3854e0",
		Items:    currentItems,
		MetaData: nil,
		Total:    total,
	}
	err = s.repository.SaveShoppingCart(ctx, sc)
	if err != nil {
		level.Error(logLogin).Log("err", err)
		return nil, err
	}

	return res, nil
}

func (s service) setItemCart(ctx context.Context, itemCart AddToCartRequest, res *model.ProductDetail, lastTotal int64, currentItems []ItemCart) ([]ItemCart, int64, error) {
	var (
		oi    []OptionsItemCart
		total int64
	)

	total = res.Price + lastTotal
	for _, v := range itemCart.Options {
		opt := res.Options[v.IndexOption]
		ops := opt.ItemOptions[v.IndexSelected]
		// convert to int64
		price, _ := strconv.ParseInt(ops.Price, 10, 64)
		oi = append(oi, OptionsItemCart{
			Title:        opt.Title,
			ItemSelected: ops.Value,
			Price:        price,
		})

		total = total + price
	}

	currentItems = append(currentItems, ItemCart{
		ProductID:       itemCart.ProductID,
		Image:           res.Gallery[0],
		Qty:             itemCart.Qty,
		Price:           res.Price,
		OptionsItemCart: oi,
	})

	return currentItems, total, nil
}
