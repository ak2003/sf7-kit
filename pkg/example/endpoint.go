package example

import (
	"context"
	"gt-kit/pkg/example/model/protoc/model"
	"gt-kit/shared/response"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	HealthCheck endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		HealthCheck: makeHealthCheckEndpoint(s),
	}
}

func makeHealthCheckEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*model.HealthCheckRequest)
		resp, err := s.HealthCheck(ctx, req)
		responseBody := response.Body{Data: resp}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}
