package product

import (
	"context"
	"database/sql"
	"gt-kit/product/model/protoc/model"
	"gt-kit/shared/utils/logger"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
	"github.com/olivere/elastic/v7"
)

type ProductService struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &ProductService{
		repository: rep,
		logger:     logger,
	}
}

// CreateProduct add new comment 1
// CreateProduct add new comment 2
// CreateProduct add new comment 3
// CreateProduct add new comment 4
// CreateProduct add new comment 5
// CreateProduct add new comment 6
// CreateProduct add new comment 7
// CreateProduct add new comment 8
// CreateProduct add new comment 9
// CreateProduct add new comment 10
// CreateProduct add new comment 11
// CreateProduct add new comment 12
// CreateProduct add new comment 13
// CreateProduct add new comment 14
// CreateProduct add new comment 15
// CreateProduct add new comment 16
// CreateProduct add new comment 17
// CreateProduct add new comment 18
// CreateProduct add new comment 19
// CreateProduct add new comment 20
func (s ProductService) CreateProduct(ctx context.Context, product interface{}) (interface{}, error) {

	var (
		logCreate = logger.MakeLogEntry("product", "CreateProduct")
		p         Product
		r         CreateProductRequest
	)

	r = product.(CreateProductRequest)
	//options = make(map[string]Options)
	//
	//for _, index := range r.Options {
	//	uidOpt, _ := uuid.NewV4()
	//	options[uidOpt.String()] = index
	//}

	uid, _ := uuid.NewV4()
	p = Product{
		ID:          uid.String(),
		ProductName: r.ProductName,
		CategoryID:  r.CategoryID,
		BrandID:     r.CategoryID,
		Description: r.Description,
		Price:       r.Price,
		DiscPrice:   r.DiscPrice,
		DiscPercent: r.DiscPercent,
		Options:     r.Options,
		Gallery:     r.Gallery,
		SupplierID:  r.SupplierID,
	}

	tx, err := s.repository.CreateProduct(ctx, p)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return &p, err
	}

	err = s.storeToEs(ctx, uid.String(), p, tx)
	if err != nil {
		level.Info(logCreate).Log("err", err)
		return &p, err
	}

	level.Info(logCreate).Log("msg", uid)
	return uid, nil

}

func (s ProductService) storeToEs(ctx context.Context, uid string, p Product, tx *sql.Tx) error {
	var (
		err    error
		client *elastic.Client
	)
	err = hystrix.Do("store_to_es", func() error {
		// Store to ES
		client, err = elastic.NewClient(
			elastic.SetURL("http://localhost:9200"),
			elastic.SetBasicAuth("elastic", "SySDisTwdLGa8Aeah2ri"),
		)
		if err != nil {
			level.Error(logCreate).Log("err", err)
			// Handle error
			return err
		}
		_, err = client.Index().BodyJson(p).Index("product").Id(uid).Do(context.Background())
		if err != nil {
			// Handle error
			level.Error(logCreate).Log("err", err)
			return err
		}
		tx.Commit()
		return nil
	}, func(err error) error {
		// do this when services are down
		tx.Rollback()
		level.Error(logCreate).Log("err", err)
		return err
	})

	return err
}

func (s ProductService) DetailProduct(ctx context.Context, param *model.ProductId) (*model.ProductDetail, error) {
	logDetail := logger.MakeLogEntry("product", "DetailProduct")
	level.Info(logDetail).Log("param-id", param.Id)

	dp, err := s.repository.DetailProduct(ctx, param.Id)
	if err != nil {
		level.Error(logDetail).Log("err", err)
		return nil, err
	}

	pd := &model.ProductDetail{
		Id:          param.Id,
		ProductName: dp.ProductName,
		Price:       dp.Price,
		Gallery:     dp.Gallery,
		Options:     dp.Options,
	}

	return pd, nil
}
