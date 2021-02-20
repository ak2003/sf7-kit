package user

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	LoginUser  endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		LoginUser:  makeLoginUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		msg, err := s.CreateUser(ctx, req.Email, req.Password)
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnprocessableEntity
		}
		responseBody := ResponseBody{Message: msg}
		return CreateResponse{HttpCode: httpCode, RespBody: responseBody}, err
	}
}

func makeLoginUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateLoginRequest)
		result, err := s.LoginUser(ctx, req.Email, req.Password)
		var msg string
		msg = "OK"
		httpCode := http.StatusOK
		if err != nil {
			httpCode = http.StatusUnauthorized
			msg = "Username and password is wrong"
		}
		data := CreateLoginResponse{
			Token: result,
		}
		responseBody := ResponseBody{Message: msg, Data: data}
		return CreateResponse{HttpCode: httpCode, RespBody: responseBody}, nil
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, _ := s.GetUser(ctx, req.Id, req.Token)
		responseBody := ResponseBody{Message: "OK", Data: email}
		httpCode := http.StatusOK
		return CreateResponse{HttpCode: httpCode, RespBody: responseBody}, nil
	}
}
