package user

import "context"

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
	LoginUser(ctx context.Context, email string) (string, string, error)
	CheckEmail(ctx context.Context, username string) (int, error)
}