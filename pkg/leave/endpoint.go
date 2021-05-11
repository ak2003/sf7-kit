package leave

import (
	"context"
	"net/http"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetLeaveRequestListing       endpoint.Endpoint
	GetLeaveRequestFilterListing endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetLeaveRequestListing:       makeGetLeaveRequestListingEndpoint(s),
		GetLeaveRequestFilterListing: makeGetLeaveRequestFilterListingEndpoint(s),
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
