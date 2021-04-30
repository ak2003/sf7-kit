package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		NoHp     string `json:"no_hp"`
	}

	CreateLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateLoginResponse struct {
		Token string `json:"token"`
	}

	GetUserRequest struct {
		Id          string `json:"id"`
		Token string `json:"tokenString"`
	}

	GetUserResponse struct {
		Email string `json:"email"`
	}

	CreateResponse struct {
		HttpCode int `json:"http_code"`
		RespBody ResponseBody
	}

	ResponseBody struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	a := response.(CreateResponse)
	w.WriteHeader(a.HttpCode)
	return json.NewEncoder(w).Encode(a.RespBody)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeLoginReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)
	token := r.Header.Get("Authorization")
	req = GetUserRequest{
		Id:   vars["id"],
		Token: token,
	}
	return req, nil
}
