package product

import (
	"context"
	"gt-kit/pkg/product/model/protoc/model"
	"gt-kit/shared/response"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateProduct endpoint.Endpoint
	DetailProduct endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateProduct: makeCreateProductEndpoint(s),
		DetailProduct: makeDetailProductEndpoint(s),
	}
}

func makeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		msg, err := s.CreateProduct(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}

func makeDetailProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.ProductId)
		msg, err := s.DetailProduct(ctx, &req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}