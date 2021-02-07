package order

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"gt-kit/product/model/protoc/model"
	"gt-kit/shared/utils/logger"
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
	logLogin := logger.MakeLogEntry("order","serviceProduct")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		level.Error(logLogin).Log("could not connect to" + port, err)
	}

	return model.NewProductsClient(conn)
}

func (s service) AddToCart(ctx context.Context, itemCart AddToCartRequest) (interface{}, error) {
	logLogin := logger.MakeLogEntry("order","AddToCart")
	// Call service product to get information product

	// validate data
	level.Info(logLogin).Log("start here")
	// Add To Cart
	prod := serviceProduct()

	// show all registered users
	res1, err := prod.List(context.Background(), &model.ProductId{Id: "1"})
	if err != nil {
		level.Error(logLogin).Log("err", err.Error())
	}
	level.Info(logLogin).Log("in here")
	return res1, nil
}