package order

import (
	"context"
	"database/sql"
	"github.com/afex/hystrix-go/hystrix"
	"gt-kit/shared/utils/logger"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
	"github.com/olivere/elastic/v7"
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

func (s service) CreateProduct(ctx context.Context, product interface{}) (interface{}, error) {

	var (
		logCreate = logger.MakeLogEntry("product", "CreateProduct")
		p         Product
		r         CreateProductRequest
	)

	r = product.(CreateProductRequest)

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
		Variant:     r.Variant,
		Gallery:     r.Gallery,
		SupplierID:  r.SupplierID,
	}

	tx, err := s.repository.CreateProduct(ctx, p)
	if  err != nil {
		level.Error(logCreate).Log("err", err)
		return &p, err
	}

	err = s.storeToEs(ctx, uid.String(), p, tx)
	if err != nil {
		level.Info(logCreate).Log("err", err)
		return &p, err
	}

	level.Info(logCreate).Log("msg", uid)
	return "problem", nil

}

func (s service) storeToEs(ctx context.Context, uid string, p Product, tx *sql.Tx) error  {
	var (
		err error
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