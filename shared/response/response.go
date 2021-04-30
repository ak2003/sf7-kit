package response

import (
	"context"
	"encoding/json"
	"gt-kit/shared/constant"
	"net/http"
)

type CreateResponse struct {
	Err   error  `json:"error,omitempty"`
	RespBody Body
}

type Body struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func EncodeJson(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	a := response.(CreateResponse)
	//w.WriteHeader(a.HttpCode)
	return json.NewEncoder(w).Encode(a.RespBody)
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	a := response.(CreateResponse)
	if a.Err != nil {
		EncodeError(ctx, a.Err, w)
		return nil
	}
	return json.NewEncoder(w).Encode(a.RespBody)
}

// Encode errors from business-logic
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case constant.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}