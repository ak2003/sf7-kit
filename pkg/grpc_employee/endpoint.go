package grpc_employee

import (
	"context"
	"net/http"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/grpc_employee/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetEmployeeInformation endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetEmployeeInformation: makeGetEmployeeInformationEndpoint(s),
	}
}

func makeGetEmployeeInformationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*model.GetEmployeeInformationRequest)
		resp, err := s.GetEmployeeInformation(ctx, req)

		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: resp}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}
