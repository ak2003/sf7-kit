package example

import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/example/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/config"

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
		responseBody := response.Body{Data: resp, Message: config.GetString("jwt.key")}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}
