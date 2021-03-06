package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string, tokenString string) (string, error)
	LoginUser(ctx context.Context, username string, password string) (string, error)
}