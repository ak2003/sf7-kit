package response

import (
	"context"
	"encoding/json"
	"net/http"
)

type CreateResponse struct {
	HttpCode int `json:"http_code"`
	RespBody Body
}

type Body struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func EncodeJson(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	a := response.(CreateResponse)
	w.WriteHeader(a.HttpCode)
	return json.NewEncoder(w).Encode(a.RespBody)
}
