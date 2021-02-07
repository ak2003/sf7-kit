package order

import (
	"context"
	"gt-kit/shared/response"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddToCart endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		AddToCart: makeAddToCartEndpoint(s),
	}
}

func makeAddToCartEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddToCartRequest)
		msg, err := s.AddToCart(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponse{HttpCode: httpCode, RespBody: responseBody}, err
	}
}