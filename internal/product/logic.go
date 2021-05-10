package product

import (
	"context"
	"database/sql"
	"github.com/afex/hystrix-go/hystrix"
	"gitlab.dataon.com/gophers/sf7-kit/internal/product/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"

	"github.com/gofrs/uuid"
	"github.com/olivere/elastic/v7"
)

type service struct {
	repository Repository
}

func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
}

func (s service) CreateProduct(ctx context.Context, product interface{}) (interface{}, error) {

	var (
		//logCreate = logger.MakeLogEntry("product", "CreateProduct")
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
		logger.Error(nil, err)
		return &p, err
	}

	err = s.storeToEs(ctx, uid.String(), p, tx)
	if err != nil {
		logger.Error(nil, err)
		return &p, err
	}

	return uid, nil
}

func (s service) storeToEs(ctx context.Context, uid string, p Product, tx *sql.Tx) error {
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
			logger.Error(nil, err)
			// Handle error
			return err
		}
		_, err = client.Index().BodyJson(p).Index("product").Id(uid).Do(context.Background())
		if err != nil {
			// Handle error
			logger.Error(nil, err)
			return err
		}
		tx.Commit()
		return nil
	}, func(err error) error {
		// do this when internal are down
		tx.Rollback()
		logger.Error(nil, err)
		return err
	})

	return err
}

func (s service) DetailProduct(ctx context.Context, param *model.ProductId) (*model.ProductDetail, error) {
	//logDetail := logger.MakeLogEntry("product", "DetailProduct")
	//level.Info(logDetail).Log("param-id", param.Id)

	dp, err := s.repository.DetailProduct(ctx, param.Id)
	if err != nil {
		logger.Error(nil, err)
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
