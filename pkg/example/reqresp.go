package example

import (
	"context"
	"encoding/json"
	"sf7-kit/pkg/example/model/protoc/model"
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
