package custom_field

import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/config"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	HealthCheck   endpoint.Endpoint
	CheckAddField endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		HealthCheck:   makeHealthCheckEndpoint(s),
		CheckAddField: makeCheckAddFieldEndpoint(s),
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

func makeCheckAddFieldEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*model.AddFieldCheckRequest)
		resp, err := s.CheckAddField(ctx, req)
		responseBody := response.Body{Data: resp}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}
