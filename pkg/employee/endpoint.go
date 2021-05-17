package employee

import (
	"context"
	"net/http"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/employee/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetEmployeeInformation      endpoint.Endpoint
	GetEmployeeEditInformation  endpoint.Endpoint
	GetEmployeeMasterAddress    endpoint.Endpoint
	UpdateEmployeeMasterAddress endpoint.Endpoint
	CreateEmployeeMasterAddress endpoint.Endpoint
	GetCity                     endpoint.Endpoint
	GetAddressType              endpoint.Endpoint
	GetOwnerStatus              endpoint.Endpoint
	GetStayStatus               endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetEmployeeInformation:      makeGetEmployeeInformationEndpoint(s),
		GetEmployeeEditInformation:  makeGetEmployeeEditInformationEndpoint(s),
		GetEmployeeMasterAddress:    makeGetEmployeeMasterAddressEndpoint(s),
		UpdateEmployeeMasterAddress: makeUpdateEmployeeMasterAddressEndpoint(s),
		CreateEmployeeMasterAddress: makeCreateEmployeeMasterAddressEndpoint(s),
		GetCity:                     makeGetCityEndpoint(s),
		GetAddressType:              makeGetAddressTypeEndpoint(s),
		GetOwnerStatus:              makeGetOwnerStatusEndpoint(s),
		GetStayStatus:               makeGetStayStatusEndpoint(s),
	}
}

func makeCreateEmployeeMasterAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.CreateEmployeeMasterAddressRequest)
		err, msg := s.CreateEmployeeMasterAddress(ctx, req)
		httpCode := http.StatusCreated
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: msg, Data: nil}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeUpdateEmployeeMasterAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.UpdateEmployeeMasterAddressRequest)
		err, msg := s.UpdateEmployeeMasterAddress(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: msg, Data: nil}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetEmployeeMasterAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetEmployeeMasterAddressRequest)
		err, msg := s.GetEmployeeMasterAddress(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetEmployeeInformationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetEmployeeInformationRequest)
		err, msg := s.GetEmployeeInformation(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetEmployeeEditInformationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetEmployeeByIdRequest)
		err, msg := s.GetEmployeeEditInformation(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      nil,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetCityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetCityRequest)
		err, msg := s.GetCity(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetAddressTypeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetAddressTypeRequest)
		err, msg := s.GetAddressType(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetOwnerStatusEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetOwnerStatusRequest)
		err, msg := s.GetOwnerStatus(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetStayStatusEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetStayStatusRequest)
		err, msg := s.GetStayStatus(ctx, req)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := response.Body{Message: http.StatusText(httpCode), Data: msg}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}
