package leave

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"
)

func decodeGetLeaveRequestListingReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetLeaveRequestListingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetLeaveRequestFilterListingReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetLeaveRequestListingFilterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetDataRequestForReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetDataRequestForReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
