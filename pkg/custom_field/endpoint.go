package custom_field

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"
	"net/http"
)

type Endpoints struct {
	HealthCheck   endpoint.Endpoint
	CheckAddField endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CheckAddField: makeCheckAddFieldEndpoint(s),
	}
}

func makeCheckAddFieldEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*model.AddFieldCheckRequest)
		resp, err := s.CheckAddField(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: resp}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}
