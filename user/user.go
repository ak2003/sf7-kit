package user

import "context"

type User struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Photo       string `json:"photo"`
	NoHp        string `json:"no_hp"`
	Address     string `json:"address"`
	SubDistrict string `json:"sub_district"`
	ZipCode     string `json:"zip_code"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
	LoginUser(ctx context.Context, email string) (string, string, error)
	CheckEmail(ctx context.Context, username string) (int, error)
}
