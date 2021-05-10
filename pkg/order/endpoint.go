package order

import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/order/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddToCart endpoint.Endpoint
	DeleteItemCart endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		AddToCart:       makeAddToCartEndpoint(s),
		DeleteItemCart : makeDeleteItemCartEndpoint(s),
	}
}

func makeAddToCartEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddToCartRequest)
		msg, err := s.AddToCart(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}

func makeDeleteItemCartEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.DeleteItemCartRequest)
		msg, err := s.DeleteItemCart(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}