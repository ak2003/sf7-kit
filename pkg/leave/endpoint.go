package leave

import (
	"context"
	"net/http"
	"strings"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetLeaveRequestListing       endpoint.Endpoint
	GetLeaveRequestFilterListing endpoint.Endpoint
	GetDataTypeOfLeave           endpoint.Endpoint
	GetDataRequestFor            endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetLeaveRequestListing:       makeGetLeaveRequestListingEndpoint(s),
		GetLeaveRequestFilterListing: makeGetLeaveRequestFilterListingEndpoint(s),
		GetDataTypeOfLeave:           makeGetDataTypeOfLeaveEndpoint(s),
		GetDataRequestFor:            makeGetDataRequestForEndpoint(s),
	}
}

func makeGetDataRequestForEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetDataRequestForReq)
		err, datas := s.GetDataRequestFor(ctx, req)
		httpCode := http.StatusOK
		var msg string
		msg = http.StatusText(httpCode)
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
			if strings.Contains(err.Error(), "mandatory") {
				httpCode = http.StatusBadRequest
				msg = err.Error()
			} else {
				msg = http.StatusText(httpCode)
			}
		}

		responseBody := response.Body{Message: msg, Data: datas}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetDataTypeOfLeaveEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetDataTypeOfLeaveReq)
		err, datas := s.GetDataTypeOfLeave(ctx, req)
		httpCode := http.StatusOK
		var msg string
		msg = http.StatusText(httpCode)
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
			if strings.Contains(err.Error(), "mandatory") {
				httpCode = http.StatusBadRequest
				msg = err.Error()
			} else {
				msg = http.StatusText(httpCode)
			}
		}

		responseBody := response.Body{Message: msg, Data: datas}
		return response.CreateResponseWithStatusCode{
			ResponseJson: response.CreateResponse{
				Err:      err,
				RespBody: responseBody,
			},
			StatusCode: httpCode,
		}, nil
	}
}

func makeGetLeaveRequestListingEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetLeaveRequestListingRequest)
		err, msg := s.GetLeaveRequestListing(ctx, req)
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

func makeGetLeaveRequestFilterListingEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetLeaveRequestListingFilterRequest)
		err, msg := s.GetLeaveRequestFilterListing(ctx, req)
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
