package leave

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"

	"github.com/gorilla/schema"
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

func decodeCreateLeaveRequestReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.CreateLeaveRequestReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateLeaveRequestFormReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.CreateLeaveRequestFormReq

	// For request body type application/x-www-form-urlencoded
	err := r.ParseForm()

	if err != nil {
		// Handle error
		return nil, err
	}

	err = schema.NewDecoder().Decode(&req, r.PostForm)
	if err != nil {
		// Handle error
		return nil, err
	}

	return req, nil
}
