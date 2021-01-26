package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gt-kit/shared/utils/config"
	"gt-kit/shared/utils/logger"
	"gt-kit/user/helper"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:    logger,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {

	var (
		logCreate = logger.MakeLogEntry("user","CreateUser")
		pass string
		err error
		numEmail int
	)
	// Validate Email
	if !helper.IsEmailValid(email) {
		return "Invalid Email", nil
	}

	// Check email availability
	numEmail, err = s.repository.CheckEmail(ctx, email)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return "", err
	}

	if numEmail != 0 {
		return "email is already used", err
	}

	// HashPassword
	pass, err = helper.HashAndSalt(ctx, []byte(password))
	if err != nil {
		return "", err
	}
	
	uid, _ := uuid.NewV4()
	id := uid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: pass,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logCreate).Log("err", err)
		return "", err
	}

	level.Info(logCreate).Log("create user", id)

	return "Success", nil
}

func (s service) LoginUser(ctx context.Context, username string, password string) (string, error) {
	logLogin := logger.MakeLogEntry("user","LoginUser")

	email, hashedPwd, err := s.repository.LoginUser(ctx, username)

	if err != nil {
		level.Error(logLogin).Log("err", err)
		return "", err
	}

	if !helper.ComparePasswords(ctx, hashedPwd, []byte(password)) {
		level.Info(logLogin).Log("msg", fmt.Sprintf("Password is wrong for email : %s", email))
		return "", errors.New("wrong password")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	var jwtKey = config.GetString("jwt.key")

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["id_users"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	level.Info(logLogin).Log("msg", fmt.Sprintf("Login is success for email : %s", email))
	return t, nil
}

func (s service) GetUser(ctx context.Context, id string, tokenString string) (string, error) {
	logging := log.With(s.logger, "method", "GetUser")

	arrString := strings.Split(tokenString," ")
	tokenString = arrString[1]

	signingKey := []byte(config.GetString("jwt.key"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	idUsers := claims["id_users"].(string)

	_, err = s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logging).Log("err", err)
		return "", err
	}

	return idUsers, nil
}
