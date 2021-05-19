package custom_field

import (
	"context"
	"encoding/json"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model/protoc/model"
	"net/http"
)

func decodeHealthCheckReq(_ context.Context, r *http.Request) (interface{}, error) {
	var req *model.HealthCheckRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCheckAddFieldReq(_ context.Context, r *http.Request) (interface{}, error) {
	var req *model.AddFieldCheckRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
